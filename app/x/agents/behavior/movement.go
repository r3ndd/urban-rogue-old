package behavior

import (
	"github.com/r3ndd/urban-rogue/app/engine/actor"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
)

func RandomWalk(self entity.InstanceId) {
	var dir string

	switch rng.Intn(4) {
	case 0:
		dir = "left"
	case 1:
		dir = "right"
	case 2:
		dir = "up"
	case 3:
		dir = "down"
	}

	turn.ConsumeTurn(self, turn.TurnCap, func() {
		actor.ActSelf(self, "move_self", dir)
	})
}

func RandomEscape(self entity.InstanceId, feared entity.InstanceId) {
	turn.ConsumeTurn(self, turn.TurnCap, func() {
		// actor.ActSelf(self, "move_self", "left")
	})
}
