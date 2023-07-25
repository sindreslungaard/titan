package hh

import (
	"titan/db"
)

var rooms = store[int, Room]()

type Room struct {
	data    db.Room
	tilemap *TileMap
	users   Store[int, RoomUser]

	sigclose chan bool
}

func loadroom(id int) (*Room, bool) {
	r, ok := rooms.find(id)

	if ok {
		return r, ok
	}

	var data db.Room
	if err := db.Conn.Take(&data, id).Error; err != nil {
		return nil, false
	}

	r = &Room{
		data:    data,
		tilemap: tilemap(data.FloorPlan, data.DoorX, data.DoorY),
		users:   store[int, RoomUser](),

		sigclose: make(chan bool),
	}

	rooms.add(r.data.ID, r)

	return r, true
}

func (r *Room) close() {
	rooms.remove(r.data.ID)
}
