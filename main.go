package main

import (
	"machine"
	"time"
)

func main() {
	robot := NewRobot()

	sensorModule := NewSensorModule(robot.ultraTrig, robot.ultraEcho, &robot.irSensors)
	motorController := NewMotorController(robot.leftMotor, robot.rightMotor)
	navigationModule := NewNavigationModule(motorController, sensorModule)
	behaviorPatterns := NewBehaviorPatterns(robot.statusLed, robot.buzzer)
	calibrationModule := NewCalibrationModule(robot, sensorModule, motorController)

	machine.I2C0.Configure(machine.I2CConfig{Frequency: 400000})
	displayModule := NewDisplayModule(machine.I2C0)

	robot.Initialize()
	calibrationModule.CalibrateComplete()
	navigationModule.SetBehaviorMode(RANDOM_WALK_MODE)
	displayModule.ShowExpression(EXPR_HAPPY)

	var lastState int = -1
	for {
		navigationModule.Update()

		currentState := navigationModule.GetCurrentState()
		if currentState != lastState {
			behaviorPatterns.IndicateStateChange(currentState)
			displayModule.ShowStateExpression(currentState)
			lastState = currentState
		}
		displayModule.UpdateAnimation()

		time.Sleep(time.Millisecond * 100)
	}
}
