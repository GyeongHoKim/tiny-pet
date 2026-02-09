// Package navlogic implements navigation state transitions with no hardware dependencies (unit-testable with go test).
package navlogic

const (
	StateIdle = iota
	StateMoving
	StateObstacleAvoidance
	StateEdgeAvoidance
	StateInteracting
)

// NextStateFromSensors returns the next state from sensor inputs; edge takes precedence over obstacle.
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
