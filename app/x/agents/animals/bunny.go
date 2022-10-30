package animals

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/actor"
	"github.com/r3ndd/urban-rogue/app/engine/agent"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/x/agents/behavior"
)

type BunnyState struct {
	agent.AgentState
	FearedEntity entity.InstanceId
}

var BunnyTypeId entity.TypeId

func init() {
	selfActions := []struct {
		entity.ActionId
		entity.SelfAction
	}{
		{"move_self", actor.MoveSelf},
	}

	targetActions := []struct {
		entity.ActionId
		entity.TargetAction
	}{}

	reactions := []struct {
		entity.ActionId
		entity.Reaction
	}{}

	regData := entity.RegData{
		Name:          "Bunny",
		Rune:          'b',
		Color:         color.RGBA{200, 200, 200, 255},
		Class:         entity.Active,
		InitState:     &BunnyState{},
		SelfActions:   selfActions,
		TargetActions: targetActions,
		Reactions:     reactions,
		ZIndex:        0,
		OnTurn:        BunnyOnTurn,
		AfterTurn:     BunnyAfterTurn,
	}
	BunnyTypeId = entity.RegisterEntityType(&regData)
}

func (state *BunnyState) Copy() entity.EntityStateBase {
	stateCopy := *state
	return &stateCopy
}

func SpawnBunny(x, y int) entity.InstanceId {
	return agent.Spawn(BunnyTypeId, x, y)
}

func BunnyOnTurn(self entity.InstanceId) {
	stateI, _ := entity.GetEntityState(self)
	state := stateI.(*BunnyState)

	switch state.SmState {
	case 0: // Random walk
		behavior.RandomWalk(self)
	case 1: // Escaping from feared entity
		behavior.RandomEscape(self, state.FearedEntity)
	}
}

func BunnyAfterTurn(self entity.InstanceId) {
	stateI, _ := entity.GetEntityState(self)
	state := stateI.(*BunnyState)
	_ = state
}
