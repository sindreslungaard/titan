package hh

import "github.com/rs/zerolog/log"

type RoomUserEnterMode int

const (
	Regular RoomUserEnterMode = iota
	Teleporter
)

type RoomUser struct {
	id   uint64
	host *Option[User]

	flatctrl int

	x int
	y int
	z float32

	prevX int
	prevY int
	prevZ float32

	direction int
	walking   bool

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
	host, ok := u.host.unwrap()

	if !ok {
		log.Debug().Msg("Can't write to roomuser without host")
		return
	}

	host.write(data)
}

func (u *RoomUser) point() Point {
	return Point{X: u.x, Y: u.y}
}
