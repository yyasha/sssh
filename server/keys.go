package server

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"os"
	"sssh/options"

	"golang.org/x/crypto/ssh"
)

// compare whitelist keys
func compareKeyWithWhitelist(key ssh.PublicKey) error {
	entites, err := os.ReadDir(options.Settings.Whitelist)
	if err != nil {
		return err
	}
	for entity := range entites {
		// get file data
		filebytes, err := os.ReadFile(fmt.Sprintf("%s/%s", options.Settings.Whitelist, entites[entity].Name()))
		if err != nil {
			return err
		}
		// parce public key from file
		pub_key, _, _, _, err := ssh.ParseAuthorizedKey(filebytes)
		if err != nil {
			continue
		}
		// compare
		if bytes.Equal(key.Marshal(), pub_key.Marshal()) {
			return nil
		}
	}
	return errors.New("no entry allowed")
}

// validate user key
func validateKey(conn ssh.ConnMetadata, key ssh.PublicKey) (*ssh.Permissions, error) {
	// compare with all files in whitelist dir
	err := compareKeyWithWhitelist(key)
	if err != nil {
		log.Printf("Failed login attempt from %s for user '%s' client: %s\n", conn.RemoteAddr(), conn.User(), conn.ClientVersion())
		return nil, err
	}
	return nil, nil
}
