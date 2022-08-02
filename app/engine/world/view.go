package world

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/r3ndd/urban-rogue/app/engine"
	"github.com/r3ndd/urban-rogue/app/engine/entity"
	"github.com/r3ndd/urban-rogue/app/utils"
	"golang.org/x/image/font"
)

const fontName string = "agave"
const baseGridSize int = 80
const topMargin int = 16
const bottomMargin int = 32
const fontRatioNum int = 3
const fontRatioDenom int = 2

var fontFace font.Face
var viewX int = 0
var viewY int = 0

func init() {
	fontFace = *engine.LoadFont(fontName, 22)
	engine.AddView("world", view)
}

func view(screen *ebiten.Image) error {
	screenWidth, screenHeight := screen.Size()
	viewSize := screenHeight - topMargin - bottomMargin
	sideMargin := (screenWidth - viewSize) / 2
	gridXSize := baseGridSize
	gridYSize := baseGridSize / fontRatioNum * fontRatioDenom
	tileXSize := viewSize / gridXSize
	tileYSize := viewSize / gridYSize
	yExcess := viewSize % gridYSize

	if yExcess >= tileYSize {
		gridYSize += yExcess / tileYSize
	}

	worldX := viewX / tileXSize
	worldY := viewY / tileYSize
	viewTiles := Tiles[worldY:utils.IntMin(worldY+gridYSize, WorldSize-1)][worldX:utils.IntMin(worldX+gridXSize, WorldSize-1)]

	for gridY := 0; gridY < gridYSize; gridY++ {
		for gridX := 0; gridX < gridXSize; gridX++ {
			var tileRune rune
			var tileColor color.Color

			if len(viewTiles) > gridY && len(viewTiles[gridY]) > gridX {
				_entity := viewTiles[gridY][gridX]
				state, exists := entity.GetEntityState(_entity.ActiveInstanceId)

				if exists {
					tileRune = state.GetRune()
					tileColor = state.GetColor()
				} else {
					var typeId entity.TypeId

					if _entity.ActiveTypeId != 0 {
						typeId = _entity.ActiveTypeId
					} else {
						typeId = _entity.PassiveTypeId
					}

					tileRune = entity.Runes[typeId]
					tileColor = entity.Colors[typeId]
				}
			} else {
				tileRune = entity.Runes[0]
				tileColor = entity.Colors[0]
			}

			x := gridX*tileXSize + sideMargin
			y := gridY*tileYSize + topMargin
			text.Draw(screen, string(tileRune), fontFace, x, y, tileColor)
		}
	}

	return nil
}
