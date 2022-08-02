package foliage

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

var GrassGroundTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:  "Grass",
		Desc:  "Simple, green grass",
		Rune:  '`',
		Color: color.RGBA{154, 181, 125, 255},
		Class: entity.Passive,
	}
	GrassGroundTypeId = entity.RegisterEntityType(&regData)
}
