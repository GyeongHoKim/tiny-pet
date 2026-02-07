// Package navlogic implements pure state transition logic for the Tiny Pet
// navigation state machine. It has no dependency on machine or hardware, so
// it can be unit-tested with the standard Go toolchain (go test).
package navlogic

// Robot navigation states. Values must stay in sync with navigation.go usage.
const (
	StateIdle = iota
	StateMoving
	StateObstacleAvoidance
	StateEdgeAvoidance
	StateInteracting
)

// NextStateFromSensors returns the next navigation state based on the current
// state and sensor readings. It encodes only the sensor-driven transitions
// from IDLE and MOVING; avoidance states transition to MOVING after physical
// actions (handled in the main package).
// Edge takes precedence over obstacle when both are detected.
func NextStateFromSensors(currentState int, obstacleDetected, edgeDetected bool) int {
	switch currentState {
	case StateIdle:
		if edgeDetected {
			return StateEdgeAvoidance
		}
		if obstacleDetected {
			return StateObstacleAvoidance
		}
		return StateMoving

	case StateMoving:
		if edgeDetected {
			return StateEdgeAvoidance
		}
		if obstacleDetected {
			return StateObstacleAvoidance
		}
		return StateMoving

	default:
		return currentState
	}
}
