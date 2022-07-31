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
	TreeTypeId = entity.RegisterEntityType("Tree", "A basic tree", '♠', color.RGBA{33, 191, 83, 255}, entity.Hybrid, &entity.EntityState{}, nil, nil, nil)
}
