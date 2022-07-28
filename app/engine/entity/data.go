package entity

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

var Actives = map[TypeId]bool{}
var EntityTypeAbilities = map[TypeId][]string{}

func RegisterEntityType(name, desc string, rune rune, active bool, abilities []string) TypeId {
	numEntities++
	typeId := TypeId(numEntities)

	RegisterRune(typeId, rune)
	RegisterName(typeId, name)
	RegisterDesc(typeId, desc)
	RegisterActive(typeId, active)

	if _, exists := EntityTypeAbilities[typeId]; !exists {
		EntityTypeAbilities[typeId] = []string{}
	}

	for _, abilityName := range abilities {
		EntityTypeAbilities[typeId] = append(EntityTypeAbilities[typeId], abilityName)
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

func RegisterActive(id TypeId, active bool) {
	Actives[id] = active
}
