package main

import (
	"time"
)

func main() {
	// Initialize the robot hardware
	robot := NewRobot()
	
	// Initialize all modules
	sensorModule := NewSensorModule(robot.ultraTrig, robot.ultraEcho, robot.irSensors)
	motorController := NewMotorController(robot.leftMotor, robot.rightMotor)
	navigationModule := NewNavigationModule(motorController, sensorModule)
	behaviorPatterns := NewBehaviorPatterns(robot.statusLed, robot.buzzer)
	calibrationModule := NewCalibrationModule(robot, sensorModule, motorController)
	
	// Initialize the robot (blink lights, beep to indicate ready)
	robot.Initialize()
	
	// Perform initial calibration
	calibrationModule.CalibrateComplete()
	
	// Set initial behavior mode
	navigationModule.SetBehaviorMode(RANDOM_WALK_MODE)
	
	var lastStateName string
	// Main loop
	for {
		// Update navigation logic
		navigationModule.Update()
		
		// Only indicate with LED when state has actually changed
		currentStateName := navigationModule.GetStateName()
		if currentStateName != lastStateName {
			behaviorPatterns.IndicateStateChange(currentStateName)
			lastStateName = currentStateName
		}
		
		// Occasionally perform special behaviors
		if time.Now().UnixNano()%5000 == 0 {
			behaviorPatterns.HeartbeatPattern()
		}
		
		// Small delay to prevent overwhelming the processor
		time.Sleep(time.Millisecond * 100)
	}
}
