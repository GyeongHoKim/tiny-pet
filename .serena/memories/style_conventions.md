# Code Style & Conventions

- **Language:** Go with TinyGo constraints
- **Naming:** UPPER_SNAKE for constants, CamelCase for types/exported, camelCase for unexported
- **Comments:** Minimal. Only essential GODoc (e.g. conversion semantics, non-obvious types). No redundant inline comments. Project description and usage live in README.md.
- **Modules:** Each hardware subsystem in its own file with a struct (MotorController, SensorModule, etc.). Board-specific code: build tags (`bluepill` / `!bluepill`) in hardware_*.go and sensors_*.go.
- **Hardware isolation:** All `machine` imports in hardware-touching files only
- **Testability:** Pure logic in internal/navlogic/ — standard Go tests, no TinyGo
- **Debug:** debugPrint() gated by build tag (debug_debug.go vs debug_release.go)
- **Motor control:** Digital HIGH/LOW via L298N (no PWM). Differential drive.
