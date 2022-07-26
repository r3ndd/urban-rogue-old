package worldgen

import (
	"math/rand"

	"github.com/r3ndd/urban-rogue/app/engine/world"
	"github.com/r3ndd/urban-rogue/app/x/agents/animals"
	"github.com/r3ndd/urban-rogue/app/x/foliage"
)

func init() {
	world.Generator = func() {
		GenTestHouses()
		animals.SpawnBunny(5, 5)
	}
}

func GenTestHouses() {
	for y := 0; y < world.WorldSize; y++ {
		for x := 0; x < world.WorldSize; x++ {
			world.CreateEntity(foliage.GrassGroundTypeId, x, y, false)

			if x == 0 && y == 0 {
				continue
			}

			if rand.Float32() > 0.95 {
				world.CreateEntity(foliage.TreeTypeId, x, y, false)
			}
		}
	}

	GenBrickHouse(10, 10, 20, 10, true)
	GenWoodHouse(30, 30, 20, 10, true)
}
