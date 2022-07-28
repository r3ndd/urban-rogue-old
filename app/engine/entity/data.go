package entity

type EntityClass int

const (
	Passive EntityClass = 0
	Active              = 1
	Hybrid              = 2
)

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

var Classes = map[TypeId]EntityClass{}
var InitStates = map[TypeId]EntityStateBase{}

func RegisterEntityType(
	name, desc string,
	rune rune,
	class EntityClass,
	initState EntityStateBase,
	selfActions []struct {
		ActionId
		SelfAction
	},
	targetActions []struct {
		ActionId
		TargetAction
	},
	reactions []struct {
		ActionId
		Reaction
	},
) TypeId {
	numEntities++
	typeId := TypeId(numEntities)

	RegisterRune(typeId, rune)
	RegisterName(typeId, name)
	RegisterDesc(typeId, desc)
	RegisterClass(typeId, class)
	RegisterInitState(typeId, initState)

	if _, exists := EntitySelfActions[typeId]; !exists {
		EntitySelfActions[typeId] = map[ActionId]SelfAction{}
	}

	if _, exists := EntityTargetActions[typeId]; !exists {
		EntityTargetActions[typeId] = map[ActionId]TargetAction{}
	}

	if _, exists := EntityReactions[typeId]; !exists {
		EntityReactions[typeId] = map[ActionId]Reaction{}
	}

	for _, data := range selfActions {
		EntitySelfActions[typeId][data.ActionId] = data.SelfAction
	}

	for _, data := range targetActions {
		EntityTargetActions[typeId][data.ActionId] = data.TargetAction
	}

	for _, data := range reactions {
		EntityReactions[typeId][data.ActionId] = data.Reaction
	}

	return typeId
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

func RegisterClass(id TypeId, class EntityClass) {
	Classes[id] = class
}

func RegisterInitState(id TypeId, initState EntityStateBase) {
	InitStates[id] = initState
}
