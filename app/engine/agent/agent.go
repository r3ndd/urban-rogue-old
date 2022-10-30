package agent

import (
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
	"github.com/r3ndd/urban-rogue/app/engine/world"
)

type AgentState struct {
	entity.EntityState
	SmState int
}

func Spawn(typeId entity.TypeId, x, y int) entity.InstanceId {
	instId, _ := world.CreateEntity(typeId, x, y, true)
	turn.RegisterActor(instId)
	return instId
}
