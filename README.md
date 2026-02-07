# Tiny Pet

Small desk pet robot (TinyGo): random movement, obstacle avoidance (ultrasonic), edge avoidance (IR). Runs on Arduino Uno/Nano.

## Requirements

- **Board:** Arduino Uno (`arduino`) or Nano (`arduino-nano`). 5 V logic.
- **Software:** Go 1.20+, [TinyGo](https://tinygo.org/getting-started/install/), avrdude (for flash).

## Features

- **Random movement** — Drives forward and occasionally turns at random to wander on a flat surface.
- **Obstacle avoidance** — Ultrasonic sensor (HC-SR04) detects obstacles ahead; robot stops, reverses, then turns away. Threshold: `OBSTACLE_DISTANCE_THRESHOLD` in `sensors.go`.
- **Edge detection** — Four IR sensors (A1–A4) detect desk edges; robot stops, reverses, and turns to avoid falling. Threshold: `EDGE_DETECTION_THRESHOLD` in `sensors.go`.
- **Interaction (optional)** — Status LED (D13) and buzzer (D8) indicate current state (moving, avoiding obstacle, avoiding edge). Calibration on startup is indicated by LED blinks and beeps.

## Parts list

### Board
| Item | Note |
|------|------|
| Arduino Uno or Arduino Nano | Uno: target `arduino`. Nano: target `arduino-nano`. 5 V. |

### Required
| Component | Qty | Spec |
|-----------|-----|------|
| DC gear motors + motor driver | 2 motors, 1 driver | Driver with logic-level inputs (e.g. L298N, TB6612). Arduino pins → driver IN1/IN2; motor power from separate supply. |
| Ultrasonic distance sensor | 1 | HC-SR04 or compatible. Trig + Echo (digital). |
| IR sensors (analog) | 4 | Analog output to A1–A4. Lower ADC = edge (e.g. TCRT5000-style). |
| Power supply | 1 | 5 V for Uno/Nano (USB or regulated). For battery: step-up to 5 V or USB power bank. |
| Wheels | 2 | To fit motor shafts (e.g. 40–65 mm). |
| Caster wheel | 1 | Front or rear, for balance. |

### Optional
| Component | Pin in code |
|-----------|-------------|
| Status LED | D13 (often built-in) |
| Buzzer | D8 (other leg GND) |
| MPU6050 (I2C) | SDA, SCL |
| Button | Free digital pin (not in current code) |

## Wiring (Arduino pins)

```
D5, D6  → Motor driver (left, right). Do not power motors from Arduino 5 V.
D7, A0  → Ultrasonic Trig, Echo (HC-SR04)
A1–A4   → IR edge sensors (analog). Lower ADC = edge.
D13, D8 → Optional: LED, Buzzer
```

Pin constants: `hardware.go`. Thresholds: `sensors.go` (`OBSTACLE_DISTANCE_THRESHOLD`, `EDGE_DETECTION_THRESHOLD`).

## Build & flash

```bash
go mod tidy
tinygo build -o firmware.hex -target arduino .          # Uno
tinygo flash -target arduino -port /dev/cu.usbmodem* . # set port: ls /dev/cu.* (macOS), COM* (Windows)
```

Nano: use `-target arduino-nano`.

## Run

Wire → power 5 V → flash. On startup: short calibration (LED/beep). Then it wanders and avoids obstacles/edges.

## Development

### Project layout

| Path | Description |
|------|-------------|
| `main.go` | Entry point, main loop, module wiring |
| `hardware.go` | Pin constants, `Motor`, `Robot`, board init |
| `motors.go` | `MotorController` — direction, speed, timed moves |
| `sensors.go` | `SensorModule` — ultrasonic, IR, thresholds |
| `navigation.go` | `NavigationModule` — state machine, behavior mode |
| `behaviors.go` | `BehaviorPatterns` — LED and buzzer feedback |
| `calibration.go` | `CalibrationModule` — sensor/motor calibration |
| `internal/navlogic/` | Pure state logic (no hardware); unit-testable |

### Emulator (no board)

Run firmware under simavr to check that the program starts and the main loop runs. Sensor and motor I/O are simulated (default values), so behavior will not match real hardware.

```bash
brew install simavr   # macOS; or install simavr for your OS
tinygo run -target=arduino-nano .
```

### Unit tests

Navigation state logic only (sensor inputs → next state). Uses the standard Go toolchain; no TinyGo or board needed.

```bash
go test ./internal/navlogic/... -v
```

### Tuning

- Obstacle/edge thresholds: `sensors.go` (`OBSTACLE_DISTANCE_THRESHOLD`, `EDGE_DETECTION_THRESHOLD`).
- Avoidance timings: `navigation.go`. Runtime adjustment via `CalibrationModule.AdjustThresholds()`.

## License

MIT. See [LICENSE](LICENSE).
