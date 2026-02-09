package main

import (
	"machine"
	"time"
)

// Pin constants for Arduino Uno/Nano.
const (
	LEFT_MOTOR_IN1     = machine.D5 // Mini L298N IN1 (left)
	LEFT_MOTOR_IN2     = machine.D4 // Mini L298N IN2 (left)
	RIGHT_MOTOR_IN1    = machine.D6 // Mini L298N IN3 (right)
	RIGHT_MOTOR_IN2    = machine.D9 // Mini L298N IN4 (right)
	ULTRA_TRIG_PIN     = machine.D7
	ULTRA_ECHO_PIN     = machine.ADC0
	IR_FRONT_LEFT_PIN  = machine.ADC1
	IR_FRONT_RIGHT_PIN = machine.ADC2
	DISPLAY_SDA_PIN    = machine.ADC4 // I2C SDA (shared with hardware I2C0)
	DISPLAY_SCL_PIN    = machine.ADC5 // I2C SCL (shared with hardware I2C0)
	STATUS_LED_PIN     = machine.D13
	BUZZER_PIN         = machine.D8
)

// IR sensor indices.
const (
	IR_FRONT_LEFT = iota
	IR_FRONT_RIGHT
	IR_SENSOR_COUNT
)

// Motor represents a DC motor controlled via H-bridge (Mini L298N IN1/IN2).
type Motor struct {
	in1 machine.Pin
	in2 machine.Pin
}

// NewMotor creates and configures a motor with two H-bridge control pins.
func NewMotor(in1, in2 machine.Pin) *Motor {
	motor := &Motor{in1: in1, in2: in2}
	motor.in1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	motor.in2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return motor
}

// Forward drives the motor forward (in1 HIGH, in2 LOW).
func (m *Motor) Forward() {
	m.in1.High()
	m.in2.Low()
}

// Backward drives the motor in reverse (in1 LOW, in2 HIGH).
func (m *Motor) Backward() {
	m.in1.Low()
	m.in2.High()
}

// Stop brakes the motor (in1 LOW, in2 LOW).
func (m *Motor) Stop() {
	m.in1.Low()
	m.in2.Low()
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
		leftMotor:  NewMotor(LEFT_MOTOR_IN1, LEFT_MOTOR_IN2),
		rightMotor: NewMotor(RIGHT_MOTOR_IN1, RIGHT_MOTOR_IN2),
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
