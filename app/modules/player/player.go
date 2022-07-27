package player

import (
	"reflect"

	"github.com/r3ndd/urban-rogue/app/modules/entity"
)

type PlayerState struct {
	entity.EntityState
}

func init() {
	entity.RegisterEntityType("Yourself", "", '@', reflect.TypeOf(PlayerState{}))
}
