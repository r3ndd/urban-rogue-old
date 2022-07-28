package entity

import "encoding/binary"

type TypeId uint16
type InstanceId [4]byte

type EntityStateBase interface {
	GetTypeId() TypeId
	GetPos() (int, int)
	SetTypeId(TypeId)
	SetPos(int, int)
}

type EntityState struct {
	typeId TypeId
	x      int
	y      int
}

type EntityInfo struct {
	TypeId TypeId
	X      int
	Y      int
	State  EntityStateBase
}

var numActiveEntities uint32 = 0
var EntityStates = map[InstanceId]EntityStateBase{}

func (e *EntityState) GetTypeId() TypeId {
	return e.typeId
}

func (e *EntityState) GetPos() (int, int) {
	return e.x, e.y
}

func (e *EntityState) SetTypeId(typeId TypeId) {
	e.typeId = typeId
}

func (e *EntityState) SetPos(x, y int) {
	e.x = x
	e.y = y
}

func GetNextEntityInstanceId() InstanceId {
	numActiveEntities++
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, numActiveEntities)
	id := InstanceId(*(*[4]byte)(idBytes))
	return id
}
