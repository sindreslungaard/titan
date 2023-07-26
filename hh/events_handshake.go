package hh

import (
	"titan/db"
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

func e_releaseversion(s Session, b protocol.Buffer) {
	b.ReadString()
}

func e_securelogin(s Session, b protocol.Buffer) {
	sso := b.ReadString()

	var data db.User
	if err := db.Conn.Take(&data, "sso = ?", sso).Error; err != nil {
		log.Debug().Str("sso", sso).Msg("Failed to find sso for user")
		s.Close()
		return
	}

	u := newuser(s, data)

	s.OnReceive(u.EventHandler)
	u.welcome()
}
