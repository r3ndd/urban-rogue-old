package actor

import (
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/modules/world"
)

func init() {
	mobileAbility := entity.Ability{
		Actions: map[string]entity.Action{
			"move_self": entity.Action(OnMove),
		},
	}
	entity.RegisterAbility("mobile", mobileAbility)
}

func OnMove(self entity.InstanceId, args ...interface{}) {
	dir := args[0].(string)
	state, exists := entity.EntityStates[self]

	if !exists || state == nil {
		return
	}

	switch dir {
	case "left":
		MoveLeft(state)
	case "right":
		MoveRight(state)
	case "up":
		MoveUp(state)
	case "down":
		MoveDown(state)
	}
}

func MoveLeft(state entity.EntityStateBase) {
	x, y := state.GetPos()
	world.MoveEntity(x, y, x-1, y, true)
}

func MoveRight(state entity.EntityStateBase) {
	x, y := state.GetPos()
	world.MoveEntity(x, y, x+1, y, true)
}

func MoveUp(state entity.EntityStateBase) {
	x, y := state.GetPos()
	world.MoveEntity(x, y, x, y-1, true)
}

func MoveDown(state entity.EntityStateBase) {
	x, y := state.GetPos()
	world.MoveEntity(x, y, x, y+1, true)
}
