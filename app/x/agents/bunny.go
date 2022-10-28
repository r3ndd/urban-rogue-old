package agents

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/actor"
	"github.com/r3ndd/urban-rogue/app/engine/agent"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

type BunnyState struct {
	agent.AgentState
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
		OnTurn:        OnTurn,
		AfterTurn:     AfterTurn,
	}
	BunnyTypeId = entity.RegisterEntityType(&regData)
}

func Spawn(x, y int) entity.InstanceId {
	return agent.Spawn(BunnyTypeId, x, y, OnTurn, AfterTurn)
}

func OnTurn() {

}

func AfterTurn() {

}
