package main

import (
	"time"

	"github.com/GyeongHoKim/tiny-pet/internal/navlogic"
)

// Robot states (re-export navlogic state constants for use in main and other packages)
const (
	IDLE_STATE                  = navlogic.StateIdle
	MOVING_STATE                = navlogic.StateMoving
	OBSTACLE_AVOIDANCE_STATE    = navlogic.StateObstacleAvoidance
	EDGE_AVOIDANCE_STATE        = navlogic.StateEdgeAvoidance
	INTERACTING_STATE           = navlogic.StateInteracting
)

// Behavior modes
const (
	RANDOM_WALK_MODE = iota
	GUARD_MODE
	INTERACTIVE_MODE
)

// NavigationModule handles the robot's decision-making and navigation
type NavigationModule struct {
	motorController *MotorController
	sensorModule    *SensorModule
	currentState    int
	behaviorMode    int
	lastDirection   int
	obstacleCount   int
	edgeCount       int
}

// NewNavigationModule creates a new navigation module instance
func NewNavigationModule(motorController *MotorController, sensorModule *SensorModule) *NavigationModule {
	return &NavigationModule{
		motorController: motorController,
		sensorModule:    sensorModule,
		currentState:    navlogic.StateIdle,
		behaviorMode:    RANDOM_WALK_MODE,
		lastDirection:   MOVE_FORWARD,
		obstacleCount:   0,
		edgeCount:       0,
	}
}

// SetBehaviorMode sets the robot's behavior mode
func (nm *NavigationModule) SetBehaviorMode(mode int) {
	nm.behaviorMode = mode
}

// GetCurrentState returns the current state
func (nm *NavigationModule) GetCurrentState() int {
	return nm.currentState
}

// ProcessState processes the current state and updates as needed
func (nm *NavigationModule) ProcessState() {
	// Read sensor data
	obstacleDetected := nm.sensorModule.IsObstacleDetected()
	edgeDetected := nm.sensorModule.IsEdgeDetected()
	
	// Update state based on sensor data and current state
	switch nm.currentState {
	case navlogic.StateIdle:
		nm.currentState = navlogic.NextStateFromSensors(nm.currentState, obstacleDetected, edgeDetected)

	case navlogic.StateMoving:
		nextState := navlogic.NextStateFromSensors(nm.currentState, obstacleDetected, edgeDetected)
		nm.currentState = nextState
		if nextState == navlogic.StateEdgeAvoidance {
			nm.edgeCount++
		} else if nextState == navlogic.StateObstacleAvoidance {
			nm.obstacleCount++
		} else {
			// Still moving: continue in current direction or change randomly
			if nm.behaviorMode == RANDOM_WALK_MODE && time.Now().UnixNano()%5 == 0 {
				nm.motorController.MoveRandomly()
			} else {
				nm.motorController.SetDirection(nm.lastDirection)
			}
		}

	case navlogic.StateObstacleAvoidance:
		// Stop motors
		nm.motorController.SetDirection(STOP)
		
		// Move backward to create space
		nm.motorController.MoveForTime(MOVE_BACKWARD, time.Millisecond*300)
		
		// Turn away from obstacle
		// For now, choose a random turn direction
		if time.Now().UnixNano()%2 == 0 {
			nm.motorController.TurnForTime(TURN_LEFT, time.Millisecond*400)
		} else {
			nm.motorController.TurnForTime(TURN_RIGHT, time.Millisecond*400)
		}
		
		// Reset counter and prefer forward after avoidance
		nm.obstacleCount = 0
		nm.lastDirection = MOVE_FORWARD
		
		// Return to moving state
		nm.currentState = navlogic.StateMoving

	case navlogic.StateEdgeAvoidance:
		// Stop motors
		nm.motorController.SetDirection(STOP)
		
		// Move backward to get away from edge
		nm.motorController.MoveForTime(MOVE_BACKWARD, time.Millisecond*500)
		
		// Turn toward the center of the desk
		// For now, choose a random turn direction
		if time.Now().UnixNano()%2 == 0 {
			nm.motorController.TurnForTime(TURN_LEFT, time.Millisecond*600)
		} else {
			nm.motorController.TurnForTime(TURN_RIGHT, time.Millisecond*600)
		}
		
		// Reset counter and prefer forward after avoidance
		nm.edgeCount = 0
		nm.lastDirection = MOVE_FORWARD
		
		// Return to moving state
		nm.currentState = navlogic.StateMoving

	case navlogic.StateInteracting:
		// Handle interaction behavior
		// For now, just blink lights or make sounds
		// This would be triggered by proximity sensors or buttons
		nm.currentState = navlogic.StateMoving
	}
}

// Update runs one cycle of the navigation logic
func (nm *NavigationModule) Update() {
	nm.ProcessState()
}

// EmergencyStop stops the robot immediately
func (nm *NavigationModule) EmergencyStop() {
	nm.motorController.SetDirection(STOP)
	nm.currentState = navlogic.StateIdle
}

// GetStateName returns a string representation of the current state
func (nm *NavigationModule) GetStateName() string {
	switch nm.currentState {
	case navlogic.StateIdle:
		return "IDLE"
	case navlogic.StateMoving:
		return "MOVING"
	case navlogic.StateObstacleAvoidance:
		return "AVOIDING OBSTACLE"
	case navlogic.StateEdgeAvoidance:
		return "AVOIDING EDGE"
	case navlogic.StateInteracting:
		return "INTERACTING"
	default:
		return "UNKNOWN"
	}
}

// GetBehaviorModeName returns a string representation of the current behavior mode
func (nm *NavigationModule) GetBehaviorModeName() string {
	switch nm.behaviorMode {
	case RANDOM_WALK_MODE:
		return "RANDOM WALK"
	case GUARD_MODE:
		return "GUARD"
	case INTERACTIVE_MODE:
		return "INTERACTIVE"
	default:
		return "UNKNOWN"
	}
}

// ChangeDirection changes the robot's movement direction
func (nm *NavigationModule) ChangeDirection() {
	// Store the previous direction
	prevDirection := nm.lastDirection
	
	// Choose a new direction that's not the opposite of the previous
	// to avoid oscillating back and forth
	var newDirection int
	for {
		newDirection = int(time.Now().UnixNano() % 4) // Random direction
		// Avoid choosing the opposite direction immediately
		if (prevDirection == MOVE_FORWARD && newDirection != MOVE_BACKWARD) ||
		   (prevDirection == MOVE_BACKWARD && newDirection != MOVE_FORWARD) ||
		   (prevDirection == TURN_LEFT && newDirection != TURN_RIGHT) ||
		   (prevDirection == TURN_RIGHT && newDirection != TURN_LEFT) ||
		   prevDirection == STOP {
			break
		}
	}
	
	nm.lastDirection = newDirection
	nm.motorController.SetDirection(newDirection)
}