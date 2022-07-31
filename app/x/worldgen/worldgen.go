package worldgen

import (
	"math/rand"

	"github.com/r3ndd/urban-rogue/app/engine/world"
	"github.com/r3ndd/urban-rogue/app/x/foliage"
)

func init() {
	world.Generator = func() {
		for y := 0; y < world.WorldSize; y++ {
			for x := 0; x < world.WorldSize; x++ {
				if x == 0 && y == 0 {
					continue
				}

				if rand.Float32() > 0.8 {
					world.CreateEntity(foliage.TreeTypeId, x, y)
				}
			}
		}
	}
}
