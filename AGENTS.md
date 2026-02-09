# AGENTS.md

This file provides guidance to AI agents when working with code in this repository.

## Project Overview

TinyGo-based desk pet robot firmware for Arduino Uno/Nano and STM32 Blue Pill. Uses ultrasonic (HC-SR04) for obstacle avoidance, IR sensors for edge detection, and SSD1306 OLED (I2C) for face expressions. Requires TinyGo toolchain.

## Build Commands

```bash
make build          # Build for Blue Pill (default) → firmware_bluepill.elf
make build-uno      # Build for Arduino Uno → firmware.hex
make build-nano     # Build for Arduino Nano
make flash          # Flash to Uno (auto-detects port on macOS/Linux)
make flash-nano     # Flash to Nano
make flash-bluepill # Flash to Blue Pill (ST-Link v2 + OpenOCD required)
make test           # Run unit tests (internal/navlogic only)
make run            # Run in simavr emulator (no board needed)
make fmt            # Format code (go fmt + gofmt -s -w)
```

To set port explicitly: `make flash PORT=/dev/cu.usbmodem14101` (macOS) or `PORT=COM3` (Windows).

## Architecture

**Module dependency flow:**
```
main.go → Robot → SensorModule, MotorController → NavigationModule → BehaviorPatterns
                                                                   → DisplayModule (faces.go)
                                                → CalibrationModule
```

**Key design patterns:**

1. **Hardware abstraction**: Board-specific files define pin constants and low-level `Motor`/`Robot` types: `hardware_arduino.go` (Uno/Nano), `hardware_bluepill.go` (STM32 Blue Pill). All `machine` package imports are isolated to hardware-touching files. Build tag `bluepill` selects Blue Pill; default is Arduino.

2. **State machine**: Navigation uses a state machine (IDLE → MOVING → OBSTACLE_AVOIDANCE/EDGE_AVOIDANCE). Edge detection takes priority over obstacle detection.

3. **Testable logic separation**: `internal/navlogic/` contains pure state transition logic with no hardware dependencies. This is the only unit-testable code—test with standard `go test`, not TinyGo.

4. **Debug build tag**: `debug_debug.go` and `debug_release.go` provide `debugPrint()`. Release builds (default) use a no-op to save flash. Build with `-tags=debug` for serial output during development.

## Key Constants

- **Thresholds** (`sensors.go` / `sensors_bluepill.go`): `OBSTACLE_DISTANCE_THRESHOLD` (cm), `EDGE_DETECTION_THRESHOLD` (ADC value)
- **Pins (Arduino)** `hardware_arduino.go`: Motor D5/D4 (left), D6/D9 (right), ultrasonic D7/A0, IR A1-A2, I2C A4/A5, LED D13, buzzer D8
- **Pins (Blue Pill)** `hardware_bluepill.go`: Motor PA8/PA9 (left), PA10/PA11 (right), ultrasonic PA12/PB10, IR PA1/PA2 (ADC), I2C PB7/PB6 (I2C0), LED PC13, buzzer PB15

## Firmware Size Constraints

Arduino Uno/Nano have 32KB flash. The Makefile uses `-scheduler=none -gc=leaking` to minimize size. If builds overflow, the `debugPrint` calls are the main target for reduction (already gated by build tag in release). Blue Pill has 64KB+ flash and 20KB SRAM, so size pressure is lower.

## Blue Pill Flashing

Flash Blue Pill with ST-Link v2 and OpenOCD: connect SWIO, SWCLK, 3V3, GND to the programmer, then `make flash-bluepill`. Install OpenOCD if needed (e.g. `brew install openocd` on macOS). Ultrasonic timing on Blue Pill uses `time.Sleep(10*time.Microsecond)` for trigger and a loop-to-µs constant (`bluepillLoopsPerMicrosecond` in `sensors_bluepill.go`); tune that constant if distance readings are off.
