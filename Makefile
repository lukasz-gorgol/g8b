# Makefile for optimized Galoping Top 8-bit emulator and assembler

# Build variables
GO := go
GOFLAGS := -v
LDFLAGS := -ldflags="-s -w"
BUILDFLAGS := -trimpath
OPTIMIZED_FLAGS := $(BUILDFLAGS) $(LDFLAGS) -tags=netgo
RELEASE_FLAGS := $(OPTIMIZED_FLAGS) -ldflags="-s -w -X main.Version=$(shell git describe --tags --always)"

# Output directories
BIN_DIR := bin
DIST_DIR := dist

# Binary names
EMULATOR := $(BIN_DIR)/emulator
ASSEMBLER := $(BIN_DIR)/assembler

# Source files
EMULATOR_SRC := src/emulator/emulator.go
ASSEMBLER_SRC := src/assembler/assembler.go
CPU_SRC := $(wildcard src/cpu/*.go)

# Default target
.PHONY: all
all: $(EMULATOR) $(ASSEMBLER)

# Ensure directories exist
$(BIN_DIR):
	mkdir -p $(BIN_DIR)

$(DIST_DIR):
	mkdir -p $(DIST_DIR)

# Build emulator with optimizations
$(EMULATOR): $(EMULATOR_SRC) $(CPU_SRC) | $(BIN_DIR)
	$(GO) build $(GOFLAGS) $(OPTIMIZED_FLAGS) -o $@ $<

# Build assembler with optimizations
$(ASSEMBLER): $(ASSEMBLER_SRC) $(CPU_SRC) | $(BIN_DIR)
	$(GO) build $(GOFLAGS) $(OPTIMIZED_FLAGS) -o $@ $<

# Clean build artifacts
.PHONY: clean
clean:
	rm -rf $(BIN_DIR) $(DIST_DIR)

# Build release versions with maximum optimization
.PHONY: release
release: | $(DIST_DIR)
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/emulator-linux-amd64 $(EMULATOR_SRC)
	GOOS=linux GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/assembler-linux-amd64 $(ASSEMBLER_SRC)
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/emulator-darwin-amd64 $(EMULATOR_SRC)
	GOOS=darwin GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/assembler-darwin-amd64 $(ASSEMBLER_SRC)
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/emulator-windows-amd64.exe $(EMULATOR_SRC)
	GOOS=windows GOARCH=amd64 $(GO) build $(GOFLAGS) $(RELEASE_FLAGS) -o $(DIST_DIR)/assembler-windows-amd64.exe $(ASSEMBLER_SRC)

# Run benchmarks
.PHONY: bench
bench:
	$(GO) test -bench=. -benchmem ./src/cpu/...

# Run tests
.PHONY: test
test:
	$(GO) test ./src/...

# Install binaries to $GOPATH/bin
.PHONY: install
install:
	$(GO) install $(GOFLAGS) $(OPTIMIZED_FLAGS) $(EMULATOR_SRC)
	$(GO) install $(GOFLAGS) $(OPTIMIZED_FLAGS) $(ASSEMBLER_SRC)

# Build with profiling enabled
.PHONY: profile
profile: | $(BIN_DIR)
	$(GO) build $(GOFLAGS) -tags=profile $(OPTIMIZED_FLAGS) -o $(BIN_DIR)/emulator-profile $(EMULATOR_SRC)
	$(GO) build $(GOFLAGS) -tags=profile $(OPTIMIZED_FLAGS) -o $(BIN_DIR)/assembler-profile $(ASSEMBLER_SRC)

# Help target
.PHONY: help
help:
	@echo "Available targets:"
	@echo "  all      - Build emulator and assembler binaries (default)"
	@echo "  clean    - Remove build artifacts"
	@echo "  release  - Build optimized release binaries for multiple platforms"
	@echo "  bench    - Run benchmarks"
	@echo "  test     - Run tests"
	@echo "  install  - Install binaries to GOPATH/bin"
	@echo "  profile  - Build with profiling enabled" 