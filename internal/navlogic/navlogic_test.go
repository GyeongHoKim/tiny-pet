package navlogic

import "testing"

func TestNextStateFromSensors_Idle(t *testing.T) {
	tests := []struct {
		obstacle bool
		edge     bool
		want     int
	}{
		{false, false, StateMoving},
		{true, false, StateObstacleAvoidance},
		{false, true, StateEdgeAvoidance},
		{true, true, StateEdgeAvoidance}, // edge takes precedence
	}
	for _, tt := range tests {
		got := NextStateFromSensors(StateIdle, tt.obstacle, tt.edge)
		if got != tt.want {
			t.Errorf("NextStateFromSensors(StateIdle, obstacle=%v, edge=%v) = %v, want %v",
				tt.obstacle, tt.edge, got, tt.want)
		}
	}
}

func TestNextStateFromSensors_Moving(t *testing.T) {
	tests := []struct {
		obstacle bool
		edge     bool
		want     int
	}{
		{false, false, StateMoving},
		{true, false, StateObstacleAvoidance},
		{false, true, StateEdgeAvoidance},
		{true, true, StateEdgeAvoidance}, // edge takes precedence
	}
	for _, tt := range tests {
		got := NextStateFromSensors(StateMoving, tt.obstacle, tt.edge)
		if got != tt.want {
			t.Errorf("NextStateFromSensors(StateMoving, obstacle=%v, edge=%v) = %v, want %v",
				tt.obstacle, tt.edge, got, tt.want)
		}
	}
}

func TestNextStateFromSensors_OtherStatesUnchanged(t *testing.T) {
	otherStates := []int{StateObstacleAvoidance, StateEdgeAvoidance, StateInteracting}
	for _, s := range otherStates {
		got := NextStateFromSensors(s, true, true)
		if got != s {
			t.Errorf("NextStateFromSensors(%v, true, true) = %v, want %v (unchanged)", s, got, s)
		}
	}
}
