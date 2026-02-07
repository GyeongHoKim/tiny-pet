package main

import (
	"machine"
	"time"
)

// Pin configurations for the Tiny Pet robot
const (
	// Motor pins
	LEFT_MOTOR_PIN  = machine.D5
	RIGHT_MOTOR_PIN = machine.D6
	
	// Ultrasonic sensor pins (for obstacle detection)
	ULTRA_TRIG_PIN = machine.D7
	ULTRA_ECHO_PIN = machine.A0
	
	// IR sensor pins (for edge detection)
	IR_FRONT_LEFT_PIN  = machine.A1
	IR_FRONT_RIGHT_PIN = machine.A2
	IR_REAR_LEFT_PIN   = machine.A3
	IR_REAR_RIGHT_PIN  = machine.A4
	
	// LED and buzzer pins
	STATUS_LED_PIN = machine.D13
	BUZZER_PIN     = machine.D8
)

// Motor represents a motor with control pin
type Motor struct {
	pin machine.Pin
}

// NewMotor creates a new motor instance
func NewMotor(pin machine.Pin) *Motor {
	motor := &Motor{pin: pin}
	motor.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return motor
}

// SetSpeed sets the motor speed (HIGH/LOW for simplicity, can be extended for PWM)
func (m *Motor) SetSpeed(speed bool) {
	if speed {
		m.pin.High()
	} else {
		m.pin.Low()
	}
}

// Stop stops the motor
func (m *Motor) Stop() {
	m.pin.Low()
}

// Robot represents the Tiny Pet robot
type Robot struct {
	leftMotor  *Motor
	rightMotor *Motor
	ultraTrig  machine.Pin
	ultraEcho  machine.Pin
	irSensors  map[string]machine.ADC
	statusLed  machine.Pin
	buzzer     machine.Pin
}

// NewRobot creates a new robot instance with initialized hardware
func NewRobot() *Robot {
	robot := &Robot{
		leftMotor:  NewMotor(LEFT_MOTOR_PIN),
		rightMotor: NewMotor(RIGHT_MOTOR_PIN),
		ultraTrig:  ULTRA_TRIG_PIN,
		ultraEcho:  ULTRA_ECHO_PIN,
		irSensors:  make(map[string]machine.ADC),
		statusLed:  STATUS_LED_PIN,
		buzzer:     BUZZER_PIN,
	}
	
	// Configure ultrasonic sensor pins
	robot.ultraTrig.Configure(machine.PinConfig{Mode: machine.PinOutput})
	robot.ultraEcho.Configure(machine.PinConfig{Mode: machine.PinInput})
	
	// Configure IR sensor pins
	irPins := map[string]machine.Pin{
		"front_left":  IR_FRONT_LEFT_PIN,
		"front_right": IR_FRONT_RIGHT_PIN,
		"rear_left":   IR_REAR_LEFT_PIN,
		"rear_right":  IR_REAR_RIGHT_PIN,
	}
	
	for name, pin := range irPins {
		adc := machine.ADC{pin}
		adc.Configure()
		robot.irSensors[name] = adc
	}
	
	// Configure LED and buzzer pins
	robot.statusLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	robot.buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})
	
	return robot
}

// BlinkLED blinks the status LED a specified number of times
func (r *Robot) BlinkLED(times int) {
	for i := 0; i < times; i++ {
		r.statusLed.High()
		time.Sleep(time.Millisecond * 200)
		r.statusLed.Low()
		time.Sleep(time.Millisecond * 200)
	}
}

// Beep makes the buzzer beep
func (r *Robot) Beep(duration time.Duration) {
	r.buzzer.High()
	time.Sleep(duration)
	r.buzzer.Low()
}

// Initialize the robot and show startup sequence
func (r *Robot) Initialize() {
	r.BlinkLED(3) // Blink 3 times to indicate startup
	r.Beep(time.Millisecond * 500) // Beep to indicate ready
}