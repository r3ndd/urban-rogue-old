package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

func init() {
	// Load images
	//
}

func Run() {
	game := Game{}

	err := ebiten.RunGame(&game)

	if err != nil {
		log.Fatal(err)
	}
}

type Game struct{}

func (g *Game) Update() error {
	err := handleInput()
	return err
}

func (g *Game) Layout(windowWidth, windowHeight int) (screenWidth, screenHeight int) {
	return windowWidth, windowHeight
}
