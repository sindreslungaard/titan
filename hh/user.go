package hh

import (
	"fmt"
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

type Gender string

const (
	Male   Gender = "M"
	Female Gender = "F"
)

type User struct {
	id      int
	data    db.User
	session Session

	room        *Option[Room]
	roomuser    *Option[RoomUser]
	roomrequest *Option[RoomEnterRequest]
}

func newuser(session Session, data db.User) *User {
	clone, ok := users.find(data.ID)

	if ok {
		clone.alert("You were logged in from a different session")
		clone.logout()
	}

	u := &User{
		id:      data.ID,
		data:    data,
		session: session,

		room:        option[Room](),
		roomuser:    option[RoomUser](),
		roomrequest: option[RoomEnterRequest](),
	}

	session.OnClose(u.logout)

	users.add(u.id, u)

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

func (u *User) alert(msg string) {
	u.write(protocol.GenericAlert(msg))
}

func (u *User) gotoroom(id int) {
	r, ok := loadroom(id)

	if !ok {
		u.alert("Couldn't load the room")
		return
	}

	u.exitroom()
	r.mkuser(u)
}

func (u *User) exitroom() {
	r, ok := u.room.unwrap()
	ru, ok2 := u.roomuser.unwrap()

	if ok && ok2 {
		r.rmuser(ru.id)
	} else {
		u.room.clear()
		u.roomuser.clear()
	}
}

func (u *User) updatefigure(g Gender, figure string) {
	u.data.Gender = fmt.Sprint(g)
	u.data.Figure = figure

	db.Conn.Save(&u.data)

	r, ok := u.room.unwrap()

	if !ok {
		return
	}

	ru, ok := u.roomuser.unwrap()

	if !ok {
		return
	}

	r.broadcast(protocol.RoomUserData(ru.id, figure, u.data.Gender, u.data.Motto, 0))
}

func (u *User) logout() {
	u.exitroom()
	users.remove(u.id)
	u.session.Close()
}
