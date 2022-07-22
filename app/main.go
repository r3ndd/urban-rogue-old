package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
)

func main() {
	ebiten.SetWindowTitle("Hello, world!")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.MaximizeWindow()

	engine.Run()
}
