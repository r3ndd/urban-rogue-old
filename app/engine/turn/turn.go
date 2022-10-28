package turn

import (
	"sync"

	"github.com/r3ndd/urban-rogue/app/engine/entity"
)

const TurnCap = 128

var TurnEntity entity.InstanceId
var TurnCount int

var turnCaps = map[entity.InstanceId]int{}
var onTurns = map[entity.InstanceId]func(){}
var afterTurns = map[entity.InstanceId]func(){}
var actionQueue = map[entity.InstanceId]func(){}
var endTurn chan entity.InstanceId
var capMu sync.Mutex
var queueMu sync.Mutex

func RegisterActor(id entity.InstanceId, onTurn func(), afterTurn func()) {
	capMu.Lock()
	turnCaps[id] = TurnCap
	capMu.Unlock()

	onTurns[id] = onTurn
	afterTurns[id] = afterTurn
}

func Begin() {
	endTurn = make(chan entity.InstanceId)

	// Process turns indefinitely
	for {
		// Iterate through current actors
		for _, id := range entity.GetEntityInstanceIds() {
			state, exists := entity.GetEntityState(id)

			if !exists {
				continue
			}

			typeId := state.GetTypeId()
			onTurn := entity.OnTurns[typeId]

			if onTurn == nil {
				continue
			}

			TurnEntity = id
			TurnCount++

			queueMu.Lock()

			// Check if action is queued
			if action, exists := actionQueue[id]; exists {
				capMu.Lock()

				// Check if queued action is ready to execute
				if turnCaps[id] >= 0 {
					action()
					delete(actionQueue, id)

					// Check if more actions possible
					if turnCaps[id] > 0 {
						go onTurn()
					} else {
						go EndTurn(id)
					}
				} else {
					go EndTurn(id)
				}

				capMu.Unlock()
			} else {
				go onTurn()
			}

			queueMu.Unlock()

			// Wait for turn of current actor to end
		turn:
			for {
				switch <-endTurn {
				case id:
					break turn
				}
			}
		}
	}
}

func ConsumeTurn(id entity.InstanceId, duration int, action func()) bool {
	if TurnEntity != id {
		return false
	}

	capMu.Lock()
	_, exists := turnCaps[id]

	if !exists {
		return false
	}

	turnCaps[id] -= duration
	remaining := turnCaps[id]
	capMu.Unlock()

	if remaining >= 0 {
		action()

		if remaining > 0 {
			return false
		}
	} else {
		queueMu.Lock()
		actionQueue[id] = action
		queueMu.Unlock()
	}

	EndTurn(id)
	return true
}

func EndTurn(id entity.InstanceId) {
	state, _ := entity.GetEntityState(id)
	typeId := state.GetTypeId()
	entity.AfterTurns[typeId]()

	capMu.Lock()
	turnCaps[id] += TurnCap
	capMu.Unlock()

	endTurn <- id
}
