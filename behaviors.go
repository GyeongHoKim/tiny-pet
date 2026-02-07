package main

import (
	"machine"
	"time"
)

// BehaviorPatterns defines different behavioral patterns for the robot
type BehaviorPatterns struct {
	statusLed machine.Pin
	buzzer    machine.Pin
}

// NewBehaviorPatterns creates a new behavior patterns instance
func NewBehaviorPatterns(statusLed, buzzer machine.Pin) *BehaviorPatterns {
	return &BehaviorPatterns{
		statusLed: statusLed,
		buzzer:    buzzer,
	}
}

// IndicateStateChange indicates a state change with a short LED pattern.
// Kept brief to avoid blocking the main loop and navigation.
func (bp *BehaviorPatterns) IndicateStateChange(stateName string) {
	switch stateName {
	case "MOVING":
		bp.statusLed.High()
		time.Sleep(time.Millisecond * 80)
		bp.statusLed.Low()
	case "AVOIDING OBSTACLE":
		for i := 0; i < 3; i++ {
			bp.statusLed.High()
			time.Sleep(time.Millisecond * 50)
			bp.statusLed.Low()
			time.Sleep(time.Millisecond * 50)
		}
	case "AVOIDING EDGE":
		bp.statusLed.High()
		bp.PlayTone(523, time.Millisecond*50)
		time.Sleep(time.Millisecond * 50)
		bp.statusLed.Low()
		bp.PlayTone(392, time.Millisecond*50)
		time.Sleep(time.Millisecond * 50)
	case "INTERACTING":
		bp.statusLed.High()
		bp.PlayTone(659, time.Millisecond*80)
		time.Sleep(time.Millisecond * 80)
		bp.statusLed.Low()
	default:
		bp.statusLed.High()
		time.Sleep(time.Millisecond * 100)
		bp.statusLed.Low()
	}
}

// PlayTone plays a tone on the buzzer for a specified duration
func (bp *BehaviorPatterns) PlayTone(frequency int, duration time.Duration) {
	// For simplicity in this implementation, we'll just turn the buzzer on/off
	// A more sophisticated implementation would use PWM to generate actual tones
	bp.buzzer.High()
	time.Sleep(duration)
	bp.buzzer.Low()
}

// HappyPattern plays a happy pattern when robot is functioning well
func (bp *BehaviorPatterns) HappyPattern() {
	for i := 0; i < 3; i++ {
		bp.statusLed.High()
		bp.PlayTone(523, time.Millisecond*100) // C5
		time.Sleep(time.Millisecond * 100)
		
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 100)
		
		bp.statusLed.High()
		bp.PlayTone(659, time.Millisecond*100) // E5
		time.Sleep(time.Millisecond * 100)
		
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 100)
		
		bp.statusLed.High()
		bp.PlayTone(784, time.Millisecond*150) // G5
		time.Sleep(time.Millisecond * 150)
		
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 200)
	}
}

// AlertPattern plays an alert pattern when robot encounters an issue
func (bp *BehaviorPatterns) AlertPattern() {
	for i := 0; i < 5; i++ {
		bp.statusLed.High()
		bp.PlayTone(220, time.Millisecond*150) // A3
		time.Sleep(time.Millisecond * 150)
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 150)
	}
}

// SleepPattern indicates the robot is entering low-power mode
func (bp *BehaviorPatterns) SleepPattern() {
	// Slow fade out pattern
	for i := 0; i < 3; i++ {
		bp.statusLed.High()
		time.Sleep(time.Millisecond * 300)
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 300)
	}
	
	// Short buzz
	bp.PlayTone(220, time.Millisecond*300)
}

// WakeUpPattern indicates the robot is exiting low-power mode
func (bp *BehaviorPatterns) WakeUpPattern() {
	// Short buzz
	bp.PlayTone(784, time.Millisecond*300)
	
	// Quick blink sequence
	for i := 0; i < 5; i++ {
		bp.statusLed.High()
		time.Sleep(time.Millisecond * 100)
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 100)
	}
}

// RespondToUser responds to user interaction
func (bp *BehaviorPatterns) RespondToUser() {
	// Play a cheerful pattern when user interacts
	for i := 0; i < 2; i++ {
		bp.statusLed.High()
		bp.PlayTone(659, time.Millisecond*150) // E5
		time.Sleep(time.Millisecond * 150)
		
		bp.statusLed.Low()
		bp.PlayTone(784, time.Millisecond*150) // G5
		time.Sleep(time.Millisecond * 150)
		
		bp.statusLed.High()
		bp.PlayTone(1047, time.Millisecond*200) // C6
		time.Sleep(time.Millisecond * 200)
		
		bp.statusLed.Low()
		time.Sleep(time.Millisecond * 300)
	}
}

// HeartbeatPattern simulates a heartbeat pattern
func (bp *BehaviorPatterns) HeartbeatPattern() {
	// Quick double blink like a heartbeat
	bp.statusLed.High()
	bp.PlayTone(523, time.Millisecond*100) // C5
	time.Sleep(time.Millisecond * 100)
	bp.statusLed.Low()
	time.Sleep(time.Millisecond * 100)
	
	bp.statusLed.High()
	bp.PlayTone(659, time.Millisecond*150) // E5
	time.Sleep(time.Millisecond * 150)
	bp.statusLed.Low()
	time.Sleep(time.Millisecond * 500)
}