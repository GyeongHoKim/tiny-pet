package main

import (
	"machine"
	"time"
)

// Movement directions
const (
	MOVE_FORWARD = iota
	MOVE_BACKWARD
	TURN_LEFT
	TURN_RIGHT
	STOP
)

// Movement speeds
const (
	SPEED_SLOW = iota
	SPEED_MEDIUM
	SPEED_FAST
)

// MotorController manages the robot's movement
type MotorController struct {
	leftMotor  *Motor
	rightMotor *Motor
	currentDirection int
	currentSpeed     int
}

// NewMotorController creates a new motor controller instance
func NewMotorController(leftMotor, rightMotor *Motor) *MotorController {
	return &MotorController{
		leftMotor:  leftMotor,
		rightMotor: rightMotor,
		currentDirection: STOP,
		currentSpeed:     SPEED_MEDIUM,
	}
}

// SetDirection sets the movement direction
func (mc *MotorController) SetDirection(direction int) {
	mc.currentDirection = direction
	
	switch direction {
	case MOVE_FORWARD:
		mc.moveForward()
	case MOVE_BACKWARD:
		mc.moveBackward()
	case TURN_LEFT:
		mc.turnLeft()
	case TURN_RIGHT:
		mc.turnRight()
	case STOP:
		mc.stop()
	}
}

// SetSpeed sets the movement speed
func (mc *MotorController) SetSpeed(speed int) {
	mc.currentSpeed = speed
	// In a real implementation with PWM, we would adjust the speed here
	// For now, we'll just store the speed value
}

// moveForward moves the robot forward
func (mc *MotorController) moveForward() {
	mc.leftMotor.SetSpeed(true)  // Forward
	mc.rightMotor.SetSpeed(true) // Forward
}

// moveBackward moves the robot backward
func (mc *MotorController) moveBackward() {
	mc.leftMotor.SetSpeed(false) // Backward (reverse polarity)
	mc.rightMotor.SetSpeed(false) // Backward (reverse polarity)
}

// turnLeft turns the robot left
func (mc *MotorController) turnLeft() {
	mc.leftMotor.SetSpeed(false) // Left backward
	mc.rightMotor.SetSpeed(true) // Right forward
}

// turnRight turns the robot right
func (mc *MotorController) turnRight() {
	mc.leftMotor.SetSpeed(true)  // Left forward
	mc.rightMotor.SetSpeed(false) // Right backward
}

// stop stops the robot
func (mc *MotorController) stop() {
	mc.leftMotor.Stop()
	mc.rightMotor.Stop()
}

// MoveForTime moves the robot in a direction for a specified time
func (mc *MotorController) MoveForTime(direction int, duration time.Duration) {
	mc.SetDirection(direction)
	time.Sleep(duration)
	mc.SetDirection(STOP)
}

// TurnForTime turns the robot for a specified time
func (mc *MotorController) TurnForTime(direction int, duration time.Duration) {
	if direction != TURN_LEFT && direction != TURN_RIGHT {
		return // Only allow turning directions
	}
	mc.SetDirection(direction)
	time.Sleep(duration)
	mc.SetDirection(STOP)
}

// GetCurrentDirection returns the current movement direction
func (mc *MotorController) GetCurrentDirection() int {
	return mc.currentDirection
}

// GetCurrentSpeed returns the current speed
func (mc *MotorController) GetCurrentSpeed() int {
	return mc.currentSpeed
}

// MoveRandomly moves the robot in a random pattern
func (mc *MotorController) MoveRandomly() {
	// Choose a random direction
	// In a real implementation, we'd use a proper random function
	// For TinyGo, we'll simulate randomness using time
	seed := time.Now().UnixNano() % 4
	direction := int(seed)
	
	// Limit to valid directions
	if direction > TURN_RIGHT {
		direction = MOVE_FORWARD
	}
	
	// Move in the chosen direction for a random duration (between 500ms and 2s)
	duration := time.Duration(500+(time.Now().UnixNano()%1500)) * time.Millisecond
	
	mc.MoveForTime(direction, duration)
}