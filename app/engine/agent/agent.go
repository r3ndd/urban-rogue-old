package agent

import (
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
	"github.com/r3ndd/urban-rogue/app/engine/world"
)

type AgentState struct {
	entity.EntityState
	smState int
}

func Spawn(typeId entity.TypeId, x, y int, onTurn, afterTurn func()) entity.InstanceId {
	instId, _ := world.CreateEntity(typeId, x, y, true)
	turn.RegisterActor(instId, onTurn, afterTurn)
	return instId
}
