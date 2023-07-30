package program

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/rs/zerolog/log"
)

var Config Configuration

type TitanConfig struct {
	LogLevel string `toml:"log_level"`
}

type MysqlConfig struct {
	Host        string `toml:"db_host"`
	User        string `toml:"db_user"`
	Pass        string `toml:"db_pass"`
	Name        string `toml:"db_name"`
	Port        int    `toml:"db_port"`
	AutoMigrate bool   `toml:"db_auto_migrate"`
}

type NetworkConfig struct {
	Port       int  `toml:"port"`
	LogPackets bool `toml:"log_packets"`
}

type Configuration struct {
	Titan   TitanConfig   `toml:"titan"`
	Mysql   MysqlConfig   `toml:"mysql"`
	Network NetworkConfig `toml:"network"`
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
		Titan: TitanConfig{
			LogLevel: "info",
		},
		Mysql: MysqlConfig{
			Host:        "127.0.0.1",
			User:        "root",
			Pass:        "",
			Name:        "titan",
			Port:        3306,
			AutoMigrate: true,
		},
		Network: NetworkConfig{
			Port:       2096,
			LogPackets: false,
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
