package rooms

import "golang.org/x/term"

type Room struct {
	Users   []User
	Message chan Message
	Title   string
}

type User struct {
	Nickname string
	Terminal *term.Terminal
}

type Message struct {
	FromUser string
	Message  string
}
