package server

import (
	"fmt"
	"log"
	"net"
	"t_chat/rooms"
	"t_chat/utils"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/terminal"
)

// Create server config
func NewServer(privateKey []byte, passwordMode bool) (*Server, error) {
	signer, err := ssh.ParsePrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	config := ssh.ServerConfig{
		NoClientAuth: true,
	}
	config.AddHostKey(signer)

	server := Server{
		sshConfig:    &config,
		sshSigner:    &signer,
		passwordMode: passwordMode,
	}

	return &server, nil
}

func (s *Server) Start(laddr string) (<-chan struct{}, error) {
	// Once a ServerConfig has been configured, connections can be
	// accepted.
	socket, err := net.Listen("tcp", laddr)
	if err != nil {
		return nil, err
	}

	s.socket = &socket
	log.Printf("Listening on %s", laddr)

	// create main room
	rooms.CreateRoom("main")

	go func() {
		for {
			conn, err := socket.Accept()
			if err != nil {
				// TODO: Handle shutdown more gracefully.
				log.Printf("Failed to accept connection, aborting loop: %v", err)
				return
			}

			// From a standard TCP connection to an encrypted SSH connection
			sshConn, channels, requests, err := ssh.NewServerConn(conn, s.sshConfig)
			if err != nil {
				log.Printf("Failed to handshake: %v", err)
				continue
			}

			log.Printf("Connection from: %s, %s, %s", sshConn.RemoteAddr(), sshConn.User(), sshConn.ClientVersion())

			go ssh.DiscardRequests(requests)
			go s.handleChannels(channels, sshConn)
		}
	}()
	return s.done, nil
}

func (s *Server) handleChannels(channels <-chan ssh.NewChannel, conn *ssh.ServerConn) {
	for ch := range channels {
		if t := ch.ChannelType(); t != "session" {
			ch.Reject(ssh.UnknownChannelType, fmt.Sprintf("unknown channel type: %s", t))
			continue
		}

		channel, requests, err := ch.Accept()
		if err != nil {
			continue
		}
		// check username
		err = checkUsername(conn.User())
		if err != nil {
			channel.Write([]byte(fmt.Sprintf("incorrect username: %v\t", err)))
			conn.Close()
		}

		go func(in <-chan *ssh.Request) {
			defer channel.Close()
			for req := range in {
				// log.Println("Request: ", req.Type, string(req.Payload))

				ok := false
				switch req.Type {
				case "shell":
					// We don't accept any commands (Payload),
					// only the default shell.
					if len(req.Payload) == 0 {
						ok = true
					}
				case "pty-req":
					// Responding 'ok' here will let the client
					// know we have a pty ready for input
					ok = true
				case "window-change":
					continue //no response
				}
				req.Reply(ok, nil)
			}
		}(requests)
		go s.handleShell(channel, conn.User())
	}
}

func (s *Server) handleShell(channel ssh.Channel, username string) {
	defer channel.Close()
	// create terminal
	term := terminal.NewTerminal(channel, fmt.Sprintf("%s > ", username))
	// check password
	if s.passwordMode || userExists(username) {
		err := passwordRequest(term, username)
		if err != nil {
			return
		}
	}
	// print logo
	utils.PrintRandomLogo(term)
	// Join main room
	var currentRoom *rooms.Room
	currentRoom, err := rooms.JoinRoom("main", rooms.User{Nickname: username, Terminal: term})
	if err != nil {
		log.Println("Room main not found")
		term.Write([]byte("Room main not found\n"))
		return
	}
	// Recieve user input and send to room
	for {
		// get user input
		line, err := term.ReadLine()
		if err != nil {
			break
		}

		if len(line) > 0 {
			if string(line[0]) == "/" {
				switch line {
				case "/exit":
					return
				case "/help":
					term.Write([]byte(helpMessage))
				case "/new_password":
					err = updatePasswordRequest(term, username)
					if err != nil {
						term.Write([]byte("your password not updated\n"))
					}
				default:
					term.Write([]byte(helpMessage))
				}
				continue
			}

			// send message to room
			currentRoom.Message <- rooms.Message{FromUser: username, Message: line}
		}
	}
}
