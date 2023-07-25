package hh

import (
	"sync/atomic"
	"titan/db"

	"github.com/rs/zerolog/log"
)

var rooms = store[int, Room]()

type Room struct {
	id      int
	data    db.Room
	tilemap *TileMap
	users   Store[int, RoomUser]

	closed   uint32
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
		id:      data.ID,
		data:    data,
		tilemap: tilemap(data.FloorPlan, data.DoorX, data.DoorY),
		users:   store[int, RoomUser](),

		closed:   0,
		sigclose: make(chan bool),
	}

	rooms.add(r.id, r)

	return r, true
}

func (r *Room) close() {
	closed := atomic.LoadUint32(&r.closed) == 1

	if closed {
		return
	}

	log.Debug().Int("id", r.id).Msg("Closing room")

	atomic.StoreUint32(&r.closed, 1)

	// todo: remove roomusers

	r.sigclose <- true
	close(r.sigclose)
	rooms.remove(r.id)

	log.Debug().Int("id", r.id).Msg("Room closed")
}

func (r *Room) broadcast(data []byte, exclude ...*RoomUser) {
	for _, u := range r.users.iter() {
		if len(exclude) > 0 {
			match := false
			for _, e := range exclude {
				if e == u {
					match = true
					break
				}
			}

			if match {
				continue
			}
		}

		u.host.write(data)
	}
}
