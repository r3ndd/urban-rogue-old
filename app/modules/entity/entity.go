package entity

type TypeId uint16

type EntityStateBase interface {
	GetTypeId() TypeId
	GetPos(int, int)
	SetPos(int, int)
}

type EntityInfo struct {
	TypeId TypeId
	X      int
	Y      int
	State  EntityStateBase
}

type EntityState struct {
	typeId TypeId
	x      int
	y      int
}

func (e *EntityState) GetTypeId() TypeId {
	return e.typeId
}

func (e *EntityState) GetPos() (int, int) {
	return e.x, e.y
}

func (e *EntityState) SetPos(x, y int) {
	e.x = x
	e.y = y
}

func GetEntity()
