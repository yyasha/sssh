package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"t_chat/server"
	"t_chat/structs"
)

type Options struct {
	Verbose  []bool `short:"v" long:"verbose" description:"Show verbose logging."`
	Bind     string `short:"b" long:"bind" description:"Host and port to listen on." default:"0.0.0.0:22"`
	Identity string `short:"i" long:"identity" description:"Private key to identify server with." default:"~/.ssh/id_rsa"`
}

func init() {
	// print logo
	fmt.Fprintf(os.Stdout, structs.Onion_logo, structs.ColorGreen, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorPurple, structs.ColorWhite,
		structs.ColorGreen, structs.ColorNone,
	)
}

/*
TODO:
	+ passwords for users
	+ keys for users
*/

func main() {
	options := Options{Bind: "0.0.0.0:2222", Identity: "/home/yash/ssh-key/t_chat"}

	privateKey, err := ioutil.ReadFile(options.Identity)
	if err != nil {
		log.Println(err)
		return
	}

	server, err := server.NewServer(privateKey)
	if err != nil {
		log.Println(err)
		return
	}

	done, err := server.Start(options.Bind)
	if err != nil {
		log.Println(err)
		return
	}
	<-done
}
