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

func (u *User) write(data []byte) {
	u.session.Write(data)
}

func (u *User) welcome() {
	u.write(protocol.SecureLoginOK())
	u.write(protocol.UserHomeRoom(0, 1))
	u.write(protocol.UserEffectsList())
	u.write(protocol.UserClothes())
	u.write(protocol.NewUserIdentity())
	u.write(protocol.UserPermissions(1, 2, false))
	u.write(protocol.AvailabilityStatusMessage(true, false, true))
	u.write(protocol.Ping())
	u.write(protocol.EnableNotifications(true))
	u.write(protocol.UserAchievementScore(0))
	u.write(protocol.IsFirstLoginOfDay(true))
	u.write(protocol.MysteryBoxKeys())
	u.write(protocol.UserClub())

	u.write(protocol.GenericAlert("Hello world"))

	// todo:
	// inventory achievements msg
}