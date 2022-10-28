package player

import (
	"github.com/r3ndd/urban-rogue/app/engine/actor"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
)

var actionDurations = map[string]int{
	"move":        turn.TurnCap,
	"toggle_door": turn.TurnCap,
}

func DoTurn(turnId string, turn func() bool) {
	if !isTurn {
		return
	}

	for ; digitKeyActive > 0; digitKeyActive-- {
		turnRepeatQueue = append(turnRepeatQueue, turn)
	}

	digitKeyActive = 1
	ProcessActionRepeatQueue()
}

func MoveTurn(dir string) bool {
	return turn.ConsumeTurn(playerId, actionDurations["move"], func() {
		actor.ActSelf(playerId, "move_self", dir)
	})
}

func ToggleDoorTurn() bool {
	state, _ := entity.GetEntityState(playerId)
	x, y := state.GetPos()

	switch actDir {
	case "left":
		x -= 1
	case "right":
		x += 1
	case "up":
		y -= 1
	case "down":
		y += 1
	default:
		return false
	}

	return turn.ConsumeTurn(playerId, actionDurations["toggle_door"], func() {
		actor.ActTarget(playerId, "toggle", x, y)
	})
}
