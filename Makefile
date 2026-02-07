# Tiny Pet — TinyGo Arduino (Uno/Nano) desk pet robot
# Usage: make [target]. Run `make help` for targets.
# Windows: assumes PowerShell (pwsh). Unix: sh/bash.

.PHONY: build build-nano build-uno flash flash-unix flash-win flash-nano fmt tidy test run clean help

# Target board: arduino (Uno) or arduino-nano (Nano)
TARGET ?= arduino
# Serial port for flash (e.g. /dev/cu.usbmodem14101 on macOS, COM3 on Windows)
PORT ?=

FIRMWARE := firmware.hex

# TinyGo flags for smaller firmware (https://tinygo.org/docs/guides/optimizing-binaries/)
# -scheduler=none: no goroutines
# -gc=leaking: no GC (saves size)
# -panic=trap not used (causes "linker could not find symbol abort" on AVR).
# Optional: TINYGO_FLAGS="-scheduler=none" to keep GC if leaking is undesirable.
TINYGO_FLAGS ?= -scheduler=none -gc=leaking

# --- Build ---
build build-uno: TARGET = arduino
build build-uno:
	go mod tidy
	tinygo build $(TINYGO_FLAGS) -o $(FIRMWARE) -target $(TARGET) .

build-nano: TARGET = arduino-nano
build-nano:
	go mod tidy
	tinygo build $(TINYGO_FLAGS) -o $(FIRMWARE) -target $(TARGET) .

# --- Flash (auto-detect or set PORT=; Windows uses pwsh, Unix uses sh) ---
ifeq ($(OS),Windows_NT)
flash: flash-win
else
flash: flash-unix
endif

flash-unix:
	@port="$(PORT)"; \
	if [ -z "$$port" ]; then \
	  case "$$(uname -s 2>/dev/null)" in \
	    Darwin) port=$$(ls /dev/cu.usbmodem* 2>/dev/null | head -1);; \
	    Linux)  port=$$(ls /dev/ttyACM* /dev/ttyUSB* 2>/dev/null | head -1);; \
	    *)      port="";; \
	  esac; \
	fi; \
	if [ -z "$$port" ]; then \
	  echo "Error: PORT not set and could not auto-detect. Set PORT= (e.g. make flash PORT=/dev/cu.usbmodem14101, PORT=/dev/ttyACM0)"; \
	  exit 1; \
	fi; \
	echo "Using port: $$port"; \
	tinygo flash -target $(TARGET) -port "$$port" .

flash-win:
	@pwsh -NoProfile -Command "$$port = '$(PORT)'; if (-not $$port) { $$ports = [System.IO.Ports.SerialPort]::GetPortNames(); if ($$ports) { $$port = $$ports[0] } }; if (-not $$port) { Write-Error 'Error: PORT not set and could not auto-detect. Set PORT= (e.g. make flash PORT=COM3)'; exit 1 }; Write-Host ('Using port: ' + $$port); & tinygo flash -target $(TARGET) -port $$port ."

# Flash to Arduino Nano (build with build-nano first, or use: make build-nano flash-nano)
flash-nano: TARGET = arduino-nano
flash-nano: flash

# --- Format & tidy ---
fmt:
	go fmt ./...
	gofmt -s -w .

tidy:
	go mod tidy

# --- Test & run ---
test:
	go test ./internal/navlogic/... -v

# Run in emulator (no board; uses default sim I/O)
run:
	tinygo run -target=arduino-nano .

# --- Clean & help ---
# Windows: pwsh. Unix: rm -f
ifeq ($(OS),Windows_NT)
clean:
	-@pwsh -NoProfile -Command "Remove-Item -Force -ErrorAction SilentlyContinue '$(FIRMWARE)'"
else
clean:
	-rm -f $(FIRMWARE)
endif

help:
	@echo "Tiny Pet — build, flash, format"
	@echo ""
	@echo "Targets:"
	@echo "  build, build-uno   Build for Arduino Uno (default), output: $(FIRMWARE)"
	@echo "  build-nano         Build for Arduino Nano"
	@echo "  flash              Flash Uno (PORT= auto-detected; on Windows uses pwsh, set PORT=COM3 if needed)"
	@echo "  flash-nano         Flash Nano (same PORT= as flash)"
	@echo "  fmt                Format Go code (go fmt + gofmt -s -w)"
	@echo "  tidy               go mod tidy"
	@echo "  test               Run unit tests (internal/navlogic)"
	@echo "  run                Run in emulator (tinygo run, no board)"
	@echo "  clean              Remove $(FIRMWARE)"
	@echo "  help               This message"
	@echo ""
	@echo "Examples:"
	@echo "  make build"
	@echo "  make build-nano"
	@echo "  make flash"
	@echo "  make build-nano flash-nano"
	@echo "  make flash PORT=/dev/cu.usbmodem14101   # macOS"
	@echo "  make flash PORT=/dev/ttyACM0            # Linux"
	@echo "  make flash PORT=COM3                    # Windows"
	@echo "  make fmt tidy test"
