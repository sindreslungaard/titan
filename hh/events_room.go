package hh

import (
	"titan/protocol"
)

func e_walk(u *User, b protocol.Buffer) {
	x := b.ReadInt()
	y := b.ReadInt()

	ru, ok := u.roomuser.unwrap()

	if !ok {
		return
	}

	ru.target(x, y)
}
