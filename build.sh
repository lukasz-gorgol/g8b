#!/bin/bash

# Build script for the Galoping Top 8-bit emulator and assembler
# This script provides an easy way to build the project

set -e

# Print colored output
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

# Print usage information
function print_usage {
    echo -e "${YELLOW}Usage:${NC} $0 [command]"
    echo ""
    echo "Commands:"
    echo "  all       Build both emulator and assembler (default)"
    echo "  emulator  Build only the emulator"
    echo "  assembler Build only the assembler"
    echo "  clean     Remove build artifacts"
    echo "  release   Build optimized release binaries"
    echo "  bench     Run benchmarks"
    echo "  test      Run tests"
    echo "  profile   Build with profiling enabled"
    echo "  help      Show this help message"
    echo ""
}

# Check if make is installed
if ! command -v make &> /dev/null; then
    echo -e "${RED}Error: make is not installed.${NC}"
    echo "Please install make using your package manager:"
    echo "  Ubuntu/Debian: sudo apt-get install make"
    echo "  Fedora/RHEL:   sudo dnf install make"
    echo "  macOS:         brew install make"
    exit 1
fi

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo -e "${RED}Error: Go is not installed.${NC}"
    echo "Please install Go from https://golang.org/dl/"
    exit 1
fi

# Make the script executable
chmod +x "$0"

# Process command line arguments
CMD=${1:-all}

case "$CMD" in
    all)
        echo -e "${GREEN}Building emulator and assembler...${NC}"
        make all
        echo -e "${GREEN}Build successful! Binaries are in the bin/ directory.${NC}"
        ;;
    emulator)
        echo -e "${GREEN}Building emulator...${NC}"
        make bin/emulator
        echo -e "${GREEN}Build successful! Binary is at bin/emulator.${NC}"
        ;;
    assembler)
        echo -e "${GREEN}Building assembler...${NC}"
        make bin/assembler
        echo -e "${GREEN}Build successful! Binary is at bin/assembler.${NC}"
        ;;
    clean)
        echo -e "${GREEN}Cleaning build artifacts...${NC}"
        make clean
        echo -e "${GREEN}Clean successful!${NC}"
        ;;
    release)
        echo -e "${GREEN}Building release binaries...${NC}"
        make release
        echo -e "${GREEN}Build successful! Binaries are in the dist/ directory.${NC}"
        ;;
    bench)
        echo -e "${GREEN}Running benchmarks...${NC}"
        make bench
        ;;
    test)
        echo -e "${GREEN}Running tests...${NC}"
        make test
        ;;
    profile)
        echo -e "${GREEN}Building with profiling enabled...${NC}"
        make profile
        echo -e "${GREEN}Build successful! Profiling-enabled binaries are in the bin/ directory.${NC}"
        echo -e "${YELLOW}To use profiling:${NC}"
        echo "1. Run the binary: bin/emulator-profile [args]"
        echo "2. The profile will be saved to cpu.pprof"
        echo "3. Analyze with: go tool pprof cpu.pprof"
        ;;
    help)
        print_usage
        ;;
    *)
        echo -e "${RED}Unknown command: $CMD${NC}"
        print_usage
        exit 1
        ;;
esac

exit 0 