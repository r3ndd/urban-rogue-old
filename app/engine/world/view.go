package world

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/utils"
	"golang.org/x/image/font"
)

const fontName string = "agave"
const baseGridSize int = 50
const topMargin int = 16
const bottomMargin int = 32

var fontFace font.Face
var viewX int = 0
var viewY int = 0

func init() {
	fontFace = *engine.LoadFont(fontName, 20)
	engine.AddView("world", view)
}

func view(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()
	viewSize := screenHeight - topMargin - bottomMargin
	sideMargin := (screenWidth - viewSize) / 2
	gridSize := baseGridSize
	tileSize := viewSize / gridSize
	excess := viewSize % gridSize

	if excess >= tileSize {
		gridSize += excess / tileSize
	}

	worldX := viewX / tileSize
	worldY := viewY / tileSize
	viewTiles := Tiles[worldX:utils.IntMin(worldX+gridSize, WorldSize-1)][worldY:utils.IntMin(worldY+gridSize, WorldSize-1)]

	for gridY := 0; gridY < gridSize; gridY++ {
		for gridX := 0; gridX < gridSize; gridX++ {
			var typeId entity.TypeId

			if len(viewTiles) > gridY && len(viewTiles[gridY]) > gridX {
				_entity := viewTiles[gridY][gridX]

				if _entity.ActiveTypeId != 0 {
					typeId = _entity.ActiveTypeId
				} else {
					typeId = _entity.PassiveTypeId
				}
			} else {
				typeId = 0
			}

			tileRune := entity.Runes[typeId]
			tileColor := entity.Colors[typeId]
			x := gridX*tileSize + sideMargin
			y := gridY*tileSize + topMargin
			text.Draw(screen, string(tileRune), fontFace, x, y, tileColor)
		}
	}

	return nil
}
