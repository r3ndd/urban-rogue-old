package engine

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	// "github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type View func(*ebiten.Image) error

var views = map[string]View{}

func AddView(name string, view View) {
	views[name] = view
}

func (g *Game) Draw(screen *ebiten.Image) {
	// ebitenutil.DebugPrint(screen, "Hello, world!")

	for _, view := range views {
		err := view(screen)

		if err != nil {
			log.Fatal(err)
		}
	}
}
