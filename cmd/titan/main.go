package main

import (
	"os"
	"titan/network"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Level(zerolog.DebugLevel)

	log.Info().Msg("Starting..")
}

func main() {
	log.Fatal().Err(network.NewServer().OnConnect(func(s *network.Socket) {
		println("message")
	}).Listen(2096)).Msg("Server error")
}
