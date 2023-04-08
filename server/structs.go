package server

import (
	"net"

	"golang.org/x/crypto/ssh"
)

type Server struct {
	passwordMode bool
	sshConfig    *ssh.ServerConfig
	sshSigner    *ssh.Signer
	socket       *net.Listener
	done         chan struct{}
}

var helpMessage string = `server commands:
/help - show this message
/exit - disconnect from the server
/new_password - new password
`
