package entity

import "reflect"

var numEntities = 0

var Runes = map[TypeId]rune{
	0: '.',
}

var Names = map[TypeId]string{
	0: "",
}

var Descs = map[TypeId]string{
	0: "",
}

var StateTypes map[TypeId]reflect.Type

func RegisterEntityType(name, desc string, rune rune, stateType reflect.Type) {
	numEntities++
	typeId := TypeId(numEntities)

	RegisterRune(typeId, rune)
	RegisterName(typeId, name)
	RegisterDesc(typeId, desc)
}

func RegisterRune(id TypeId, rune rune) {
	Runes[id] = rune
}

func RegisterName(id TypeId, name string) {
	Names[id] = name
}

func RegisterDesc(id TypeId, desc string) {
	Descs[id] = desc
}

func RegisterStateType(id TypeId, stateType reflect.Type) {
	StateTypes[id] = stateType
}
