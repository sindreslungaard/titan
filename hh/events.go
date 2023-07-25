package hh

import (
	"time"
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
			user.welcome()
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
		u.write(protocol.UserData(
			u.data.ID,
			u.data.Username,
			u.data.Figure,
			u.data.Gender,
			u.data.Motto,
			u.data.Respect,
			u.data.RespectToGive,
			u.data.PetRespectToGive,
		))
		// todo: user perks msg
		// todo: memenu settings
	case 273:
		u.write(protocol.UserCredits(u.data.Credits))
	case 2596:
		// pong
	case 2110:
		/*
			todo:
			NewNavigatorSettings
			NewNavigatorSavedSearches
			NewNavigatorEventCategories
		*/
		u.write(protocol.NewNavigatorMetaData())
		u.write(protocol.NewNavigatorLiftedRooms())
		u.write(protocol.NewNavigatorCollapsedCategories())
	case 249:
		view := b.ReadString()
		query := b.ReadString()
		u.write(protocol.NewNavigatorSearchResults(view, query))
	case 2312:
		u.roomrequest.clear()

		// todo: leave current room

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
	case 3898, 2300:
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
	/*
		case 2752:
			RequestCreateRoom(c, p)

		case 3320:
			Walk(c, p)
		case 1314:
			Talk(c, p)
		case 2730:
			UserSaveLook(c, p)
		case 1195:
			RequestCatalogMode(c, p)
	*/
	default:
		log.Debug().Int("header", header).Msg("Unknown header")
	}
}
