package main

import (
	"image/color"

	"tinygo.org/x/drivers/ssd1306"
)

var white = color.RGBA{R: 255, G: 255, B: 255, A: 255}

func setHLine(dev *ssd1306.Device, x, y, w int16) {
	for i := int16(0); i < w; i++ {
		dev.SetPixel(x+i, y, white)
	}
}

func setFillRect(dev *ssd1306.Device, x, y, w, h int16) {
	for dy := int16(0); dy < h; dy++ {
		setHLine(dev, x, y+dy, w)
	}
}

func setFillCircle(dev *ssd1306.Device, cx, cy, r int16) {
	for dy := -r; dy <= r; dy++ {
		for dx := -r; dx <= r; dx++ {
			if dx*dx+dy*dy <= r*r {
				dev.SetPixel(cx+dx, cy+dy, white)
			}
		}
	}
}

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

const (
	eyeLeftX  = 40
	eyeRightX = 88
	eyeY      = 11
	mouthCX   = 64
	mouthY    = 24
)

func drawNeutralFace(dev *ssd1306.Device) {
	setFillRect(dev, eyeLeftX-4, eyeY-1, 8, 2)
	setFillRect(dev, eyeRightX-4, eyeY-1, 8, 2)
	setFillRect(dev, mouthCX-5, mouthY, 10, 1)
}

func drawHappyFace(dev *ssd1306.Device) {
	setFillCircle(dev, eyeLeftX, eyeY, 4)
	setFillCircle(dev, eyeRightX, eyeY, 4)
	for x := int16(mouthCX - 7); x <= mouthCX+7; x++ {
		dx := x - mouthCX
		dy := dx * dx / 14
		dev.SetPixel(x, mouthY+dy, white)
	}
}

func drawSurprisedFace(dev *ssd1306.Device) {
	setCircle(dev, eyeLeftX, eyeY, 5)
	setCircle(dev, eyeLeftX, eyeY, 4)
	setCircle(dev, eyeRightX, eyeY, 5)
	setCircle(dev, eyeRightX, eyeY, 4)
	setCircle(dev, mouthCX, mouthY+1, 3)
	setCircle(dev, mouthCX, mouthY+1, 2)
}

func drawScaredFace(dev *ssd1306.Device) {
	setCircle(dev, eyeLeftX, eyeY, 5)
	setCircle(dev, eyeLeftX, eyeY, 4)
	setCircle(dev, eyeRightX, eyeY, 5)
	setCircle(dev, eyeRightX, eyeY, 4)
	setFillCircle(dev, eyeLeftX, eyeY, 1)
	setFillCircle(dev, eyeRightX, eyeY, 1)
	for x := int16(mouthCX - 7); x <= mouthCX+7; x++ {
		dx := x - mouthCX
		dy := -(dx * dx / 14)
		dev.SetPixel(x, mouthY+2+dy, white)
	}
}

func drawExcitedFace(dev *ssd1306.Device) {
	setFillCircle(dev, eyeLeftX, eyeY, 4)
	setFillCircle(dev, eyeRightX, eyeY, 4)
	for _, cx := range [2]int16{eyeLeftX, eyeRightX} {
		setFillRect(dev, cx-1, eyeY-6, 2, 3)
		setFillRect(dev, cx-1, eyeY+4, 2, 3)
		setFillRect(dev, cx-6, eyeY-1, 3, 2)
		setFillRect(dev, cx+4, eyeY-1, 3, 2)
	}
	for x := int16(mouthCX - 9); x <= mouthCX+9; x++ {
		dx := x - mouthCX
		dy := dx * dx / 20
		dev.SetPixel(x, mouthY+dy, white)
	}
}

func drawBlinkFace(dev *ssd1306.Device) {
	setHLine(dev, eyeLeftX-4, eyeY, 8)
	setHLine(dev, eyeRightX-4, eyeY, 8)
	setFillRect(dev, mouthCX-5, mouthY, 10, 1)
}
