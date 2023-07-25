package hh

import "titan/db"

type Room struct {
	data    db.Room
	tilemap *TileMap
}
