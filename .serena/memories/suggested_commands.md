# Suggested Commands

## Build
- `make build` — Build for Arduino Uno → firmware.hex
- `make build-nano` — Build for Arduino Nano

## Flash
- `make flash` — Flash to Uno (port auto-detected on macOS)
- `make flash-nano` — Flash to Nano
- `make flash PORT=/dev/cu.usbmodem14101` — Explicit port

## Development
- `make fmt` — Format Go code (go fmt + gofmt -s -w)
- `make tidy` — go mod tidy
- `make test` — Run unit tests (internal/navlogic only, standard Go)
- `make run` — Run in simavr emulator (no board needed)
- `make clean` — Remove firmware.hex

## Task Completion Checklist
1. `make fmt` — format code
2. `make build` — verify it compiles for target
3. `make test` — run unit tests
