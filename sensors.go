package main

import (
	"machine"
	"time"
)

// Constants for sensor thresholds
const (
	OBSTACLE_DISTANCE_THRESHOLD = 20 // cm - anything closer is considered an obstacle
	EDGE_DETECTION_THRESHOLD    = 500 // analog value - below this means edge detected
)

// SensorModule handles all sensor readings
type SensorModule struct {
	ultraTrig machine.Pin
	ultraEcho machine.Pin
	irSensors map[string]machine.ADC
}

// NewSensorModule creates a new sensor module instance
func NewSensorModule(ultraTrig, ultraEcho machine.Pin, irSensors map[string]machine.ADC) *SensorModule {
	return &SensorModule{
		ultraTrig: ultraTrig,
		ultraEcho: ultraEcho,
		irSensors: irSensors,
	}
}

// ReadUltrasonicDistance measures distance using HC-SR04 ultrasonic sensor
func (s *SensorModule) ReadUltrasonicDistance() int {
	// Send trigger pulse
	s.ultraTrig.High()
	time.Sleep(time.Microsecond * 10)
	s.ultraTrig.Low()

	// Measure echo pulse duration
	start := time.Now()
	timeout := time.Now().Add(time.Millisecond * 100) // Safety timeout
	
	for !s.ultraEcho.Get() && time.Now().Before(timeout) {
		// Wait for echo to go HIGH
	}
	
	if time.Now().After(timeout) {
		return -1 // Timeout occurred
	}
	
	echoStart := time.Now()
	for s.ultraEcho.Get() && time.Now().Before(timeout) {
		// Wait for echo to go LOW
	}
	
	duration := time.Since(echoStart)
	
	// Calculate distance in cm (speed of sound = 343 m/s)
	// Distance = (duration * speed of sound) / 2
	// Convert microseconds to seconds and meters to centimeters
	distance := int((float64(duration.Microseconds()) * 343.0 / 10000.0) / 2.0)
	
	return distance
}

// ReadIRSensors reads all IR sensors to detect edges
func (s *SensorModule) ReadIRSensors() map[string]bool {
	results := make(map[string]bool)
	
	for name, sensor := range s.irSensors {
		value := sensor.Get()
		// If the value is below the threshold, we've detected an edge
		results[name] = value < EDGE_DETECTION_THRESHOLD
	}
	
	return results
}

// IsObstacleDetected checks if there's an obstacle in front
func (s *SensorModule) IsObstacleDetected() bool {
	distance := s.ReadUltrasonicDistance()
	if distance == -1 {
		// If we got a timeout, assume there's no obstacle
		return false
	}
	return distance < OBSTACLE_DISTANCE_THRESHOLD
}

// IsEdgeDetected checks if any IR sensor detects an edge
func (s *SensorModule) IsEdgeDetected() bool {
	edgeResults := s.ReadIRSensors()
	
	// Check if any of the edge sensors detect an edge
	for _, isEdge := range edgeResults {
		if isEdge {
			return true
		}
	}
	
	return false
}

// GetAllSensorData gets all sensor readings at once
func (s *SensorModule) GetAllSensorData() (int, map[string]bool) {
	distance := s.ReadUltrasonicDistance()
	edges := s.ReadIRSensors()
	
	return distance, edges
}