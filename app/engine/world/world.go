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

func CreateEntity(typeId entity.TypeId, x, y int) (entity.InstanceId, bool) {
	class, exists := entity.Classes[typeId]

	if !exists {
		log.Fatal("Attempted to create unregistered entity")
	}

	if x < 0 || y < 0 || x >= worldSize || y >= worldSize {
		return entity.InstanceId{}, false
	}

	tile := &Tiles[y][x]

	if class != entity.Passive {
		if tile.ActiveTypeId != 0 {
			return entity.InstanceId{}, false
		}

		tile.ActiveTypeId = typeId

		if class == entity.Active {
			id := entity.RegisterEntityInstance(typeId, x, y)
			tile.ActiveInstanceId = id
			return id, true
		}

		return entity.InstanceId{}, true
	} else {
		if tile.PassiveTypeId != 0 {
			return entity.InstanceId{}, false
		}

		tile.PassiveTypeId = typeId
		return entity.InstanceId{}, true
	}
}

func DestroyEntityAt(x, y int, active bool) {
	if x < 0 || y < 0 || x >= worldSize || y >= worldSize {
		return
	}

	tile := &Tiles[y][x]

	if active {
		id := tile.ActiveInstanceId

		if _, exists := entity.GetEntityState(id); exists {
			entity.DeleteEntityInstance(id)
		}

		tile.ActiveTypeId = 0
		tile.ActiveInstanceId = [4]byte{}
	} else {
		tile.PassiveTypeId = 0
	}
}

func DestroyActiveEntity(id entity.InstanceId) {
	state, exists := entity.GetEntityState(id)

	if !exists {
		return
	}

	x, y := state.GetPos()
	tile := &Tiles[y][x]

	entity.DeleteEntityInstance(id)
	tile.ActiveTypeId = 0
	tile.ActiveInstanceId = [4]byte{}
}

func GetPassiveEntityAt(x, y int) entity.TypeId {
	return Tiles[y][x].PassiveTypeId
}

func GetActiveEntityAt(x, y int) (*entity.EntityInfo, bool) {
	tile := Tiles[y][x]
	state, exists := entity.GetEntityState(tile.ActiveInstanceId)
	entity := entity.EntityInfo{
		TypeId: tile.ActiveTypeId,
		X:      x,
		Y:      y,
		State:  state,
	}

	return &entity, exists
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

		state, _ := entity.GetEntityState(fromTile.ActiveInstanceId)

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
