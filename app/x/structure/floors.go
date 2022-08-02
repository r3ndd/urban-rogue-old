package structure

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

var StoneFloorTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:  "Stone Floor",
		Desc:  "A hard floor made of stone",
		Rune:  '·',
		Color: color.RGBA{161, 161, 161, 255},
		Class: entity.Passive,
	}
	StoneFloorTypeId = entity.RegisterEntityType(&regData)
}
