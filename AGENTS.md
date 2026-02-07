# AGENTS.md

This file provides guidance to AI agents when working with code in this repository.

## Project Overview

TinyGo-based desk pet robot firmware for Arduino Uno/Nano. Uses ultrasonic (HC-SR04) for obstacle avoidance and IR sensors (A1-A4) for edge detection. Requires TinyGo toolchain.

## Build Commands

```bash
make build          # Build for Arduino Uno → firmware.hex
make build-nano     # Build for Arduino Nano
make flash          # Flash to Uno (auto-detects port on macOS/Linux)
make flash-nano     # Flash to Nano
make test           # Run unit tests (internal/navlogic only)
make run            # Run in simavr emulator (no board needed)
make fmt            # Format code (go fmt + gofmt -s -w)
```

To set port explicitly: `make flash PORT=/dev/cu.usbmodem14101` (macOS) or `PORT=COM3` (Windows).

## Architecture

**Module dependency flow:**
```
main.go → Robot → SensorModule, MotorController → NavigationModule → BehaviorPatterns
                                                → CalibrationModule
```

**Key design patterns:**

1. **Hardware abstraction**: `hardware.go` defines pin constants and low-level `Motor`/`Robot` types. All `machine` package imports are isolated to hardware-touching files.

2. **State machine**: Navigation uses a state machine (IDLE → MOVING → OBSTACLE_AVOIDANCE/EDGE_AVOIDANCE). Edge detection takes priority over obstacle detection.

3. **Testable logic separation**: `internal/navlogic/` contains pure state transition logic with no hardware dependencies. This is the only unit-testable code—test with standard `go test`, not TinyGo.

4. **Debug build tag**: `debug_debug.go` and `debug_release.go` provide `debugPrint()`. Release builds (default) use a no-op to save flash. Build with `-tags=debug` for serial output during development.

## Key Constants

- **Thresholds** (`sensors.go`): `OBSTACLE_DISTANCE_THRESHOLD` (cm), `EDGE_DETECTION_THRESHOLD` (ADC value)
- **Pins** (`hardware.go`): Motor pins D5/D6, ultrasonic D7/A0, IR sensors A1-A4, LED D13, buzzer D8

## Firmware Size Constraints

Arduino Uno/Nano have 32KB flash. The Makefile uses `-scheduler=none -gc=leaking` to minimize size. If builds overflow, the `debugPrint` calls are the main target for reduction (already gated by build tag in release).
