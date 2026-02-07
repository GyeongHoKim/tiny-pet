package main

import (
	"machine"

	"github.com/GyeongHoKim/tiny-pet/internal/navlogic"
)

// Sensor thresholds.
const (
	OBSTACLE_DISTANCE_THRESHOLD = 20
	EDGE_DETECTION_THRESHOLD    = 500
	ULTRASONIC_TIMEOUT_LOOPS    = 10000
)

// SensorModule handles ultrasonic and IR sensor readings.
type SensorModule struct {
	ultraTrig machine.Pin
	ultraEcho machine.Pin
	irSensors *[IR_SENSOR_COUNT]machine.ADC
}

// NewSensorModule creates a SensorModule with the given pins.
func NewSensorModule(ultraTrig, ultraEcho machine.Pin, irSensors *[IR_SENSOR_COUNT]machine.ADC) *SensorModule {
	return &SensorModule{
		ultraTrig: ultraTrig,
		ultraEcho: ultraEcho,
		irSensors: irSensors,
	}
}

// ReadUltrasonicDistance measures distance using HC-SR04.
// Returns distance in cm, or -1 on timeout.
func (s *SensorModule) ReadUltrasonicDistance() int {
	s.ultraTrig.High()
	for i := 0; i < 160; i++ {
	}
	s.ultraTrig.Low()

	count := 0
	for !s.ultraEcho.Get() {
		count++
		if count > ULTRASONIC_TIMEOUT_LOOPS {
			return -1
		}
	}

	echoCount := 0
	for s.ultraEcho.Get() {
		echoCount++
		if echoCount > ULTRASONIC_TIMEOUT_LOOPS {
			return -1
		}
	}

	return navlogic.EchoCountToDistanceCm(echoCount)
}

// IsObstacleDetected returns true if an obstacle is within threshold distance.
func (s *SensorModule) IsObstacleDetected() bool {
	distance := s.ReadUltrasonicDistance()
	return navlogic.IsWithinThreshold(distance, OBSTACLE_DISTANCE_THRESHOLD)
}

// IsEdgeDetected returns true if any IR sensor detects an edge.
func (s *SensorModule) IsEdgeDetected() bool {
	for i := 0; i < IR_SENSOR_COUNT; i++ {
		if s.irSensors[i].Get() < EDGE_DETECTION_THRESHOLD {
			return true
		}
	}
	return false
}

// ReadIRSensors returns edge detection status for each IR sensor.
func (s *SensorModule) ReadIRSensors() [IR_SENSOR_COUNT]bool {
	var results [IR_SENSOR_COUNT]bool
	for i := 0; i < IR_SENSOR_COUNT; i++ {
		results[i] = s.irSensors[i].Get() < EDGE_DETECTION_THRESHOLD
	}
	return results
}
