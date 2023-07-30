package hh

import (
	"time"
	"titan/db"
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

func e_talk(u *User, b protocol.Buffer) {
	message := b.ReadString()

	ru, ok := u.roomuser.unwrap()

	if !ok {
		return
	}

	r, ok := u.room.unwrap()

	if !ok {
		return
	}

	r.chat(ru, message, false)
}

func e_createroom(u *User, b protocol.Buffer) {
	name := b.ReadString()
	desc := b.ReadString()
	_ = b.ReadString() // model
	_ = b.ReadInt()    // category
	maxusers := b.ReadInt()
	_ = b.ReadInt() // trade type

	var count int64
	db.Conn.Model(&db.Room{}).Where("owner = ?", u.id).Count(&count)

	const MaxRooms = 50

	if count >= MaxRooms {
		u.alert("Max rooms reached")
		return
	}

	if maxusers > 100 {
		maxusers = 100
	} else if maxusers < 1 {
		maxusers = 1
	}

	floorPlan, doorX, doorY := `xxxxxxxxxxxx
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxx000000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxx00000000
xxxxxxxxxxxx
xxxxxxxxxxxx`, 3, 5

	data := db.Room{
		Owner:          u.id,
		Name:           name, // todo: validate
		Description:    desc,
		RoomType:       db.OpenRoom,
		MaxUsers:       maxusers,
		FloorPlan:      floorPlan,
		DoorX:          doorX,
		DoorY:          doorY,
		WalkThrough:    false,
		FloorThickness: -2,
		WallThickness:  -2,
		WallHeight:     -1,
		WallHidden:     false,
	}

	err := db.Conn.Create(&data).Error

	if err != nil {
		u.alert("Something went wrong")
		return
	}

	r, ok := loadroom(data.ID)

	if !ok {
		u.alert("Failed to load created room")
		return
	}

	u.write(protocol.RoomCreated(r.data.ID, r.data.Name))
}
