package hh

import (
	"titan/db"

	"github.com/rs/zerolog/log"
)

var user_auth_queue = make(chan AuthRequest)

type AuthRequest struct {
	sso     string
	session Session
}

func StartAuthenticating() {
	log.Info().Msg("Auth service started..")

	go func() {
		for req := range user_auth_queue {
			var data db.User
			err := db.Conn.Take(&data, "sso = ?", req.sso).Error

			if err != nil {
				log.Debug().Str("sso", req.sso).Msg("Failed to find sso for user")
				req.session.Close()
				return
			}

			u := newuser(req.session, data)

			req.session.OnReceive(u.EventHandler)
			u.welcome()
		}
	}()
}
