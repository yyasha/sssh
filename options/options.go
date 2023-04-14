package options

import (
	"os"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Host         string `short:"b" long:"bind" description:"Host and port to listen on." default:"0.0.0.0:22"`
	Identity     string `short:"i" long:"identity" description:"Private key to identify server with." default:"~/.ssh/id_rsa"`
	PasswordMode bool   `short:"p" long:"password_mode" description:"Enable mandatory password mode"`
	Whitelist    string `long:"whitelist" description:"Optional file of public keys who are allowed to connect."`
}

var Settings Options

func ParceOptions() error {
	_, err := flags.ParseArgs(&Settings, os.Args)
	return err
}
