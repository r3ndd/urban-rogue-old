package player

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/engine/world"
)

var dirSelActive bool = false
var shiftActive bool = false
var ctrlActive bool = false
var digitKeyActive int = 1
var actDir string

func AddInputListeners() {
	engine.AddKeyboardListener(ebiten.KeyShift, "keydown", func() {
		shiftActive = true
	})
	engine.AddKeyboardListener(ebiten.KeyControl, "keydown", func() {
		ctrlActive = true
	})
	engine.AddKeyboardListener(ebiten.KeyShift, "keyup", func() {
		shiftActive = false
	})
	engine.AddKeyboardListener(ebiten.KeyControl, "keyup", func() {
		ctrlActive = false
	})

	engine.AddKeyboardListener(ebiten.KeySpace, "keydown", func() {
		if !dirSelActive {
			dirSelActive = true
		} else {
			HandleDir("self")
		}
	})

	engine.AddDigitKeyListener("keydown", func(digit int) {
		if digit != 0 {
			digitKeyActive = digit
		}
	})

	engine.AddKeyboardListener(ebiten.KeyH, "keydown", func() { HandleDir("left") })
	engine.AddKeyboardListener(ebiten.KeyL, "keydown", func() { HandleDir("right") })
	engine.AddKeyboardListener(ebiten.KeyK, "keydown", func() { HandleDir("up") })
	engine.AddKeyboardListener(ebiten.KeyJ, "keydown", func() { HandleDir("down") })

	engine.AddKeyboardListener(ebiten.KeyD, "keydown", func() { DoTurn("toggle_door", ToggleDoorTurn) })
}

func HandleDir(dir string) {
	if shiftActive {
		world.MoveView(dir)
	} else if dirSelActive {
		actDir = dir
		dirSelActive = false
	} else {
		DoTurn("move", func() bool { return MoveTurn(dir) })
	}
}
