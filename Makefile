```makefile
# Project name
APP_NAME := TCPChat

# Main server entry point
MAIN_PATH := ./cmd/server

# Build output directory
BIN_DIR := bin

# Final binary path
BINARY := $(BIN_DIR)/$(APP_NAME)

# Default port used by the project
PORT ?= 8989

# Go command
GO := go

# Go files used for formatting/lint-style checks
GO_FILES := $(shell find . -name "*.go")

# Default target
.PHONY: all
all: fmt test build

# Run the server using the default port: 8989
.PHONY: run
run:
	$(GO) run $(MAIN_PATH)

# Run the server with a custom port.
# Usage:
# make run-port PORT=2525
.PHONY: run-port
run-port:
	$(GO) run $(MAIN_PATH) $(PORT)

# Build the application binary
.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	$(GO) build -o $(BINARY) $(MAIN_PATH)

# Build and run the compiled binary using the default port
.PHONY: start
start: build
	./$(BINARY)

# Build and run the compiled binary using a custom port.
# Usage:
# make start-port PORT=2525
.PHONY: start-port
start-port: build
	./$(BINARY) $(PORT)

# Run all tests in the project
.PHONY: test
test:
	$(GO) test ./...

# Run tests with verbose output
.PHONY: test-v
test-v:
	$(GO) test -v ./...

# Run tests with race detector.
# Useful for checking unsafe concurrent access in goroutines.
.PHONY: race
race:
	$(GO) test -race ./...

# Format all Go source files
.PHONY: fmt
fmt:
	$(GO) fmt ./...

# Check if all Go files are formatted.
# This is useful before submission because auditors may check formatting.
.PHONY: fmt-check
fmt-check:
	@test -z "$$($(GO)fmt -l $(GO_FILES))" || \
	(echo "These files are not formatted:"; $(GO)fmt -l $(GO_FILES); exit 1)

# Run go vet to catch suspicious code
.PHONY: vet
vet:
	$(GO) vet ./...

# Run formatting, vetting, and tests
.PHONY: check
check: fmt vet test

# Run stronger checks before submitting
.PHONY: audit
audit: fmt-check vet race test

# Remove build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

# Rebuild the project from scratch
.PHONY: rebuild
rebuild: clean build

# Download and tidy dependencies
.PHONY: tidy
tidy:
	$(GO) mod tidy

# Display the current project structure
.PHONY: tree
tree:
	@find . -not -path "./.git/*" -not -path "./bin/*" | sort

# Show helpful project commands
.PHONY: help
help:
	@echo "Available commands:"
	@echo "  make              Run fmt, test, and build"
	@echo "  make run          Run server on default port 8989"
	@echo "  make run-port     Run server on custom port, example: make run-port PORT=2525"
	@echo "  make build        Build binary into bin/TCPChat"
	@echo "  make start        Build and run binary on default port"
	@echo "  make start-port   Build and run binary on custom port"
	@echo "  make test         Run all tests"
	@echo "  make test-v       Run tests with verbose output"
	@echo "  make race         Run tests with race detector"
	@echo "  make fmt          Format Go files"
	@echo "  make fmt-check    Check Go formatting without changing files"
	@echo "  make vet          Run go vet"
	@echo "  make check        Run fmt, vet, and test"
	@echo "  make audit        Run stricter pre-submission checks"
	@echo "  make clean        Remove build files"
	@echo "  make rebuild      Clean and build again"
	@echo "  make tidy         Clean up go.mod and go.sum"
	@echo "  make tree         Show project structure"
```
