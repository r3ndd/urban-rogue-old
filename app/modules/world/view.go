package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
)

func init() {
	engine.AddView("world", view)
}

func view(screen *ebiten.Image) error {

	return nil
}
