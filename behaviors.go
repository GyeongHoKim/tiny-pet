package main

import (
	"machine"
	"time"
)

// BehaviorPatterns provides LED and buzzer feedback for robot states.
type BehaviorPatterns struct {
	statusLed machine.Pin
	buzzer    machine.Pin
}

// NewBehaviorPatterns creates a BehaviorPatterns with the given pins.
func NewBehaviorPatterns(statusLed, buzzer machine.Pin) *BehaviorPatterns {
	return &BehaviorPatterns{
		statusLed: statusLed,
		buzzer:    buzzer,
	}
}

// IndicateStateChange blinks the LED briefly to signal a state transition.
func (bp *BehaviorPatterns) IndicateStateChange(_ int) {
	bp.statusLed.High()
	time.Sleep(time.Millisecond * 80)
	bp.statusLed.Low()
}
