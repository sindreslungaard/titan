package program

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog/log"
)

var Config Configuration

type MysqlConfig struct {
	Host        string `toml:"db_host"`
	User        string `toml:"db_user"`
	Pass        string `toml:"db_pass"`
	Name        string `toml:"db_name"`
	Port        string `toml:"db_port"`
	AutoMigrate bool   `toml:"db_auto_migrate"`
}

type Configuration struct {
	Mysql MysqlConfig `toml:"mysql"`
}

func LoadConfig() {
	log.Info().Msg("Loading config")

	dat, err := os.ReadFile("./config.toml")

	if err != nil {
		dat = mkdefaultcnf()
	}

	err = toml.Unmarshal(dat, &Config)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to parse config file")
	}
}

func mkdefaultcnf() []byte {
	log.Info().Msg("Couldn't find config file, creating a default one")

	cnf := Configuration{
		Mysql: MysqlConfig{
			Host:        "127.0.0.1",
			User:        "root",
			Pass:        "",
			Name:        "titan",
			Port:        "3306",
			AutoMigrate: true,
		},
	}

	b, err := toml.Marshal(cnf)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to marshal default config struct")
	}

	err = os.WriteFile("./config.toml", b, os.ModePerm)

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to write config file to disk")
	}

	return b
}
