package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/engine/player"
	"github.com/r3ndd/urban-rogue/app/engine/turn"
	"github.com/r3ndd/urban-rogue/app/engine/world"

	_ "github.com/r3ndd/urban-rogue/app/x/worldgen"
)

func main() {
	ebiten.SetWindowTitle("Urban Rogue")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	world.Generate()
	player.Spawn()

	go turn.Begin()
	engine.Run()
}
