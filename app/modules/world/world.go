package world

import (
	"github.com/r3ndd/urban-rogue/app/modules/entity"
)

type EntityTile struct {
	PassiveTypeId entity.TypeId
	ActiveTypeId  entity.TypeId
	ActiveEntity  entity.EntityStateBase
}

const worldSize = 100

var Tiles = [worldSize][worldSize]EntityTile{}

func init() {
	// for y := 0; y < worldSize; y++ {
	// 	for x := 0; x < worldSize; x++ {
	// 		tiles[y][x] = entity.Entity{}
	// 	}
	// }
}

func GetPassiveEntityAt(x, y int) entity.TypeId {
	return Tiles[y][x].PassiveTypeId
}

func GetActiveEntityAt(x, y int) *entity.EntityInfo {
	tile := Tiles[y][x]
	entity := entity.EntityInfo{
		TypeId: tile.ActiveTypeId,
		X:      x,
		Y:      y,
		State:  tile.ActiveEntity,
	}

	return &entity
}

func MoveEntity(fromX, fromY, toX, toY int, active bool) bool {
	fromTile := Tiles[fromY][fromX]
	toTile := Tiles[fromY][fromX]

	if active {
		if toTile.ActiveTypeId != 0 {
			return false
		}

		if fromTile.ActiveEntity != nil {
			fromTile.ActiveEntity.SetPos(toX, toY)
		}

		toTile.ActiveTypeId = fromTile.ActiveTypeId
		toTile.ActiveEntity = fromTile.ActiveEntity
		fromTile.ActiveEntity = nil
		fromTile.ActiveTypeId = 0
	} else {
		if fromTile.PassiveTypeId != 0 {
			return false
		}

		toTile.PassiveTypeId = fromTile.PassiveTypeId
		fromTile.ActiveTypeId = 0
	}

	return true
}
