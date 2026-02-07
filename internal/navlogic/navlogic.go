// Package navlogic implements pure state transition logic for navigation.
// It has no hardware dependencies and can be unit-tested with go test.
package navlogic

// Navigation states.
const (
	StateIdle = iota
	StateMoving
	StateObstacleAvoidance
	StateEdgeAvoidance
	StateInteracting
)

// NextStateFromSensors returns the next state based on sensor readings.
// Edge detection takes precedence over obstacle detection.
func NextStateFromSensors(currentState int, obstacleDetected, edgeDetected bool) int {
	switch currentState {
	case StateIdle, StateMoving:
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
