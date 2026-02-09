package main

import (
	"machine"

	"tinygo.org/x/drivers/ssd1306"
)

const (
	EXPR_NEUTRAL = iota
	EXPR_HAPPY
	EXPR_SURPRISED
	EXPR_SCARED
	EXPR_EXCITED
	EXPR_BLINK
)

const (
	blinkInterval uint8 = 40
	blinkDuration uint8 = 2
)

// DisplayModule drives the SSD1306 OLED and face expressions.
type DisplayModule struct {
	device       ssd1306.Device
	currentExpr  int
	animCounter  uint8
	blinkCounter uint8
	isBlinking   bool
}

func NewDisplayModule(bus *machine.I2C) *DisplayModule {
	dm := &DisplayModule{
		device:      ssd1306.NewI2C(bus),
		currentExpr: EXPR_NEUTRAL,
	}
	dm.device.Configure(ssd1306.Config{
		Width:   128,
		Height:  32,
		Address: 0x3C,
	})
	dm.device.ClearDisplay()
	return dm
}

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

func (dm *DisplayModule) UpdateAnimation() {
	dm.animCounter++

	if dm.isBlinking {
		dm.blinkCounter++
		if dm.blinkCounter >= blinkDuration {
			dm.isBlinking = false
			dm.blinkCounter = 0
			dm.ShowExpression(dm.currentExpr)
		}
		return
	}

	if dm.animCounter >= blinkInterval {
		dm.animCounter = 0
		savedExpr := dm.currentExpr
		dm.isBlinking = true
		dm.blinkCounter = 0
		dm.device.ClearBuffer()
		drawBlinkFace(&dm.device)
		dm.device.Display()
		dm.currentExpr = savedExpr
	}
}
