package main

// Movement directions.
const (
	MOVE_FORWARD = iota
	MOVE_BACKWARD
	TURN_LEFT
	TURN_RIGHT
	STOP
)

// MotorController manages differential drive movement.
type MotorController struct {
	leftMotor        *Motor
	rightMotor       *Motor
	currentDirection int
}

// NewMotorController creates a MotorController with the given motors.
func NewMotorController(leftMotor, rightMotor *Motor) *MotorController {
	return &MotorController{
		leftMotor:        leftMotor,
		rightMotor:       rightMotor,
		currentDirection: STOP,
	}
}

// SetDirection sets the robot's movement direction.
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

// MoveForLoops moves in a direction for a specified number of loop iterations.
func (mc *MotorController) MoveForLoops(direction int, loops int) {
	mc.SetDirection(direction)
	busyWait(loops)
	mc.SetDirection(STOP)
}

// TurnForLoops turns the robot for a specified number of loop iterations.
func (mc *MotorController) TurnForLoops(direction int, loops int) {
	if direction != TURN_LEFT && direction != TURN_RIGHT {
		return
	}
	mc.SetDirection(direction)
	busyWait(loops)
	mc.SetDirection(STOP)
}

// MoveRandomly moves the robot in a pseudo-random direction based on seed.
func (mc *MotorController) MoveRandomly(seed uint8) {
	direction := int(seed % 4)
	if direction > TURN_RIGHT {
		direction = MOVE_FORWARD
	}
	mc.MoveForLoops(direction, 5000)
}
