package world

import (
	"github.com/r3ndd/urban-rogue/app/modules/entity"
)

const worldSize = 100

var Tiles = [worldSize][worldSize]entity.Entity{}

func init() {
	// for y := 0; y < worldSize; y++ {
	// 	for x := 0; x < worldSize; x++ {
	// 		tiles[y][x] = entity.Entity{}
	// 	}
	// }
}
