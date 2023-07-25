package hh

import (
	"titan/program"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

type EventHandlerUpgrade func(func(b protocol.Buffer))

func UnauthenticatedEventHandler(s Session) func(protocol.Buffer) {
	return func(b protocol.Buffer) {
		defer program.Recover()

		b.ReadInt()
		header := b.ReadShort()

		switch header {
		case 4000:
			releaseversion(s, b)
		case 2419:
			user, ok := securelogin(s, b)
			if !ok {
				return
			}
			s.OnReceive(user.EventHandler)
		default:
			log.Debug().Int("header", header).Msg("Unknown pre-auth header")
		}
	}
}

func (u *User) EventHandler(b protocol.Buffer) {
	defer program.Recover()

	b.ReadInt()
	header := b.ReadShort()

	switch header {
	default:
		log.Debug().Int("header", header).Msg("Unknown header")
	}
}
