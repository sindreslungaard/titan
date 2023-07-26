package hh

import (
	"math"
)

// Location is an x, y coordinate with its distance from the current position
type Location struct {
	x        int
	y        int
	distance float64
}

// Pathfind finds the next tile available between the current and target pos
func (r *Room) pathfind(fromX, fromY int, fromZ float32, toX, toY int) (x, y int) {
	r.tilemap.mu.RLock()
	defer r.tilemap.mu.RUnlock()

	lenY := len(r.tilemap.tiles)
	lenX := len(r.tilemap.tiles[0])

	coordinatedMap := r.tilemap.coordinatedMap
	allowWalkthrough := r.data.WalkThrough

	locations := make([]Location, 8)
	closest := Location{fromX, fromY, distance(fromX, fromY, toX, toY)}

	locations[0] = Location{fromX - 1, fromY - 1, distance(fromX-1, fromY-1, toX, toY)}
	locations[1] = Location{fromX - 1, fromY + 1, distance(fromX-1, fromY+1, toX, toY)}
	locations[2] = Location{fromX + 1, fromY + 1, distance(fromX+1, fromY+1, toX, toY)}
	locations[3] = Location{fromX + 1, fromY - 1, distance(fromX+1, fromY-1, toX, toY)}
	locations[4] = Location{fromX, fromY - 1, distance(fromX, fromY-1, toX, toY)}
	locations[5] = Location{fromX - 1, fromY, distance(fromX-1, fromY, toX, toY)}
	locations[6] = Location{fromX, fromY + 1, distance(fromX, fromY+1, toX, toY)}
	locations[7] = Location{fromX + 1, fromY, distance(fromX+1, fromY, toX, toY)}

	for _, location := range locations {

		// make sure no fishy business is going on
		if location.x < 0 || location.y < 0 {
			continue
		}

		// make sure the tile is not out of bounds
		if location.x >= lenX || location.y >= lenY {
			continue
		}

		// make sure the tile is free
		if r.tilemap.tiles[location.y][location.x].state == 0 {
			continue
		}

		// if the square is not free
		if !allowWalkthrough && !coordinatedMap.IsFree(Point{X: location.x, Y: location.y}) {
			continue
		}

		var z float32
		var hasChair bool
		walkable := true
		var deductable float32

		/* for _, val := range room.ItemSystem.itemPoints.Get(internalPoint{X: location.x, Y: location.y}).Iter() {

			item := val.(*FloorItem)

			height := item.Height()

			if item.Definition().Chair {
				hasChair = true
				deductable = item.Height()
				z = item.z + height
				break
			}

			if !item.behaviour.Walkable(item) {
				walkable = false
			}

			if item.z+height > z {
				z = item.z + height
			}

		} */

		// don't walk into chairs unless it's our target destination
		if hasChair && (location.x != toX || location.y != toY) {
			continue
		}

		if !walkable && !hasChair {
			continue
		}

		z = float32(r.tilemap.tiles[location.y][location.x].height) + z

		if deductable > 0 {
			z -= deductable / 2
		}

		if z < 0 {
			z = 0
		}

		if z-fromZ > 1.5 {
			continue
		}

		if location.distance < closest.distance {
			closest = location
		}

	}

	return closest.x, closest.y
}

func distance(fromX, fromY, toX, toY int) float64 {
	return math.Sqrt(math.Pow(float64(fromX)-float64(toX), 2) + math.Pow(float64(fromY)-float64(toY), 2))
}

func direction(fromX, fromY, toX, toY int) int {
	dX := math.Round(float64(toX - fromX))
	dY := math.Round(float64(toY - fromY))

	dir := math.Round(4*math.Atan2(dY, dX)/math.Pi + 2)

	if dir < 0 || dir > 7 {
		dir = 7
	}

	return int(dir)
}

func tilestouch(x1, y1, x2, y2 int) bool {

	if !(math.Abs(float64(x1-x2)) > 1 || math.Abs(float64(y1-y2)) > 1) {
		return true
	}

	if x1 == x2 && y1 == y2 {
		return true
	}

	return false

}
