package main

import (
	"github.com/GyeongHoKim/tiny-pet/internal/navlogic"
)

const (
	IDLE_STATE               = navlogic.StateIdle
	MOVING_STATE             = navlogic.StateMoving
	OBSTACLE_AVOIDANCE_STATE = navlogic.StateObstacleAvoidance
	EDGE_AVOIDANCE_STATE     = navlogic.StateEdgeAvoidance
	INTERACTING_STATE        = navlogic.StateInteracting
)

const (
	RANDOM_WALK_MODE = iota
	GUARD_MODE
	INTERACTIVE_MODE
)

type NavigationModule struct {
	motorController *MotorController
	sensorModule    *SensorModule
	currentState    int
	behaviorMode    int
	lastDirection   int
	loopCounter     uint8
}

func NewNavigationModule(motorController *MotorController, sensorModule *SensorModule) *NavigationModule {
	return &NavigationModule{
		motorController: motorController,
		sensorModule:    sensorModule,
		currentState:    navlogic.StateIdle,
		behaviorMode:    RANDOM_WALK_MODE,
		lastDirection:   MOVE_FORWARD,
	}
}

func (nm *NavigationModule) SetBehaviorMode(mode int) {
	nm.behaviorMode = mode
}

func (nm *NavigationModule) GetCurrentState() int {
	return nm.currentState
}

func (nm *NavigationModule) ProcessState() {
	nm.loopCounter++

	obstacleDetected := nm.sensorModule.IsObstacleDetected()
	edgeDetected := nm.sensorModule.IsEdgeDetected()

	switch nm.currentState {
	case navlogic.StateIdle:
		nm.currentState = navlogic.NextStateFromSensors(nm.currentState, obstacleDetected, edgeDetected)

	case navlogic.StateMoving:
		nextState := navlogic.NextStateFromSensors(nm.currentState, obstacleDetected, edgeDetected)
		nm.currentState = nextState
		if nextState == navlogic.StateMoving {
			if nm.behaviorMode == RANDOM_WALK_MODE && nm.loopCounter%50 == 0 {
				nm.motorController.MoveRandomly(nm.loopCounter)
			} else {
				nm.motorController.SetDirection(nm.lastDirection)
			}
		}

	case navlogic.StateObstacleAvoidance:
		nm.motorController.SetDirection(STOP)
		nm.motorController.MoveForLoops(MOVE_BACKWARD, 2500)
		if nm.loopCounter%2 == 0 {
			nm.motorController.TurnForLoops(TURN_LEFT, 3000)
		} else {
			nm.motorController.TurnForLoops(TURN_RIGHT, 3000)
		}
		nm.lastDirection = MOVE_FORWARD
		nm.currentState = navlogic.StateMoving

	case navlogic.StateEdgeAvoidance:
		nm.motorController.SetDirection(STOP)
		nm.motorController.MoveForLoops(MOVE_BACKWARD, 4000)
		if nm.loopCounter%2 == 0 {
			nm.motorController.TurnForLoops(TURN_LEFT, 4000)
		} else {
			nm.motorController.TurnForLoops(TURN_RIGHT, 4000)
		}
		nm.lastDirection = MOVE_FORWARD
		nm.currentState = navlogic.StateMoving

	case navlogic.StateInteracting:
		nm.currentState = navlogic.StateMoving
	}
}

func (nm *NavigationModule) Update() {
	nm.ProcessState()
}

func (nm *NavigationModule) EmergencyStop() {
	nm.motorController.SetDirection(STOP)
	nm.currentState = navlogic.StateIdle
}
