package structure

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

type DoorState struct {
	entity.EntityState
	Open bool
}

var WoodenDoorTypeId entity.TypeId

func init() {
	regData := entity.RegData{
		Name:      "Wooden Door",
		Desc:      "A simple wooden door",
		Rune:      '_',
		Color:     color.RGBA{135, 95, 55, 255},
		Class:     entity.Active,
		InitState: &DoorState{},
	}
	WoodenDoorTypeId = entity.RegisterEntityType(&regData)
}

func (state *DoorState) GetRune() rune {
	if state.Open {
		return '/'
	} else {
		return '_'
	}
}

func (state *DoorState) GetOverlappable() bool {
	return state.Open
}

func (state *DoorState) Copy() entity.EntityStateBase {
	stateCopy := *state
	return &stateCopy
}
