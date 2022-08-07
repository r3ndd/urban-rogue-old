package structure

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

var AsphaltTypeId entity.TypeId
var YellowAsphaltTypeId entity.TypeId
var WhiteAsphaltTypeId entity.TypeId
var PavementTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:  "Asphalt",
		Desc:  "Asphalt for a road",
		Rune:  '·',
		Color: color.RGBA{38, 38, 38, 255},
		Class: entity.Passive,
	}
	AsphaltTypeId = entity.RegisterEntityType(&regData)

	regData = entity.RegData{
		Name:  "Yellow Asphalt",
		Desc:  "Asphalt painted yellow to mark lanes",
		Rune:  '·',
		Color: color.RGBA{142, 148, 84, 255},
		Class: entity.Passive,
	}
	YellowAsphaltTypeId = entity.RegisterEntityType(&regData)

	regData = entity.RegData{
		Name:  "White Asphalt",
		Desc:  "Asphalt painted white to mark lanes",
		Rune:  '·',
		Color: color.RGBA{209, 209, 209, 255},
		Class: entity.Passive,
	}
	WhiteAsphaltTypeId = entity.RegisterEntityType(&regData)

	regData = entity.RegData{
		Name:  "Pavement",
		Desc:  "Standard concrete pavement",
		Rune:  '·',
		Color: color.RGBA{161, 161, 161, 255},
		Class: entity.Passive,
	}
	PavementTypeId = entity.RegisterEntityType(&regData)
}
