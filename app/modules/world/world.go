package world

import (
	"fmt"
	"log"
	"unsafe"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

type EntityTile struct {
	PassiveTypeId    entity.TypeId
	ActiveTypeId     entity.TypeId
	ActiveInstanceId entity.InstanceId
}

const worldSize = 100

var Tiles = [worldSize][worldSize]EntityTile{}

func init() {
	fmt.Println(unsafe.Sizeof(EntityTile{}))
}

func CreateEntity(typeId entity.TypeId, initState entity.EntityStateBase, x, y int) bool {
	active, exists := entity.Actives[typeId]

	if !exists {
		log.Fatal("Attempted to create unregistered entity")
	}

	if x < 0 || y < 0 || x >= worldSize || y >= worldSize {
		return false
	}

	tile := &Tiles[y][x]

	if active {
		if tile.ActiveTypeId != 0 {
			return false
		}

		if initState != nil {
			initState.SetTypeId(typeId)
			initState.SetPos(x, y)
		}

		id := entity.GetNextEntityInstanceId()
		tile.ActiveTypeId = typeId
		tile.ActiveInstanceId = id
		entity.EntityStates[id] = initState
	} else {
		if tile.PassiveTypeId != 0 {
			return false
		}

		tile.PassiveTypeId = typeId
	}

	return true
}

func DestroyEntityAt(x, y int, active bool) bool {
	if x < 0 || y < 0 || x >= worldSize || y >= worldSize {
		return false
	}

	tile := &Tiles[y][x]

	if active {
		id := tile.ActiveInstanceId

		if _, exists := entity.EntityStates[id]; exists {
			delete(entity.EntityStates, id)
		}

		tile.ActiveTypeId = 0
		tile.ActiveInstanceId = [4]byte{}
	} else {
		tile.PassiveTypeId = 0
	}

	return true
}

func DestroyActiveEntity(id entity.InstanceId) bool {
	state, exists := entity.EntityStates[id]

	if !exists {
		return false
	}

	if state == nil {
		log.Fatal("Active entities without state must be destroyed by position")
	}

	x, y := state.GetPos()
	tile := &Tiles[y][x]

	delete(entity.EntityStates, id)
	tile.ActiveTypeId = 0
	tile.ActiveInstanceId = [4]byte{}

	return true
}

func GetPassiveEntityAt(x, y int) entity.TypeId {
	return Tiles[y][x].PassiveTypeId
}

func GetActiveEntityAt(x, y int) *entity.EntityInfo {
	tile := Tiles[y][x]
	state := entity.EntityStates[tile.ActiveInstanceId]
	entity := entity.EntityInfo{
		TypeId: tile.ActiveTypeId,
		X:      x,
		Y:      y,
		State:  state,
	}

	return &entity
}

func MoveEntity(fromX, fromY, toX, toY int, active bool) bool {
	if toX < 0 || toY < 0 || toX >= worldSize || toY >= worldSize {
		return false
	}

	fromTile := &Tiles[fromY][fromX]
	toTile := &Tiles[toY][toX]

	if active {
		if toTile.ActiveTypeId != 0 {
			return false
		}

		state := entity.EntityStates[fromTile.ActiveInstanceId]

		if state != nil {
			state.SetPos(toX, toY)
		}

		toTile.ActiveTypeId = fromTile.ActiveTypeId
		toTile.ActiveInstanceId = fromTile.ActiveInstanceId
		fromTile.ActiveTypeId = 0
		fromTile.ActiveInstanceId = [4]byte{}
	} else {
		if fromTile.PassiveTypeId != 0 {
			return false
		}

		toTile.PassiveTypeId = fromTile.PassiveTypeId
		fromTile.PassiveTypeId = 0
	}

	return true
}

func Generate() {
}
