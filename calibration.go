package main

type CalibrationModule struct {
	robot           *Robot
	sensorModule    *SensorModule
	motorController *MotorController
	calibrated      bool
}

func NewCalibrationModule(robot *Robot, sensorModule *SensorModule, motorController *MotorController) *CalibrationModule {
	return &CalibrationModule{
		robot:           robot,
		sensorModule:    sensorModule,
		motorController: motorController,
		calibrated:      false,
	}
}

func (cm *CalibrationModule) IsCalibrated() bool {
	return cm.calibrated
}

func (cm *CalibrationModule) CalibrateSensors() {
	cm.robot.BlinkLED(2)
	debugPrint("Starting sensor calibration...")
	debugPrint("Measuring baseline distance...")
	distance := cm.sensorModule.ReadUltrasonicDistance()
	debugPrint("Ultrasonic distance:", distance)
	busyWait(20000)
	debugPrint("Measuring IR sensor baselines...")
	irValues := cm.sensorModule.ReadIRSensors()
	for i, value := range irValues {
		debugPrint("IR sensor", i, "edge:", value)
	}
	busyWait(20000)
	debugPrint("Sensor calibration complete!")
	cm.robot.BeepLoops(10000)
	cm.calibrated = true
}

func (cm *CalibrationModule) CalibrateMotors() {
	cm.robot.BlinkLED(3)
	debugPrint("Starting motor calibration...")
	for _, dir := range []int{MOVE_FORWARD, MOVE_BACKWARD, TURN_LEFT, TURN_RIGHT} {
		debugPrint("Testing movement...")
		cm.motorController.SetDirection(dir)
		busyWait(50000)
		cm.motorController.SetDirection(STOP)
		busyWait(20000)
	}
	debugPrint("Motor calibration complete!")
	cm.robot.BeepLoops(5000)
}

func (cm *CalibrationModule) CalibrateComplete() {
	debugPrint("Starting complete calibration...")
	cm.CalibrateSensors()
	cm.CalibrateMotors()
	debugPrint("Complete calibration finished!")
	cm.robot.BlinkLED(5)
}
