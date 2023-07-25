package hh

import (
	"strconv"
	"strings"
	"sync"
)

const (
	FallbackHeightmap = "xx\nx0"
)

type Point struct {
	X int
	Y int
}

type Tile struct {
	height int
	state  int
}

type TileMap struct {
	heightmap string
	tiles     [][]Tile
	mu        sync.RWMutex

	SizeX int
	SizeY int

	doorX int
	doorY int

	coordinatedMap *CoordinatedMap
}

func tilemap(heightmap string, doorX, doorY int) (ret *TileMap) {

	defer func() {
		if r := recover(); r != nil {
			ret = tilemap(FallbackHeightmap, 0, 0)
		}
	}()

	heightmap = strings.ReplaceAll(heightmap, "\r", "")
	heightmap = strings.ReplaceAll(heightmap, " ", "\n")

	rows := strings.Split(heightmap, "\n")

	if len(rows) < 1 {
		panic("invalid heightmap")
	}

	sizeX := len(rows)
	sizeY := len(rows[0])

	tilemap := make([][]Tile, sizeX)

	for i, row := range rows {

		tilemap[i] = make([]Tile, sizeY)

		for j, char := range row {

			height := (int(char) - '0')

			if height > 9 {
				height = 9
			}

			state := 0
			if char != 'x' {
				state = 1
				if doorX == 0 && doorY == 0 {
					doorX = j
					doorY = i
				}
			}

			tilemap[i][j] = Tile{
				height: height,
				state:  state,
			}
		}

	}

	if string(rows[doorY][doorX]) == "x" {
		doorH := rune('0')

		if len(rows[doorY]) > doorX+1 {
			doorH = rune(rows[doorY][doorX+1])
		}

		modified := []rune(rows[doorY])
		modified[doorX] = doorH
		rows[doorY] = string(modified)
	}

	doorH, err := strconv.Atoi(string(rows[doorY][doorX]))
	if err == nil {
		tilemap[doorY][doorX].state = 1
		tilemap[doorY][doorX].height = doorH
	}

	heightmap = strings.Join(rows, "\n")

	return &TileMap{
		heightmap: heightmap,
		tiles:     tilemap,
		coordinatedMap: &CoordinatedMap{
			points: map[Point]int{},
		},

		SizeX: sizeX - 1,
		SizeY: sizeY - 1,

		doorX: doorX,
		doorY: doorY,
	}

}

func (r *Room) ValidTile(x, y int) bool {

	r.Tilemap.mu.Lock()
	defer r.Tilemap.mu.Unlock()

	if x > r.Tilemap.SizeY || y > r.Tilemap.SizeX {
		return false
	}

	if r.Tilemap.tiles[y][x].state == 0 {
		return false
	}

	return true

}

func (r *Room) TileHeight(x int, y int) int {

	r.Tilemap.mu.RLock()
	defer r.Tilemap.mu.RUnlock()

	return r.Tilemap.tiles[y][x].height

}

func (r *Room) CoordinatedMap() *CoordinatedMap {
	return r.Tilemap.coordinatedMap
}

func HeightmapForModel(model int) string {
	switch model {
	case 1:
		{
			return "xxxxxxxxxxxx\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxx00000000\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 2:
		{
			return "xxxxxxxxxxxx\nxxxxx0000000\nxxxxx0000000\nxxxxx0000000\nxxxxx0000000\nx00000000000\nx00000000000\nx00000000000\nx00000000000\nx00000000000\nx00000000000\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 3:
		{
			return "xxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 4:
		{
			return "xxxxxxxxxxxx\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxx000000x\nxxxxxxxxxxxx"
		}
	case 5:
		{
			return "xxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxx0000000000\nxx0000000000\nxx0000000000\nxx0000000000\nxx0000000000\nxx0000000000\nxx0000000000\nxx0000000000\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 6:
		{
			return "xxxxxxxxxxxx\nxxxxxxx0000x\nxxxxxxx0000x\nxxx00000000x\nxxx00000000x\nxxx00000000x\nxxx00000000x\nx0000000000x\nx0000000000x\nx0000000000x\nx0000000000x\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 7:
		{
			return "xxxxxxxxxxxxxxxxx\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nx0000000000000000\nxxxxxxxxxxxxxxxxx"
		}
	case 8:
		{
			return "xxxxxxxxxxxxxxxxxxxxx\nxxxxxxxxxxx0000000000\nxxxxxxxxxxx0000000000\nxxxxxxxxxxx0000000000\nxxxxxxxxxxx0000000000\nxxxxxxxxxxx0000000000\nxxxxxxxxxxx0000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx0000000000xxxxxxxxxx\nx0000000000xxxxxxxxxx\nx0000000000xxxxxxxxxx\nx0000000000xxxxxxxxxx\nx0000000000xxxxxxxxxx\nx0000000000xxxxxxxxxx\nxxxxxxxxxxxxxxxxxxxxx"
		}
	case 9:
		{
			return "xxxxxxxxxxxxxxxxxxxxxxxxx\nxxxxxxxxxxxxxxxxx00000000\nxxxxxxxxxxxxxxxxx00000000\nxxxxxxxxxxxxxxxxx00000000\nxxxxxxxxxxxxxxxxx00000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nx000000000000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxx0000000000000000\nxxxxxxxxxxxxxxxxxxxxxxxxx"
		}
	case 10:
		{
			return "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nx0000000000000000000000000000\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxx00000000xxxxxxxxxx\nxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
		}
	case 11:
		{
			return "xxxxxxxxxxxxxxxxxxxxx\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx000000xxxxxxxx000000\nx000000x000000x000000\nx000000x000000x000000\nx000000x000000x000000\nx000000x000000x000000\nx000000x000000x000000\nx000000x000000x000000\nx000000xxxxxxxx000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nx00000000000000000000\nxxxxxxxxxxxxxxxxxxxxx"
		}
	case 12:
		{
			return "xxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxx00000\nxxxxxxx00000\nxxxxxxx00000\nxx1111000000\nxx1111000000\nxx1111000000\nxx1111000000\nxx1111000000\nxxxxxxx00000\nxxxxxxx00000\nxxxxxxx00000\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx\nxxxxxxxxxxxx"
		}
	case 13:
		{
			return "xxxxxxxxxxxxxxxxxxx\nxxxxxxxxxxx22222222\nxxxxxxxxxxx22222222\nxxxxxxxxxxx22222222\nxxxxxxxxxxx22222222\nxxxxxxxxxxx22222222\nxxxxxxxxxxx22222222\nx222222222222222222\nx222222222222222222\nx222222222222222222\nx222222222222222222\nx222222222222222222\nx222222222222222222\nx2222xxxxxxxxxxxxxx\nx2222xxxxxxxxxxxxxx\nx2222211111xx000000\nx222221111110000000\nx222221111110000000\nx2222211111xx000000\nxx22xxx1111xxxxxxxx\nxx11xxx1111xxxxxxxx\nx1111xx1111xx000000\nx1111xx111110000000\nx1111xx111110000000\nx1111xx1111xx000000\nxxxxxxxxxxxxxxxxxxx"
		}
	default:
		{
			return "xx\nx0"
		}
	}
}

type CoordinatedMap struct {
	fastLock sync.RWMutex
	points   map[Point]int
}

func (c *CoordinatedMap) Add(point Point) {

	c.fastLock.Lock()
	defer c.fastLock.Unlock()

	_, ok := c.points[point]

	if !ok {
		c.points[point] = 0
	}

	c.points[point]++

}

func (c *CoordinatedMap) Remove(point Point) {

	c.fastLock.Lock()
	defer c.fastLock.Unlock()

	n, ok := c.points[point]

	if !ok || n <= 0 {
		c.points[point] = 0
		return
	}

	c.points[point]--

}

func (c *CoordinatedMap) IsFree(point Point) bool {

	c.fastLock.RLock()
	defer c.fastLock.RUnlock()

	n, ok := c.points[point]

	if !ok || n <= 0 {
		return true
	}

	return false

}
