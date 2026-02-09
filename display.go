package main

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"
)

// Expression constants.
const (
	EXPR_NEUTRAL = iota
	EXPR_HAPPY
	EXPR_SURPRISED
	EXPR_SCARED
	EXPR_EXCITED
	EXPR_BLINK
)

// Blink timing: ~4 seconds at 100ms loop interval = 40 ticks.
const (
	blinkInterval uint8 = 40
	blinkDuration uint8 = 2 // blink lasts 2 ticks (~200ms)
)

// DisplayModule wraps the SSD1306 OLED driver and manages face expressions.
type DisplayModule struct {
	device       ssd1306.Device
	currentExpr  int
	animCounter  uint8
	blinkCounter uint8
	isBlinking   bool
}

// NewDisplayModule initializes the SSD1306 OLED on the given I2C bus.
func NewDisplayModule(bus *machine.I2C) *DisplayModule {
	dm := &DisplayModule{
		device:      ssd1306.NewI2C(bus),
		currentExpr: EXPR_NEUTRAL,
	}
	// 128x32: 512-byte buffer to fit 2KB SRAM (Uno/Nano).
	dm.device.Configure(ssd1306.Config{
		Width:   128,
		Height:  32,
		Address: 0x3C,
	})
	dm.device.ClearDisplay()
	return dm
}

// ShowExpression clears the buffer, draws the given expression, and sends to display.
func (dm *DisplayModule) ShowExpression(expr int) {
	dm.currentExpr = expr
	dm.device.ClearBuffer()
	switch expr {
	case EXPR_NEUTRAL:
		drawNeutralFace(&dm.device)
	case EXPR_HAPPY:
		drawHappyFace(&dm.device)
	case EXPR_SURPRISED:
		drawSurprisedFace(&dm.device)
	case EXPR_SCARED:
		drawScaredFace(&dm.device)
	case EXPR_EXCITED:
		drawExcitedFace(&dm.device)
	case EXPR_BLINK:
		drawBlinkFace(&dm.device)
	}
	dm.device.Display()
}

// ShowStateExpression maps a navigation state to an expression and displays it.
func (dm *DisplayModule) ShowStateExpression(state int) {
	var expr int
	switch state {
	case IDLE_STATE:
		expr = EXPR_NEUTRAL
	case MOVING_STATE:
		expr = EXPR_HAPPY
	case OBSTACLE_AVOIDANCE_STATE:
		expr = EXPR_SURPRISED
	case EDGE_AVOIDANCE_STATE:
		expr = EXPR_SCARED
	case INTERACTING_STATE:
		expr = EXPR_EXCITED
	default:
		expr = EXPR_NEUTRAL
	}
	dm.ShowExpression(expr)
}

// UpdateAnimation handles periodic blink animation.
// Call this every main loop iteration (~100ms).
func (dm *DisplayModule) UpdateAnimation() {
	dm.animCounter++

	if dm.isBlinking {
		dm.blinkCounter++
		if dm.blinkCounter >= blinkDuration {
			// End blink: restore previous expression
			dm.isBlinking = false
			dm.blinkCounter = 0
			dm.ShowExpression(dm.currentExpr)
		}
		return
	}

	if dm.animCounter >= blinkInterval {
		dm.animCounter = 0
		// Start blink
		savedExpr := dm.currentExpr
		dm.isBlinking = true
		dm.blinkCounter = 0
		dm.device.ClearBuffer()
		drawBlinkFace(&dm.device)
		dm.device.Display()
		dm.currentExpr = savedExpr // preserve so we restore after blink
	}
}
