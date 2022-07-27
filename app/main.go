package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"

	_ "github.com/r3ndd/urban-rogue/app/modules/actor"
	_ "github.com/r3ndd/urban-rogue/app/modules/entity"
	_ "github.com/r3ndd/urban-rogue/app/modules/player"
	_ "github.com/r3ndd/urban-rogue/app/modules/world"
)

func main() {
	ebiten.SetWindowTitle("Urban Rogue")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	engine.Run()
}
