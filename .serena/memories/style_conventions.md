# Code Style & Conventions

- **Language:** Go with TinyGo constraints
- **Naming:** UPPER_SNAKE for constants, CamelCase for types/exported, camelCase for unexported
- **Modules:** Each hardware subsystem in its own file with a struct (e.g., MotorController, SensorModule)
- **Hardware isolation:** All `machine` package imports confined to hardware-touching files
- **Testability:** Pure logic in internal/navlogic/ (standard Go tests, no TinyGo)
- **Debug:** `debugPrint()` gated by build tag (debug_debug.go vs debug_release.go)
- **Motor control:** Simple digital HIGH/LOW (no PWM currently), differential drive
- **Timing:** busyWait loops instead of time.Sleep where needed (no scheduler)
