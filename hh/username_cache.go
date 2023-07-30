package hh

import (
	"titan/db"
	"titan/program"

	lru "github.com/hashicorp/golang-lru/v2"
	"github.com/rs/zerolog/log"
)

var usernamecache *UsernameCache

type UsernameCache struct {
	lru *lru.Cache[int, string]
}

func InitializeUsernameCache() {
	size := program.Config.Titan.UsernameCacheSize
	log.Info().Int("size", size).Msg("Initialized username cache")

	lru, err := lru.NewWithEvict[int, string](size, func(key int, value string) {
		log.Debug().Int("id", key).Str("username", value).Msg("Cleared username from cache")
	})

	if err != nil {
		log.Fatal().Err(err).Msg("Failed to create username cache")
	}

	usernamecache = &UsernameCache{
		lru: lru,
	}
}

func (cache *UsernameCache) get(id int) string {
	username, ok := cache.lru.Get(id)
	if ok {
		return username
	}

	var u db.User
	err := db.Conn.First(&u, "id = ?", id).Error

	if err != nil {
		return "undefined"
	}

	cache.lru.Add(id, u.Username)
	return u.Username
}
