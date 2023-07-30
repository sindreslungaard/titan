package program

import (
	"runtime/debug"

	"github.com/rs/zerolog/log"
)

func Recover() {
	err := recover()
	if err != nil {
		log.Debug().Str(
			"stack",
			string(debug.Stack())).Interface(
			"error",
			err,
		).Msg("Recovered from error")
	}
}

func RecoverAnd(f func()) {
	err := recover()
	if err != nil {
		log.Debug().Str(
			"stack",
			string(debug.Stack())).Interface(
			"error",
			err,
		).Msg("Recovered from error")
		f()
	}
}
