package foliage

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

type TreeState struct {
	entity.EntityState
}

var TreeTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:      "Tree",
		Desc:      "A basic tree",
		Rune:      '♠',
		Color:     color.RGBA{33, 191, 83, 255},
		Class:     entity.Hybrid,
		InitState: &TreeState{},
	}
	TreeTypeId = entity.RegisterEntityType(&regData)
}
