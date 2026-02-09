# Tiny Pet - Project Overview

## Purpose
TinyGo-based desk pet robot firmware for Arduino Uno/Nano. A small wandering robot with obstacle avoidance, edge detection, and OLED face expressions.

## Tech Stack
- **Language:** Go (TinyGo for AVR cross-compilation)
- **Target:** Arduino Uno (ATmega328P) / Arduino Nano — 32KB flash, 2KB SRAM, 5V logic
- **Dependencies:** tinygo.org/x/drivers v0.27.0 (SSD1306 OLED driver)
- **Build tool:** Make + TinyGo

## Code Structure
| File | Role |
|------|------|
| main.go | Entry point, main loop |
| hardware.go | Pin constants, Motor, Robot types (all `machine` imports) |
| motors.go | MotorController — differential drive |
| sensors.go | SensorModule — ultrasonic HC-SR04, IR edge sensors |
| navigation.go | NavigationModule — state machine |
| behaviors.go | BehaviorPatterns — LED/buzzer feedback |
| display.go | DisplayModule — SSD1306 OLED |
| faces.go | Procedural face drawing |
| calibration.go | CalibrationModule |
| internal/navlogic/ | Pure state logic, unit-testable with standard Go |

## Hardware
- 2x DC gear motors (D5, D6) with motor driver (L298N or similar)
- HC-SR04 ultrasonic sensor (D7, A0)
- 2x IR edge sensors (A1, A2) — analog
- SSD1306 128x32 OLED (I2C: A4 SDA, A5 SCL)
- Optional: LED (D13), Buzzer (D8)

## Key Constraints
- 32KB flash limit (optimization flags: -scheduler=none -gc=leaking)
- 2KB SRAM (128x32 OLED = 512-byte buffer)
- No goroutines (scheduler=none)
- Debug output gated by build tag
