package entity

type AbilityId string
type ActionId string
type Action func(self InstanceId, args ...interface{})
type Reaction func(self InstanceId, actor InstanceId, args ...interface{})
type Ability struct {
	Actions   map[ActionId]Action
	Reactions map[ActionId]Reaction
}

var Abilities = map[AbilityId]Ability{}
var Actions = map[ActionId]

func RegisterAbility(id AbilityId, ability Ability) {
	Abilities[id] = ability
}

func Act(id InstanceId, abilityId AbilityId, actionId ActionId, args ...interface{}) {
	ability, exists := Abilities[abilityId]

	if exists {
		ability.Actions[actionId](id, args...)
	}
}

func React(id InstanceId, actor InstanceId, ability AbilityId, action ActionId, args ...interface{}) {
	ability, exists := Abilities[abilityId]

	Abilities[ability].Reactions[action](id, actor, args...)
}
