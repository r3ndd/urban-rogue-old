package worldgen

import (
	"github.com/r3ndd/urban-rogue/app/engine/world"
	"github.com/r3ndd/urban-rogue/app/x/foliage"
	"github.com/r3ndd/urban-rogue/app/x/structure"
)

const BlockSize = 128
const SuperBlockSize = BlockSize * 4
const MegaBlockSize = SuperBlockSize * 4
const LaneWidth = 4

func GenCity() {
	var margin int = (world.WorldSize - MegaBlockSize*4) / 2

	for y := margin; y < world.WorldSize-margin; y += MegaBlockSize {
		for x := margin; x < world.WorldSize-margin; x += MegaBlockSize {
			GenMegaBlock(x, y)
		}
	}
}

func GenBlock(x, y int) {
	for _y := y; _y < y+BlockSize; _y++ {
		for _x := x; _x < x+BlockSize; _x++ {
			world.CreateEntity(foliage.GrassGroundTypeId, _x, _y, true)
		}
	}

	GenBlockRoads(x, y, BlockSize, BlockSize, 1, true)
}

func GenSuperBlock(x, y int) {
	for _y := y; _y < y+SuperBlockSize; _y += BlockSize {
		for _x := x; _x < x+SuperBlockSize; _x += BlockSize {
			GenBlock(_x, _y)
		}
	}

	GenBlockRoads(x, y, SuperBlockSize, SuperBlockSize, 2, true)
}

func GenMegaBlock(x, y int) {
	for _y := y; _y < y+MegaBlockSize; _y += SuperBlockSize {
		for _x := x; _x < x+MegaBlockSize; _x += SuperBlockSize {
			GenSuperBlock(_x, _y)
		}
	}

	GenBlockRoads(x, y, MegaBlockSize, MegaBlockSize, 3, false)
}

func GenBlockRoads(x, y, width, height, numLanes int, sidewalks bool) {
	if sidewalks {
		for _y := y; _y < y+height; _y++ {
			for _x := x; _x < x+width; _x++ {
				if _x < x+2 || _y < y+2 || _x >= x+width-2 || _y >= y+height-2 {
					world.CreateEntity(structure.PavementTypeId, _x, _y, true)
				}
			}
		}

		x += 2
		y += 2
		width -= 2
		height -= 2
	}

	roadWidth := numLanes*2*LaneWidth + numLanes*2 - 1

	for _y := y; _y < y+height; _y++ {
		for _x := x; _x < x+width; _x++ {
			if _x < x+roadWidth || _y < y+roadWidth || _x >= x+width-roadWidth || _y >= y+height-roadWidth {
				if _x > x && _x < x+width-1 && (_x-x)%LaneWidth == 0 {
					if _x-x == width/2 {
						world.CreateEntity(structure.YellowAsphaltTypeId, _x, _y, true)
					} else {
						world.CreateEntity(structure.WhiteAsphaltTypeId, _x, _y, true)
					}
				} else if _y > y && _y < y+height-1 && (_y-y)%LaneWidth == 0 {
					if _y-y == width/2 {
						world.CreateEntity(structure.YellowAsphaltTypeId, _x, _y, true)
					} else {
						world.CreateEntity(structure.WhiteAsphaltTypeId, _x, _y, true)
					}
				} else {
					world.CreateEntity(structure.AsphaltTypeId, _x, _y, true)
				}
			}
		}
	}
}
