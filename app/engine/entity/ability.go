package entity

type ActionId string
type SelfAction func(self InstanceId, args ...interface{})
type TargetAction func(self InstanceId, target InstanceId, args ...interface{})
type Reaction func(self InstanceId, actor InstanceId, args ...interface{})

var EntitySelfActions = map[TypeId]map[ActionId]SelfAction{}
var EntityTargetActions = map[TypeId]map[ActionId]TargetAction{}
var EntityReactions = map[TypeId]map[ActionId]Reaction{}
