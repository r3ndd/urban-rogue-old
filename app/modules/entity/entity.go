package entity

type TypeId uint16

type Entity struct {
	TypeId TypeId
	State  *EntityState
}

type EntityState interface {
	GetPos() (int, int)
	SetPos(int, int) bool
}
