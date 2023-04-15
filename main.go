package main

import (
	"log"
	"os"
	"os/user"
	"sssh/options"
	"sssh/server"
	"sssh/utils"
	"strings"
)

/*
TODO:
	+ allow or disallow registration
	+ keys for users
	+ passphrase support
	+ allowed ips
	+ add rooms
	+ add personal messages
	+ @username and color names
	+ setting up the number of threads for a room (?)
	+ notif that this username is used in other session
	+ notif that user connected (join room) or disconnected
	+ show message for joining users about online
	+ authentificator (Google or another)
	+ notifications
*/

func main() {
	// parse arguments
	err := options.ParceOptions()
	if err != nil {
		log.Fatalln(err)
	}
	// print logo
	utils.PrintRandomLogo(os.Stdout)
	// get server privkey from file
	if strings.HasPrefix(options.Settings.Identity, "~/") {
		user, err := user.Current()
		if err == nil {
			options.Settings.Identity = strings.Replace(options.Settings.Identity, "~", user.HomeDir, 1)
		}
	}
	privateKey, err := os.ReadFile(options.Settings.Identity)
	if err != nil {
		log.Println(err)
		return
	}
	// create server
	server, err := server.NewServer(privateKey, options.Settings.PasswordMode, options.Settings.Whitelist != "")
	if err != nil {
		log.Println(err)
		return
	}
	// start server
	done, err := server.Start(options.Settings.Host)
	if err != nil {
		log.Println(err)
		return
	}
	<-done
}
