package hh

import (
	"titan/protocol"
)

type RoomUserEnterMode int

const (
	Regular RoomUserEnterMode = iota
	Teleporter
)

type RoomUser struct {
	id       uint64
	username string
	host     *Option[User]

	actions Actions

	flatctrl int

	x int
	y int
	z float32

	prevX int
	prevY int
	prevZ float32

	direction int
	walking   bool
	setstep   bool

	targetX int
	targetY int

	lockmvmnt bool
	kickcycle int

	needsupdate bool

	walkedon  bool
	walkedoff bool

	entermode RoomUserEnterMode
}

func (u *RoomUser) write(data []byte) {
	h, ok := u.host.unwrap()

	if !ok {
		return
	}

	h.write(data)
}

func (u *RoomUser) point() Point {
	return Point{X: u.x, Y: u.y}
}

func (u *RoomUser) target(x, y int) {
	u.targetX = x
	u.targetY = y
}

func (u *RoomUser) serialize() protocol.SerializedRoomUser {
	h, ok := u.host.unwrap()

	if !ok {
		return protocol.SerializedRoomUser{
			Username: "undefined_host",
		}
	}

	return protocol.SerializedRoomUser{
		ID:               h.id,
		Username:         h.data.Username,
		Motto:            h.data.Motto,
		Look:             h.data.Figure,
		VID:              int(u.id),
		X:                u.x,
		Y:                u.y,
		Z:                u.z,
		Direction:        u.direction,
		Gender:           h.data.Gender,
		AchievementScore: 0,
	}
}

func (u *RoomUser) serializestatus() protocol.SerializedRoomUserStatus {
	fromX := u.x
	fromY := u.y
	fromZ := u.z

	if u.walking {
		fromX = u.prevX
		fromY = u.prevY
		fromZ = u.prevZ
	}

	return protocol.SerializedRoomUserStatus{
		VID:          int(u.id),
		PrevX:        fromX,
		PrevY:        fromY,
		PrevZ:        fromZ,
		HeadRotation: u.direction,
		BodyRotation: u.direction,
		Statuses:     u.actions.serialized(),
	}
}
