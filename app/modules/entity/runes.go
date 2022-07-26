package entity

var Runes = map[TypeId]rune{
	0: '.',
}

func RegisterRune(id TypeId, rune rune) {
	Runes[id] = rune
}
