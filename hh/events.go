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
			e_releaseversion(s, b)
		case 2419:
			e_securelogin(s, b)

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
	case 357:
		e_userdata(u, b)
	case 273:
		e_credits(u, b)
	case 2596:
		// pong
	case 2110:
		e_navigatordata(u, b)
	case 249:
		e_navigatorsearch(u, b)
	case 2312:
		e_requestroom(u, b)
	case 3898, 2300:
		e_gotoroom(u, b)
	case 3320:
		e_walk(u, b)
	case 1314:
		e_talk(u, b)
	case 2730:
		e_updatefigure(u, b)
	/*
		case 2752:
			RequestCreateRoom(c, p)
		case 2730:
			UserSaveLook(c, p)
		case 1195:
			RequestCatalogMode(c, p)
	*/
	default:
		log.Debug().Int("header", header).Msg("Unknown header")
	}
}
