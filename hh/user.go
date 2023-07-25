package hh

import (
	"titan/db"
	"titan/protocol"
)

var users = store[int, User]()

type Session interface {
	Close()
	Write([]byte)
	OnReceive(func(protocol.Buffer))
	OnClose(func())
}

type User struct {
	data    db.User
	session Session
}

func newuser(session Session, data db.User) *User {
	_, ok := users.find(data.ID)

	if ok {
		// block login or log the user out
	}

	u := &User{
		data:    data,
		session: session,
	}

	users.add(data.ID, u)

	return u
}
