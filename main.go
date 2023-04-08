package main

import (
	"io/ioutil"
	"log"
	"os"
	"t_chat/server"
	"t_chat/utils"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Host         string `short:"a" long:"host" description:"Host and port to listen on." default:"0.0.0.0:22"`
	Identity     string `short:"i" long:"identity" description:"Private key to identify server with." default:"~/.ssh/id_rsa"`
	PasswordMode bool   `short:"p" long:"password_mode" description:"Enable mandatory password mode"`
}

/*
TODO:
	+ keys for users
	+ cached passwords
	+ setting up the number of threads for a room
	+ notif that this username is used
	+ notif that user connected or disconnected
	+ show message for joining users about online
*/

func main() {
	// parse arguments
	var options Options
	_, err := flags.ParseArgs(&options, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%+v", options)
	// print logo
	utils.PrintRandomLogo(os.Stdout)
	// get server privkey from file
	privateKey, err := ioutil.ReadFile(options.Identity)
	if err != nil {
		log.Println(err)
		return
	}
	// create server
	server, err := server.NewServer(privateKey, options.PasswordMode)
	if err != nil {
		log.Println(err)
		return
	}
	// start server
	done, err := server.Start(options.Host)
	if err != nil {
		log.Println(err)
		return
	}
	<-done
}
