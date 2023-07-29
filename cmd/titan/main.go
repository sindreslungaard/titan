package main

import (
	"os"
	"titan/db"
	"titan/hh"
	"titan/network"
	"titan/program"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Level(zerolog.DebugLevel)

	log.Info().Msg("Starting..")
}

func main() {
	program.LoadConfig()
	cnf := program.Config

	if err := db.Connect(
		cnf.Mysql.User,
		cnf.Mysql.Pass,
		cnf.Mysql.Host,
		cnf.Mysql.Name,
		cnf.Mysql.Port,
		cnf.Mysql.AutoMigrate,
	); err != nil {
		log.Fatal().Err(err).Msg("Database connection error")
	}

	hh.StartAuthenticating()

	log.Fatal().Err(network.NewServer().OnConnect(func(s *network.Socket) {
		s.OnReceive(hh.UnauthenticatedEventHandler(s))
	}).Listen(2096)).Msg("Server error")
}
