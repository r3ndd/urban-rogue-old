package player

import (
	"image/color"

	"github.com/r3ndd/urban-rogue/app/engine/actor"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
	"github.com/r3ndd/urban-rogue/app/engine/world"
)

type PlayerState struct {
	entity.EntityState
}

var playerTypeId entity.TypeId
var playerId entity.InstanceId
var playerState PlayerState = PlayerState{}
var isTurn = false
var turnRepeatQueue = []func() bool{}

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
	}{
		{"open", nil},
		{"close", nil},
		{"toggle", nil},
	}

	regData := entity.RegData{
		Name:          "Yourself",
		Rune:          '@',
		Color:         color.White,
		Class:         entity.Active,
		InitState:     &playerState,
		SelfActions:   selfActions,
		TargetActions: targetActions,
		ZIndex:        1,
		OnTurn:        OnTurn,
		AfterTurn:     AfterTurn,
	}
	playerTypeId = entity.RegisterEntityType(&regData)
}

func (state *PlayerState) Copy() entity.EntityStateBase {
	stateCopy := *state
	return &stateCopy
}

func Spawn() {
	playerState = PlayerState{}
	playerId, _ = world.CreateEntity(playerTypeId, 0, 0, true)

	turn.RegisterActor(playerId)
	AddInputListeners()
}

func OnTurn(self entity.InstanceId) {
	isTurn = true
	ProcessActionRepeatQueue()
}

func AfterTurn(self entity.InstanceId) {
	isTurn = false
}

func ProcessActionRepeatQueue() {
	finished := false

	for i := 0; !finished && len(turnRepeatQueue) > 0; i++ {
		turn := turnRepeatQueue[0]
		turnRepeatQueue = turnRepeatQueue[1:]
		finished = turn()
	}
}
