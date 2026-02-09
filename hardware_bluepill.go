//go:build bluepill

package main

import (
	"machine"
	"time"
)

const (
	LEFT_MOTOR_IN1     = machine.PA8
	LEFT_MOTOR_IN2     = machine.PA9
	RIGHT_MOTOR_IN1    = machine.PA10
	RIGHT_MOTOR_IN2    = machine.PA11
	ULTRA_TRIG_PIN     = machine.PA12
	ULTRA_ECHO_PIN     = machine.PB10
	IR_FRONT_LEFT_PIN  = machine.PA1
	IR_FRONT_RIGHT_PIN = machine.PA2
	DISPLAY_SDA_PIN    = machine.PB7
	DISPLAY_SCL_PIN    = machine.PB6
	STATUS_LED_PIN     = machine.PC13
	BUZZER_PIN         = machine.PB15
)

const (
	IR_FRONT_LEFT = iota
	IR_FRONT_RIGHT
	IR_SENSOR_COUNT
)

// Motor is a DC motor controlled via H-bridge (L298N IN1/IN2).
type Motor struct {
	in1 machine.Pin
	in2 machine.Pin
}

func NewMotor(in1, in2 machine.Pin) *Motor {
	motor := &Motor{in1: in1, in2: in2}
	motor.in1.Configure(machine.PinConfig{Mode: machine.PinOutput})
	motor.in2.Configure(machine.PinConfig{Mode: machine.PinOutput})
	return motor
}

func (m *Motor) Forward() {
	m.in1.High()
	m.in2.Low()
}

func (m *Motor) Backward() {
	m.in1.Low()
	m.in2.High()
}

func (m *Motor) Stop() {
	m.in1.Low()
	m.in2.Low()
}

// Robot holds pins and drivers for the desk pet hardware.
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

func (r *Robot) BlinkLED(times int) {
	for i := 0; i < times; i++ {
		r.statusLed.High()
		time.Sleep(time.Millisecond * 200)
		r.statusLed.Low()
		time.Sleep(time.Millisecond * 200)
	}
}

func (r *Robot) Beep(duration time.Duration) {
	r.buzzer.High()
	time.Sleep(duration)
	r.buzzer.Low()
}

func (r *Robot) BeepLoops(loops int) {
	r.buzzer.High()
	for i := 0; i < loops; i++ {
	}
	r.buzzer.Low()
}

func (r *Robot) Initialize() {
	r.BlinkLED(3)
	r.Beep(time.Millisecond * 500)
}
