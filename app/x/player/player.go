package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
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

func init() {
	selfActions := []struct {
		entity.ActionId
		entity.SelfAction
	}{
		{"move_self", actor.MoveSelf},
	}

	playerTypeId = entity.RegisterEntityType("Yourself", "", '@', entity.Active, &entity.EntityState{}, selfActions, nil, nil)
}

func Spawn() {
	playerState = PlayerState{}
	playerId, _ = world.CreateEntity(playerTypeId, 0, 0)

	turn.RegisterActor(playerId, OnTurn, AfterTurn)

	engine.AddKeyboardListener(ebiten.KeyH, "keydown", func() error {
		if !isTurn {
			return nil
		}

		turn.ConsumeTurn(playerId, turn.TurnCap, func() {
			actor.ActSelf(playerId, "move_self", "left")
		})
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyL, "keydown", func() error {
		if !isTurn {
			return nil
		}

		turn.ConsumeTurn(playerId, turn.TurnCap, func() {
			actor.ActSelf(playerId, "move_self", "right")
		})
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyK, "keydown", func() error {
		if !isTurn {
			return nil
		}

		turn.ConsumeTurn(playerId, turn.TurnCap, func() {
			actor.ActSelf(playerId, "move_self", "up")
		})
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyJ, "keydown", func() error {
		if !isTurn {
			return nil
		}

		turn.ConsumeTurn(playerId, turn.TurnCap, func() {
			actor.ActSelf(playerId, "move_self", "down")
		})
		return nil
	})
}

func OnTurn() {
	isTurn = true
}

func AfterTurn() {
	isTurn = false
}
