package main

const (
	MOVE_FORWARD = iota
	MOVE_BACKWARD
	TURN_LEFT
	TURN_RIGHT
	STOP
)

type MotorController struct {
	leftMotor        *Motor
	rightMotor       *Motor
	currentDirection int
}

func NewMotorController(leftMotor, rightMotor *Motor) *MotorController {
	return &MotorController{
		leftMotor:        leftMotor,
		rightMotor:       rightMotor,
		currentDirection: STOP,
	}
}

func (mc *MotorController) SetDirection(direction int) {
	mc.currentDirection = direction

	switch direction {
	case MOVE_FORWARD:
		mc.leftMotor.Forward()
		mc.rightMotor.Forward()
	case MOVE_BACKWARD:
		mc.leftMotor.Backward()
		mc.rightMotor.Backward()
	case TURN_LEFT:
		mc.leftMotor.Backward()
		mc.rightMotor.Forward()
	case TURN_RIGHT:
		mc.leftMotor.Forward()
		mc.rightMotor.Backward()
	case STOP:
		mc.leftMotor.Stop()
		mc.rightMotor.Stop()
	}
}

func busyWait(loops int) {
	for i := 0; i < loops; i++ {
	}
}

func (mc *MotorController) MoveForLoops(direction int, loops int) {
	mc.SetDirection(direction)
	busyWait(loops)
	mc.SetDirection(STOP)
}

func (mc *MotorController) TurnForLoops(direction int, loops int) {
	if direction != TURN_LEFT && direction != TURN_RIGHT {
		return
	}
	mc.SetDirection(direction)
	busyWait(loops)
	mc.SetDirection(STOP)
}

func (mc *MotorController) MoveRandomly(seed uint8) {
	direction := int(seed % 4)
	if direction > TURN_RIGHT {
		direction = MOVE_FORWARD
	}
	mc.MoveForLoops(direction, 5000)
}
