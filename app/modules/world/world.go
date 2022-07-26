package world

import (
	"github.com/r3ndd/urban-rogue/app/engine"
)

type WorldTile struct {
	PassiveEntity *engine.Entity
	ActiveEntity  *engine.Entity
	WorldX        int
	WorldY        int
}

const worldSize = 100

var tiles = [worldSize][worldSize]WorldTile{}

func init() {
	for y := 0; y < worldSize; y++ {
		for x := 0; x < worldSize; x++ {
			tiles[y][x] = WorldTile{
				WorldX: x,
				WorldY: y,
			}
		}
	}
}
