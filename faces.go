package main

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

var white = color.RGBA{R: 255, G: 255, B: 255, A: 255}

// setHLine draws a horizontal line.
func setHLine(dev *ssd1306.Device, x, y, w int16) {
	for i := int16(0); i < w; i++ {
		dev.SetPixel(x+i, y, white)
	}
}

// setFillRect draws a filled rectangle.
func setFillRect(dev *ssd1306.Device, x, y, w, h int16) {
	for dy := int16(0); dy < h; dy++ {
		setHLine(dev, x, y+dy, w)
	}
}

// setFillCircle draws a filled circle using brute-force dx²+dy²≤r².
func setFillCircle(dev *ssd1306.Device, cx, cy, r int16) {
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				dev.SetPixel(cx+dx, cy+dy, white)
			}
		}
	}
}

// setCircle draws a circle outline using the midpoint algorithm.
func setCircle(dev *ssd1306.Device, cx, cy, r int16) {
	x := r
	y := int16(0)
	p := 1 - r

	for x >= y {
		dev.SetPixel(cx+x, cy+y, white)
		dev.SetPixel(cx-x, cy+y, white)
		dev.SetPixel(cx+x, cy-y, white)
		dev.SetPixel(cx-x, cy-y, white)
		dev.SetPixel(cx+y, cy+x, white)
		dev.SetPixel(cx-y, cy+x, white)
		dev.SetPixel(cx+y, cy-x, white)
		dev.SetPixel(cx-y, cy-x, white)
		y++
		if p <= 0 {
			p += 2*y + 1
		} else {
			x--
			p += 2*(y-x) + 1
		}
	}
}

// Face layout constants (128x64 screen).
const (
	eyeLeftX  = 40
	eyeRightX = 88
	eyeY      = 22
	mouthCX   = 64
	mouthY    = 48
)

// drawNeutralFace draws half-closed horizontal line eyes and a small straight mouth.
func drawNeutralFace(dev *ssd1306.Device) {
	// Half-closed eyes: short horizontal bars
	setFillRect(dev, eyeLeftX-8, eyeY-1, 16, 3)
	setFillRect(dev, eyeRightX-8, eyeY-1, 16, 3)
	// Small straight mouth
	setFillRect(dev, mouthCX-10, mouthY, 20, 2)
}

// drawHappyFace draws round open eyes and a curved smile.
func drawHappyFace(dev *ssd1306.Device) {
	// Round open eyes
	setFillCircle(dev, eyeLeftX, eyeY, 7)
	setFillCircle(dev, eyeRightX, eyeY, 7)
	// Curved smile (parabola opening downward in screen coords)
	for x := int16(mouthCX - 14); x <= mouthCX+14; x++ {
		dx := x - mouthCX
		dy := dx * dx / 28 // gentle curve
		dev.SetPixel(x, mouthY+dy, white)
		dev.SetPixel(x, mouthY+dy+1, white) // 2px thick
	}
}

// drawSurprisedFace draws large circle-outline eyes and an O-shaped mouth.
func drawSurprisedFace(dev *ssd1306.Device) {
	// Large open circle eyes
	setCircle(dev, eyeLeftX, eyeY, 9)
	setCircle(dev, eyeLeftX, eyeY, 8)
	setCircle(dev, eyeRightX, eyeY, 9)
	setCircle(dev, eyeRightX, eyeY, 8)
	// O-shaped mouth
	setCircle(dev, mouthCX, mouthY+2, 5)
	setCircle(dev, mouthCX, mouthY+2, 4)
}

// drawScaredFace draws large circle eyes with tiny pupils and a frown.
func drawScaredFace(dev *ssd1306.Device) {
	// Large circle outline eyes
	setCircle(dev, eyeLeftX, eyeY, 9)
	setCircle(dev, eyeLeftX, eyeY, 8)
	setCircle(dev, eyeRightX, eyeY, 9)
	setCircle(dev, eyeRightX, eyeY, 8)
	// Tiny pupils
	setFillCircle(dev, eyeLeftX, eyeY, 2)
	setFillCircle(dev, eyeRightX, eyeY, 2)
	// Frown (inverted parabola)
	for x := int16(mouthCX - 14); x <= mouthCX+14; x++ {
		dx := x - mouthCX
		dy := -(dx * dx / 28)
		dev.SetPixel(x, mouthY+4+dy, white)
		dev.SetPixel(x, mouthY+4+dy+1, white)
	}
}

// drawExcitedFace draws round eyes with sparkle crosses and a wide smile.
func drawExcitedFace(dev *ssd1306.Device) {
	// Round eyes
	setFillCircle(dev, eyeLeftX, eyeY, 7)
	setFillCircle(dev, eyeRightX, eyeY, 7)
	// Sparkle crosses on each eye
	for _, cx := range [2]int16{eyeLeftX, eyeRightX} {
		setFillRect(dev, cx-1, eyeY-12, 3, 5) // vertical top
		setFillRect(dev, cx-1, eyeY+8, 3, 5)  // vertical bottom
		setFillRect(dev, cx-12, eyeY-1, 5, 3)  // horizontal left
		setFillRect(dev, cx+8, eyeY-1, 5, 3)   // horizontal right
	}
	// Wide smile
	for x := int16(mouthCX - 18); x <= mouthCX+18; x++ {
		dx := x - mouthCX
		dy := dx * dx / 40
		dev.SetPixel(x, mouthY+dy, white)
		dev.SetPixel(x, mouthY+dy+1, white)
	}
}

// drawBlinkFace draws thin closed-eye lines and a small mouth (for blink animation).
func drawBlinkFace(dev *ssd1306.Device) {
	// Closed eyes: thin horizontal lines
	setHLine(dev, eyeLeftX-8, eyeY, 16)
	setHLine(dev, eyeRightX-8, eyeY, 16)
	// Small straight mouth
	setFillRect(dev, mouthCX-10, mouthY, 20, 2)
}
