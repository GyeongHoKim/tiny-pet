package main

import (
	"machine"
	"time"
)

type BehaviorPatterns struct {
	statusLed machine.Pin
	buzzer    machine.Pin
}

func NewBehaviorPatterns(statusLed, buzzer machine.Pin) *BehaviorPatterns {
	return &BehaviorPatterns{
		statusLed: statusLed,
		buzzer:    buzzer,
	}
}

func (bp *BehaviorPatterns) IndicateStateChange(_ int) {
	bp.statusLed.High()
	time.Sleep(time.Millisecond * 80)
	bp.statusLed.Low()
}
