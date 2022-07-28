package entity

import "encoding/binary"

type TypeId uint16
type InstanceId [InstanceIdBytes]byte
type StatusId string

type EntityStateBase interface {
	GetTypeId() TypeId
	GetPos() (int, int)
	SetTypeId(TypeId)
	SetPos(int, int)
	Copy() EntityStateBase
	IncStatus(StatusId)
	DecStatus(StatusId)
}

type EntityState struct {
	typeId TypeId
	x      int
	y      int
	status map[StatusId]int
}

type EntityInfo struct {
	TypeId TypeId
	X      int
	Y      int
	State  EntityStateBase
}

const InstanceIdBytes = 4

var numActiveEntities uint32 = 0
var entityStates = map[InstanceId]EntityStateBase{}

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

func (e *EntityState) Copy() EntityStateBase {
	stateCopy := *e
	return &stateCopy
}

func (e *EntityState) IncStatus(name StatusId) {
	e.status[name]++
}

func (e *EntityState) DecStatus(name StatusId) {
	e.status[name]--
}

func RegisterEntityInstance(typeId TypeId, x, y int) InstanceId {
	id := getNextEntityInstanceId()
	initState := InitStates[typeId].Copy()
	entityStates[id] = initState
	initState.SetTypeId(typeId)
	initState.SetPos(x, y)
	return id
}

func DeleteEntityInstance(id InstanceId) {
	delete(entityStates, id)
}

func GetEntityState(id InstanceId) (EntityStateBase, bool) {
	state, exists := entityStates[id]
	return state, exists
}

func getNextEntityInstanceId() InstanceId {
	numActiveEntities++
	idBytes := make([]byte, 4)
	binary.LittleEndian.PutUint32(idBytes, numActiveEntities)
	id := InstanceId(*(*[InstanceIdBytes]byte)(idBytes))
	return id
}
