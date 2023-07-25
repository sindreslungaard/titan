package hh

import (
	"titan/db"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

func releaseversion(s Session, b protocol.Buffer) {
	b.ReadString()
}

func securelogin(s Session, b protocol.Buffer) (*User, bool) {
	sso := b.ReadString()

	var data db.User
	if err := db.Conn.Take(&data, "sso = ?", sso).Error; err != nil {
		log.Debug().Str("sso", sso).Msg("Failed to find sso for user")
		s.Close()
		return nil, false
	}

	u := newuser(s, data)

	return u, true
}
