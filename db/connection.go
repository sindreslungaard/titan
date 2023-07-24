package db

import (
	"fmt"

	"github.com/rs/zerolog/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Conn *gorm.DB

func Connect(user string, password string, host string, dbname string, port string, automigrate bool) error {
	dialector := mysql.Open(fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		dbname,
	))

	config := gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
	}

	db, err := gorm.Open(dialector, &config)

	if err != nil {
		return err
	}

	log.Info().Str("dbname", dbname).Str("user", user).Str("host", host).Str("port", port).Msg("Connected to database")

	if automigrate {
		log.Info().Msg("Migrating database models")

		err = db.AutoMigrate()

		if err != nil {
			return err
		}
	}

	Conn = db

	return nil
}
