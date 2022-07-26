package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/utils"
	"golang.org/x/image/font"
)

const fontName string = "agave"
const baseGridSize int = 100
const bottomMargin int = 32
const emptyRune rune = '.'

var fontFace font.Face
var viewX int = 0
var viewY int = 0

func init() {
	fontFace = *engine.LoadFont(fontName)
	engine.AddView("world", view)
}

func view(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()
	viewSize := screenHeight - bottomMargin
	sideMargin := (screenWidth - viewSize) / 2
	gridSize := baseGridSize
	tileSize := viewSize / gridSize
	excess := viewSize % gridSize

	if excess >= tileSize {
		gridSize += excess / tileSize
	}

	worldX := viewX / tileSize
	worldY := viewY / tileSize
	viewTiles := tiles[worldX:utils.IntMin(worldX+gridSize, worldSize-1)][worldY:utils.IntMin(worldY+gridSize, worldSize-1)]

	for gridY := 0; gridY < gridSize; gridY++ {
		for gridX := 0; gridX < gridSize; gridX++ {
			var tileRune rune

			if len(viewTiles) > gridY && len(viewTiles[gridY]) > gridX {
				tile := viewTiles[gridY][gridX]

				if tile.ActiveEntity != nil {
					tileRune = tile.ActiveEntity.Rune
				} else if tile.PassiveEntity != nil {
					tileRune = tile.PassiveEntity.Rune
				} else {
					tileRune = emptyRune
				}
			} else {
				tileRune = emptyRune
			}

			x := gridX*tileSize + sideMargin
			y := gridY * tileSize
			text.Draw(screen, string(tileRune), fontFace, x, y, color.White)
		}
	}

	return nil
}
