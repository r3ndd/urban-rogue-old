package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
)

var dirSelActive bool = false
var actDir string

func AddInputListeners() {
	engine.AddKeyboardListener(ebiten.KeySpace, "keydown", func() {
		if !dirSelActive {
			dirSelActive = true
		} else {
			HandleDir("self")
		}
	})

	engine.AddKeyboardListener(ebiten.KeyH, "keydown", func() { HandleDir("left") })
	engine.AddKeyboardListener(ebiten.KeyL, "keydown", func() { HandleDir("right") })
	engine.AddKeyboardListener(ebiten.KeyK, "keydown", func() { HandleDir("up") })
	engine.AddKeyboardListener(ebiten.KeyJ, "keydown", func() { HandleDir("down") })

	engine.AddKeyboardListener(ebiten.KeyD, "keydown", ToggleDoorTurn)
}

func HandleDir(dir string) {
	if !dirSelActive {
		MoveTurn(dir)
	} else {
		actDir = dir
		dirSelActive = false
	}
}
