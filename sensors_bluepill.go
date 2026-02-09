//go:build bluepill

package main

import (
	"machine"
	"time"

	"github.com/GyeongHoKim/tiny-pet/internal/navlogic"
)

const (
	OBSTACLE_DISTANCE_THRESHOLD = 20
	EDGE_DETECTION_THRESHOLD    = 500
)

const bluepillLoopsPerMicrosecond = 4
const bluepillUltrasonicTimeoutLoops = 50000

// SensorModule reads ultrasonic (HC-SR04) and IR edge sensors.
type SensorModule struct {
	ultraTrig machine.Pin
	ultraEcho machine.Pin
	irSensors *[IR_SENSOR_COUNT]machine.ADC
}

func NewSensorModule(ultraTrig, ultraEcho machine.Pin, irSensors *[IR_SENSOR_COUNT]machine.ADC) *SensorModule {
	return &SensorModule{
		ultraTrig: ultraTrig,
		ultraEcho: ultraEcho,
		irSensors: irSensors,
	}
}

// ReadUltrasonicDistance returns distance in cm, or -1 on timeout.
func (s *SensorModule) ReadUltrasonicDistance() int {
	s.ultraTrig.High()
	time.Sleep(10 * time.Microsecond)
	s.ultraTrig.Low()

	count := 0
	for !s.ultraEcho.Get() {
		count++
		if count > bluepillUltrasonicTimeoutLoops {
			return -1
		}
	}

	echoCount := 0
	for s.ultraEcho.Get() {
		echoCount++
		if echoCount > bluepillUltrasonicTimeoutLoops {
			return -1
		}
	}

	us := echoCount / bluepillLoopsPerMicrosecond
	return navlogic.EchoMicrosecondsToDistanceCm(us)
}

func (s *SensorModule) IsObstacleDetected() bool {
	distance := s.ReadUltrasonicDistance()
	return navlogic.IsWithinThreshold(distance, OBSTACLE_DISTANCE_THRESHOLD)
}

func (s *SensorModule) IsEdgeDetected() bool {
	for i := 0; i < IR_SENSOR_COUNT; i++ {
		if s.irSensors[i].Get() < EDGE_DETECTION_THRESHOLD {
			return true
		}
	}
	return false
}

func (s *SensorModule) ReadIRSensors() [IR_SENSOR_COUNT]bool {
	var results [IR_SENSOR_COUNT]bool
	for i := 0; i < IR_SENSOR_COUNT; i++ {
		results[i] = s.irSensors[i].Get() < EDGE_DETECTION_THRESHOLD
	}
	return results
}
