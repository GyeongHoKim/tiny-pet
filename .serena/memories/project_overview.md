# Tiny Pet - Project Overview

## Purpose
TinyGo-based desk pet robot firmware for Arduino Uno/Nano and STM32 Blue Pill. Wandering robot with obstacle avoidance (ultrasonic), edge detection (IR), and OLED face expressions.

## Tech Stack
- **Language:** Go (TinyGo)
- **Targets:** Arduino Uno / Arduino Nano (32KB flash, 2KB SRAM, 5V); STM32 Blue Pill (STM32F103C8, 64KB+ flash, 20KB SRAM, 3.3V). Build tag `bluepill` selects Blue Pill; default is Arduino.
- **Dependencies:** tinygo.org/x/drivers v0.27.0 (SSD1306 OLED)
- **Build:** Make + TinyGo. Blue Pill flash: ST-Link v2 + OpenOCD.

## Code Structure
| File | Role |
|------|------|
| main.go | Entry point, main loop |
| hardware_arduino.go / hardware_bluepill.go | Pin constants, Motor, Robot (build tag selects; all `machine` imports here) |
| motors.go | MotorController — differential drive |
| sensors.go / sensors_bluepill.go | SensorModule — ultrasonic, IR (Blue Pill uses time-based ultrasonic) |
| navigation.go | NavigationModule — state machine |
| behaviors.go | BehaviorPatterns — LED/buzzer |
| display.go | DisplayModule — SSD1306 OLED |
| faces.go | Procedural face drawing |
| calibration.go | CalibrationModule |
| internal/navlogic/ | Pure state logic, unit-testable with standard Go |

## Hardware (see README for full wiring)
- Arduino: motors D4–D9, ultrasonic D7/A0, IR A1–A2, I2C A4/A5, LED D13, buzzer D8.
- Blue Pill: motors PA8–PA11, ultrasonic PA12/PB10, IR PA1/PA2, I2C PB7/PB6, LED PC13, buzzer PB15.

## Key Constraints
- Uno/Nano: 32KB flash, 2KB SRAM (Makefile: -scheduler=none -gc=leaking). Blue Pill has more headroom.
- 128x32 OLED = 512-byte buffer. Debug output gated by build tag `debug`.
