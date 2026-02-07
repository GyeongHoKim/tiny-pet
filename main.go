package main

import (
	"time"
)

func main() {
	robot := NewRobot()

	sensorModule := NewSensorModule(robot.ultraTrig, robot.ultraEcho, &robot.irSensors)
	motorController := NewMotorController(robot.leftMotor, robot.rightMotor)
	navigationModule := NewNavigationModule(motorController, sensorModule)
	behaviorPatterns := NewBehaviorPatterns(robot.statusLed, robot.buzzer)
	calibrationModule := NewCalibrationModule(robot, sensorModule, motorController)

	robot.Initialize()
	calibrationModule.CalibrateComplete()
	navigationModule.SetBehaviorMode(RANDOM_WALK_MODE)

	var lastState int = -1
	for {
		navigationModule.Update()

		currentState := navigationModule.GetCurrentState()
		if currentState != lastState {
			behaviorPatterns.IndicateStateChange(currentState)
			lastState = currentState
		}

		time.Sleep(time.Millisecond * 100)
	}
}
