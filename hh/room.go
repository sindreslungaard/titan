package hh

import (
	"strings"
	"sync/atomic"
	"time"
	"titan/db"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

var rooms = store[int, Room]()

type RoomEnterRequest struct {
	id        int
	password  string
	timestamp time.Time
}

type Room struct {
	id      int
	data    db.Room
	tilemap *TileMap
	users   Store[uint64, RoomUser]

	nextuserid uint64
	nextitemid uint64

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
		users:   store[uint64, RoomUser](),

		closed:   0,
		sigclose: make(chan bool),
	}

	rooms.add(r.id, r)
	go r.process()

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

		u.write(data)
	}
}

func (r *Room) newroomuser(host *User) *RoomUser {
	h, ok := users.find(host.id)

	if ok && h.roomuser.some() {
		// todo: remove from current room
	}

	id := atomic.AddUint64(&r.nextuserid, 1)

	u := &RoomUser{
		id:   id,
		host: option[User]().set(host),

		x: r.tilemap.doorX,
		y: r.tilemap.doorY,
		z: float32(r.tilemap.tiles[r.tilemap.doorY][r.tilemap.doorX].height),

		prevX: -1,
		prevY: -1,

		targetX: r.tilemap.doorX,
		targetY: r.tilemap.doorY,

		kickcycle: -1,
		direction: 2,
		flatctrl:  0,
	}

	// todo: handle teleporting
	// todo: handle room rights

	r.tilemap.coordinatedMap.Add(u.point())
	r.users.add(u.id, u)

	host.room.set(r)
	host.roomuser.set(u)

	u.write(protocol.RoomHeightmap(
		strings.ReplaceAll(r.data.FloorPlan, "\n", "\r"),
		r.data.WallHeight,
	))

	return u
}
