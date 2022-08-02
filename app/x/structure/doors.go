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
	reactions := []struct {
		entity.ActionId
		entity.Reaction
	}{
		{"open", OpenReaction},
		{"close", CloseReaction},
		{"toggle", ToggleReaction},
	}

	regData := entity.RegData{
		Name:      "Wooden Door",
		Desc:      "A simple wooden door",
		Rune:      '_',
		Color:     color.RGBA{135, 95, 55, 255},
		Class:     entity.Active,
		InitState: &DoorState{},
		Reactions: reactions,
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

func OpenReaction(self entity.InstanceId, actor entity.InstanceId, args ...interface{}) {
	stateI, _ := entity.GetEntityState(self)
	state := stateI.(*DoorState)
	state.Open = true
}

func CloseReaction(self entity.InstanceId, actor entity.InstanceId, args ...interface{}) {
	stateI, _ := entity.GetEntityState(self)
	state := stateI.(*DoorState)
	state.Open = false
}

func ToggleReaction(self entity.InstanceId, actor entity.InstanceId, args ...interface{}) {
	stateI, _ := entity.GetEntityState(self)
	state := stateI.(*DoorState)
	state.Open = !state.Open
}
