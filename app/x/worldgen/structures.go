package worldgen

import (
	"github.com/r3ndd/urban-rogue/app/engine/world"
	"github.com/r3ndd/urban-rogue/app/x/structure"
)

func GenHouse(left, top, width, height int, destroy bool) {
	for y := top; y < top+height; y++ {
		world.CreateEntity(structure.StoneFloorTypeId, left, y, destroy)
		world.CreateEntity(structure.StoneFloorTypeId, left+width-1, y, destroy)
		world.CreateEntity(structure.BrickWallTypeId, left, y, destroy)
		world.CreateEntity(structure.BrickWallTypeId, left+width-1, y, destroy)
	}

	for x := left; x < left+width; x++ {
		world.CreateEntity(structure.StoneFloorTypeId, x, top, destroy)
		world.CreateEntity(structure.StoneFloorTypeId, x, top+height-1, destroy)

		world.CreateEntity(structure.BrickWallTypeId, x, top, destroy)
		world.CreateEntity(structure.BrickWallTypeId, x, top+height-1, destroy)
	}

	for y := top + 1; y < top+height-1; y++ {
		for x := left + 1; x < left+width-1; x++ {
			world.DestroyEntityAt(x, y, true)
			world.CreateEntity(structure.StoneFloorTypeId, x, y, destroy)
		}
	}

	world.CreateEntity(structure.WoodenDoorTypeId, left+1, top, true)
}
