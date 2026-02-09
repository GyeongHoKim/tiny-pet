# Suggested Commands

## Build
- `make build` — Arduino Uno → firmware.hex
- `make build-nano` — Arduino Nano
- `make build-bluepill` — STM32 Blue Pill → firmware_bluepill.elf

## Flash
- `make flash` — Flash Uno (port auto-detected on macOS)
- `make flash-nano` — Flash Nano
- `make flash-bluepill` — Flash Blue Pill (ST-Link v2 + OpenOCD)
- `make flash PORT=/dev/cu.usbmodem14101` — Explicit port (Arduino)

## Development
- `make fmt` — Format Go (go fmt + gofmt -s -w)
- `make tidy` — go mod tidy
- `make test` — Unit tests (internal/navlogic only, standard Go)
- `make run` — simavr emulator (no board)
- `make clean` — Remove firmware.hex, firmware_bluepill.elf

## Task completion
1. `make fmt`
2. `make build` or `make build-bluepill` as needed
3. `make test`
