package hh

import (
	"titan/protocol"

	"github.com/rs/zerolog/log"
)

func e_releaseversion(s Session, b protocol.Buffer) {
	b.ReadString()
}

func e_securelogin(s Session, b protocol.Buffer) {
	sso := b.ReadString()

	log.Debug().Str("sso", sso).Msg("Submitting auth request to queue")

	user_auth_queue <- AuthRequest{
		sso:     sso,
		session: s,
	}
}
