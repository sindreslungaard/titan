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

	for _, u := range r.users.iter() {
		r.rmuser(u.id)
	}

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

func (r *Room) mkuser(host *User) *RoomUser {
	h, ok := users.find(host.id)

	if ok && h.roomuser.some() {
		// todo: remove from current room
	}

	id := atomic.AddUint64(&r.nextuserid, 1)

	u := &RoomUser{
		id:       id,
		username: host.data.Username,
		host:     option[User]().set(host),
		actions:  actions(),

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

	// for users already in the room
	r.broadcast(protocol.RoomUsers([]protocol.SerializedRoomUser{u.serialize()}), u)
	r.broadcast(protocol.RoomUserStatus([]protocol.SerializedRoomUserStatus{u.serializestatus()}))

	// for the new user
	currentusers := []protocol.SerializedRoomUser{}
	currentstatuses := []protocol.SerializedRoomUserStatus{}

	for _, x := range r.users.iter() {
		currentusers = append(currentusers, x.serialize())
		currentstatuses = append(currentstatuses, x.serializestatus())
	}

	u.write(protocol.RoomHeightmap(
		strings.ReplaceAll(r.data.FloorPlan, "\n", "\r"),
		r.data.WallHeight,
	))

	u.write(protocol.RoomUsers(currentusers))
	u.write(protocol.RoomUserStatus(currentstatuses))

	return u
}

func (r *Room) rmuser(id uint64) {
	u, ok := r.users.find(id)

	if !ok {
		return
	}

	r.tilemap.coordinatedMap.Remove(u.point())
	r.users.remove(u.id)

	h, ok := u.host.unwrap()
	if ok {
		h.roomuser.clear()
		h.room.clear()
	}

	u.host.clear()

	r.broadcast(protocol.RoomUserRemove(u.id))
}

func (r *Room) mv(u *RoomUser, x int, y int, z float32) {
	point := u.point()

	//system.avatarPoints.Remove(point, a.VID)
	r.tilemap.coordinatedMap.Remove(point)

	newPoint := Point{X: x, Y: y}

	//system.avatarPoints.Set(newPoint, a.VID, a)
	r.tilemap.coordinatedMap.Add(newPoint)

	u.prevX = u.x
	u.prevY = u.y
	u.prevZ = u.z

	u.x = x
	u.y = y
	u.z = z

	log.Debug().Uint64("id", u.id).Int("x", x).Int("y", y).Str("user", u.username).Msg("Roomuser moved")
}

func (r *Room) chat(u *RoomUser, message string, shouting bool) {
	r.broadcast(protocol.Talk(u.id, 0, 0, message))
}
