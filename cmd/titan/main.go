package main

import (
	"os"
	"strings"
	"titan/db"
	"titan/hh"
	"titan/network"
	"titan/program"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Starting..")
}

func main() {
	program.LoadConfig()
	cnf := program.Config

	level := zerolog.InfoLevel

	switch strings.ToLower(cnf.Titan.LogLevel) {
	case "trace":
		level = zerolog.TraceLevel
	case "debug":
		level = zerolog.DebugLevel
	case "info":
		level = zerolog.InfoLevel
	case "warn":
		level = zerolog.WarnLevel
	case "error":
		level = zerolog.ErrorLevel
	case "fatal":
		level = zerolog.FatalLevel
	default:
		level = zerolog.InfoLevel
		log.Error().Str("level", cnf.Titan.LogLevel).Msg("Invalid log level, using 'info'")
	}

	log.Info().Msgf("Using log level %s", level.String())
	zerolog.SetGlobalLevel(level)

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
