package main

import (
	"machine"
	"time"
)

// CalibrationModule handles the calibration of sensors and movement parameters
type CalibrationModule struct {
	robot           *Robot
	sensorModule    *SensorModule
	motorController *MotorController
	calibrated      bool
}

// NewCalibrationModule creates a new calibration module instance
func NewCalibrationModule(robot *Robot, sensorModule *SensorModule, motorController *MotorController) *CalibrationModule {
	return &CalibrationModule{
		robot:           robot,
		sensorModule:    sensorModule,
		motorController: motorController,
		calibrated:      false,
	}
}

// CalibrateSensors runs the sensor calibration process
func (cm *CalibrationModule) CalibrateSensors() {
	cm.robot.BlinkLED(2) // Indicate calibration start
	
	println("Starting sensor calibration...")
	
	// Calibrate ultrasonic sensor - measure baseline distance
	println("Measuring baseline distance...")
	for i := 0; i < 10; i++ {
		distance := cm.sensorModule.ReadUltrasonicDistance()
		println("Ultrasonic distance:", distance)
		time.Sleep(time.Millisecond * 500)
	}
	
	// Calibrate IR sensors - measure baseline values over different surfaces
	println("Measuring IR sensor baselines...")
	for i := 0; i < 10; i++ {
		irValues := cm.sensorModule.ReadIRSensors()
		for sensorName, value := range irValues {
			println("IR sensor", sensorName, "value:", value)
		}
		time.Sleep(time.Millisecond * 500)
	}
	
	println("Sensor calibration complete!")
	cm.robot.Beep(time.Millisecond * 1000) // Long beep to indicate completion
	cm.calibrated = true
}

// CalibrateMotors runs the motor calibration process
func (cm *CalibrationModule) CalibrateMotors() {
	cm.robot.BlinkLED(3) // Indicate motor calibration start
	
	println("Starting motor calibration...")
	
	// Test each movement direction
	movements := []struct {
		name string
		dir  int
	}{
		{"Forward", MOVE_FORWARD},
		{"Backward", MOVE_BACKWARD},
		{"Left", TURN_LEFT},
		{"Right", TURN_RIGHT},
	}
	
	for _, move := range movements {
		println("Testing", move.name, "movement...")
		cm.motorController.SetDirection(move.dir)
		time.Sleep(time.Millisecond * 1000)
		cm.motorController.SetDirection(STOP)
		time.Sleep(time.Millisecond * 500)
	}
	
	println("Motor calibration complete!")
	cm.robot.Beep(time.Millisecond * 500) // Medium beep
}

// CalibrateComplete runs full calibration
func (cm *CalibrationModule) CalibrateComplete() {
	println("Starting complete calibration...")
	
	// Calibrate sensors first
	cm.CalibrateSensors()
	
	// Then calibrate motors
	cm.CalibrateMotors()
	
	println("Complete calibration finished!")
	cm.robot.BlinkLED(5) // Indicate completion
}

// IsCalibrated returns whether the system has been calibrated
func (cm *CalibrationModule) IsCalibrated() bool {
	return cm.calibrated
}

// AdjustThresholds allows runtime adjustment of sensor thresholds
func (cm *CalibrationModule) AdjustThresholds(obstacleDist, edgeDetect int) {
	// In a real implementation, we would adjust global constants
	// For now, we'll just print the new values
	println("Adjusting thresholds - obstacle distance:", obstacleDist, "edge detection:", edgeDetect)
}