package hh

import (
	"time"
	"titan/protocol"
)

func e_requestroom(u *User, b protocol.Buffer) {
	u.roomrequest.clear()

	id := b.ReadInt()
	password := b.ReadString()

	r, ok := loadroom(id)

	if !ok {
		u.alert("Failed to load room")
		return
	}

	if r.data.Password != password {
		u.alert("Incorrect password")
		return
	}

	u.roomrequest.set(&RoomEnterRequest{
		id:        r.id,
		password:  password,
		timestamp: time.Now(),
	})

	u.write(protocol.RoomOpen())
	u.write(protocol.RoomModel(r.id))
}

func e_gotoroom(u *User, b protocol.Buffer) {
	roomrequest, ok := u.roomrequest.unwrap()
	if !ok {
		u.alert("Invalid room request")
		return
	}

	if roomrequest.timestamp.After(time.Now().Add(time.Second * 5)) {
		u.alert("Expired room request")
		return
	}

	u.gotoroom(roomrequest.id)
}

func e_walk(u *User, b protocol.Buffer) {
	x := b.ReadInt()
	y := b.ReadInt()

	ru, ok := u.roomuser.unwrap()

	if !ok {
		return
	}

	ru.target(x, y)
}
