package actor

import (
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/world"
)

func MoveSelf(self entity.InstanceId, args ...interface{}) {
	state, exists := entity.GetEntityState(self)

	if !exists || state == nil {
		return
	}

	x, y := state.GetPos()
	dir := args[0].(string)

	switch dir {
	case "left":
		MoveLeft(x, y)
	case "right":
		MoveRight(x, y)
	case "up":
		MoveUp(x, y)
	case "down":
		MoveDown(x, y)
	}
}

func MoveLeft(x, y int) {
	world.MoveEntity(x, y, x-1, y, true)
}

func MoveRight(x, y int) {
	world.MoveEntity(x, y, x+1, y, true)
}

func MoveUp(x, y int) {
	world.MoveEntity(x, y, x, y-1, true)
}

func MoveDown(x, y int) {
	world.MoveEntity(x, y, x, y+1, true)
}
