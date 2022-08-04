package structure

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

type WallState struct {
	entity.EntityState
}

var BrickWallTypeId entity.TypeId
var WoodWallTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:      "Brick Wall",
		Desc:      "A wall made of brick",
		Rune:      '#',
		Color:     color.RGBA{189, 79, 79, 255},
		Class:     entity.Hybrid,
		InitState: &WallState{},
	}
	BrickWallTypeId = entity.RegisterEntityType(&regData)

	regData = entity.RegData{
		Name:      "Wood Wall",
		Desc:      "A wall made of wood",
		Rune:      'ǁ',
		Color:     color.RGBA{128, 88, 28, 255},
		Class:     entity.Hybrid,
		InitState: &WallState{},
	}
	WoodWallTypeId = entity.RegisterEntityType(&regData)
}
