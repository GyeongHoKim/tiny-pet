package main

import (
	"machine"
	"time"
)

// Pin constants for Arduino Uno/Nano.
const (
	LEFT_MOTOR_PIN     = machine.D5
	RIGHT_MOTOR_PIN    = machine.D6
	ULTRA_TRIG_PIN     = machine.D7
	ULTRA_ECHO_PIN     = machine.ADC0
	IR_FRONT_LEFT_PIN  = machine.ADC1
	IR_FRONT_RIGHT_PIN = machine.ADC2
	IR_REAR_LEFT_PIN   = machine.ADC3
	IR_REAR_RIGHT_PIN  = machine.ADC4
	STATUS_LED_PIN     = machine.D13
	BUZZER_PIN         = machine.D8
)

// IR sensor indices.
const (
	IR_FRONT_LEFT = iota
	IR_FRONT_RIGHT
	IR_REAR_LEFT
	IR_REAR_RIGHT
	IR_SENSOR_COUNT
)

// Motor represents a DC motor controlled by a single pin.
type Motor struct {
	pin machine.Pin
}

// NewMotor creates and configures a motor on the given pin.
func NewMotor(pin machine.Pin) *Motor {
	motor := &Motor{pin: pin}
	motor.pin.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return motor
}

// SetSpeed sets the motor output (true=HIGH, false=LOW).
func (m *Motor) SetSpeed(speed bool) {
	if speed {
		m.pin.High()
	} else {
		m.pin.Low()
	}
}

// Stop sets the motor output to LOW.
func (m *Motor) Stop() {
	m.pin.Low()
}

// Robot represents the Tiny Pet hardware.
type Robot struct {
	leftMotor  *Motor
	rightMotor *Motor
	ultraTrig  machine.Pin
	ultraEcho  machine.Pin
	irSensors  [IR_SENSOR_COUNT]machine.ADC
	statusLed  machine.Pin
	buzzer     machine.Pin
}

// NewRobot initializes all hardware and returns a configured Robot.
func NewRobot() *Robot {
	robot := &Robot{
		leftMotor:  NewMotor(LEFT_MOTOR_PIN),
		rightMotor: NewMotor(RIGHT_MOTOR_PIN),
		ultraTrig:  ULTRA_TRIG_PIN,
		ultraEcho:  ULTRA_ECHO_PIN,
		statusLed:  STATUS_LED_PIN,
		buzzer:     BUZZER_PIN,
	}

	robot.ultraTrig.Configure(machine.PinConfig{Mode: machine.PinOutput})
	robot.ultraEcho.Configure(machine.PinConfig{Mode: machine.PinInput})

	machine.InitADC()

	irPins := [IR_SENSOR_COUNT]machine.Pin{
		IR_FRONT_LEFT_PIN,
		IR_FRONT_RIGHT_PIN,
		IR_REAR_LEFT_PIN,
		IR_REAR_RIGHT_PIN,
	}
	for i, pin := range irPins {
		robot.irSensors[i] = machine.ADC{Pin: pin}
		robot.irSensors[i].Configure(machine.ADCConfig{})
	}

	robot.statusLed.Configure(machine.PinConfig{Mode: machine.PinOutput})
	robot.buzzer.Configure(machine.PinConfig{Mode: machine.PinOutput})

	return robot
}

// BlinkLED blinks the status LED the specified number of times.
func (r *Robot) BlinkLED(times int) {
	for i := 0; i < times; i++ {
		r.statusLed.High()
		time.Sleep(time.Millisecond * 200)
		r.statusLed.Low()
		time.Sleep(time.Millisecond * 200)
	}
}

// Beep activates the buzzer for the specified duration.
func (r *Robot) Beep(duration time.Duration) {
	r.buzzer.High()
	time.Sleep(duration)
	r.buzzer.Low()
}

// BeepLoops activates the buzzer using busy-wait timing.
func (r *Robot) BeepLoops(loops int) {
	r.buzzer.High()
	for i := 0; i < loops; i++ {
	}
	r.buzzer.Low()
}

// Initialize runs the startup sequence (LED blinks and beep).
func (r *Robot) Initialize() {
	r.BlinkLED(3)
	r.Beep(time.Millisecond * 500)
}
