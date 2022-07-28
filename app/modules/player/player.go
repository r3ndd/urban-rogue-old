package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/modules/world"
)

type PlayerState struct {
	entity.EntityState
}

var playerTypeId entity.TypeId
var playerState PlayerState = PlayerState{}

func init() {
	abilities := []string{"mobile"}
	playerTypeId = entity.RegisterEntityType("Yourself", "", '@', true, abilities)
}

func Spawn() {
	playerState = PlayerState{}
	world.CreateEntity(playerTypeId, &playerState, 0, 0)

	engine.AddKeyboardListener(ebiten.KeyH, "keydown", func() error {
		MoveLeft()
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyL, "keydown", func() error {
		MoveRight()
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyK, "keydown", func() error {
		MoveUp()
		return nil
	})

	engine.AddKeyboardListener(ebiten.KeyJ, "keydown", func() error {
		MoveDown()
		return nil
	})
}
