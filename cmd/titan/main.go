package main

import (
	"os"
	"titan/db"
	"titan/hh"
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
	if err := db.Connect(
		os.Getenv("db_user"),
		os.Getenv("db_pass"),
		os.Getenv("db_host"),
		os.Getenv("db_name"),
		os.Getenv("db_port"),
		os.Getenv("db_automigrate") == "true",
	); err != nil {
		log.Fatal().Err(err).Msg("Database connection error")
	}

	go hh.StartAuthenticating()

	log.Fatal().Err(network.NewServer().OnConnect(func(s *network.Socket) {
		s.OnReceive(hh.UnauthenticatedEventHandler(s))
	}).Listen(2096)).Msg("Server error")
}
