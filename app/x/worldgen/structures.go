package worldgen

import (
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/world"
	"github.com/r3ndd/urban-rogue/app/x/structure"
)

func GenHouse(left, top, width, height int, destroy bool, wallTypeId, floorTypeId, doorTypeId entity.TypeId) {
	for y := top; y < top+height; y++ {
		world.CreateEntity(floorTypeId, left, y, destroy)
		world.CreateEntity(floorTypeId, left+width-1, y, destroy)
		world.CreateEntity(wallTypeId, left, y, destroy)
		world.CreateEntity(wallTypeId, left+width-1, y, destroy)
	}

	for x := left; x < left+width; x++ {
		world.CreateEntity(floorTypeId, x, top, destroy)
		world.CreateEntity(floorTypeId, x, top+height-1, destroy)

		world.CreateEntity(wallTypeId, x, top, destroy)
		world.CreateEntity(wallTypeId, x, top+height-1, destroy)
	}

	for y := top + 1; y < top+height-1; y++ {
		for x := left + 1; x < left+width-1; x++ {
			world.DestroyEntityAt(x, y, true)
			world.CreateEntity(floorTypeId, x, y, destroy)
		}
	}

	world.CreateEntity(doorTypeId, left+1, top, true)
}

func GenBrickHouse(left, top, width, height int, destroy bool) {
	GenHouse(left, top, width, height, destroy, structure.BrickWallTypeId, structure.StoneFloorTypeId, structure.WoodDoorTypeId)
}

func GenWoodHouse(left, top, width, height int, destroy bool) {
	GenHouse(left, top, width, height, destroy, structure.WoodWallTypeId, structure.WoodFloorTypeId, structure.WoodDoorTypeId)
}
