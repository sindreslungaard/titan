package hh

import (
	"sync"
	"titan/db"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

var navigator = &Navigator{}

type RoomWithOwnerName struct {
	db.Room
	OwnerName string
}

type Navigator struct {
	sync.RWMutex
}

func (n *Navigator) query(u *User, view string, query string) {
	log.Debug().Str("view", view).Str("query", query).Msg("navigator query")

	var result []protocol.NavigatorRoom

	switch view {
	case "myworld_view":
		result = n.myrooms(u.id)
	case "hotel_view":
		result = n.loadedrooms()
	default:
		result = []protocol.NavigatorRoom{}
	}

	u.write(protocol.NewNavigatorSearchResults(view, query, result))
}

func (n *Navigator) loadedrooms() []protocol.NavigatorRoom {
	result := []protocol.NavigatorRoom{}

	for _, r := range rooms.iter() {
		result = append(result, protocol.NavigatorRoom{
			ID:          r.data.ID,
			Name:        r.data.Name,
			Description: r.data.Description,
			Owner:       r.data.Owner,
			OwnerName:   r.ownername,
			RoomType:    int(r.data.RoomType),
			MaxUsers:    r.data.MaxUsers,
			UserCount:   r.users.count(),
		})
	}

	return result
}

func (n *Navigator) myrooms(owner int) []protocol.NavigatorRoom {
	result := []protocol.NavigatorRoom{}

	var data []RoomWithOwnerName
	err := db.Conn.Model(&db.Room{}).Select("rooms.*, users.username as owner_name").Joins("left join users on users.id = rooms.owner").Scan(&data).Error

	if err != nil {
		log.Error().Err(err).Int("owner", owner).Msg("Failed to find navigator rooms for user")
		return result
	}

	for _, d := range data {
		usercount := 0
		r, ok := rooms.find(d.ID)
		if ok {
			usercount = r.users.count()
		}

		result = append(result, protocol.NavigatorRoom{
			ID:          d.ID,
			Name:        d.Name,
			Description: d.Description,
			Owner:       d.Owner,
			OwnerName:   d.OwnerName,
			RoomType:    int(d.RoomType),
			MaxUsers:    d.MaxUsers,
			UserCount:   usercount,
		})
	}

	return result
}
