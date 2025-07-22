package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/lukasz-gorgol/g8b/src/cpu"
	"github.com/lukasz-gorgol/g8b/src/debugger"
)

// Config represents the emulator configuration
type Config struct {
	Binary     string `json:"binary"`                // Path to the binary file
	StartAddr  string `json:"start_addr,omitempty"`  // Start address as hex string (e.g., "0x8000")
	MemorySize uint   `json:"memory_size,omitempty"` // Memory size in bytes (default: 65536)
	DumpAddrs  string `json:"dump_addrs,omitempty"`  // Memory addresses to dump
	CPUType    string `json:"cpu,omitempty"`         // CPU type (default: 8008)
	CPUSpeed   uint   `json:"speed,omitempty"`       // CPU speed in Hz (default: 1000000 for 1MHz)
	Verbose    bool   `json:"verbose,omitempty"`     // Enable verbose output
}

func main() {
	// Define command-line flags
	configFile := flag.String("c", "", "Path to JSON configuration file")
	startAddr := flag.String("s", "0x8000", "Start address for program loading and PC initialization (hex string)")
	memorySize := flag.Uint("m", 65536, "Memory size in bytes")
	dumpAddrs := flag.String("d", "", "Memory addresses to dump")
	cpuType := flag.String("cpu", "8008", "CPU type (default: 8008)")
	cpuSpeed := flag.Uint("speed", 1000000, "CPU speed in Hz (default: 1000000 for 1MHz)")
	debug := flag.Bool("debug", false, "Run in debug mode")
	verbose := flag.Bool("v", false, "Enable verbose output (show PC, registers, and flags)")
	flag.Parse()

	// Parse command-line arguments
	args := flag.Args()
	if len(args) != 1 && *configFile == "" {
		fmt.Println("Usage: ./bin/emulator [options] <program.bin>")
		fmt.Println("\nOptions:")
		fmt.Println("  -c <file>    Path to JSON configuration file")
		fmt.Println("  -s <addr>    Start address (hex string, e.g., 0x8000)")
		fmt.Println("  -m <size>    Memory size in bytes (default: 65536)")
		fmt.Println("  -d <addrs>   Memory addresses to dump")
		fmt.Println("  -cpu <type>  CPU type (default: 8008)")
		fmt.Println("  -speed <hz>  CPU speed in Hz (default: 1000000 for 1MHz)")
		fmt.Println("  -debug       Run in debug mode")
		fmt.Println("  -v           Enable verbose output")
		os.Exit(1)
	}

	fmt.Println("✅ All systems go! Emulator starting...")

	var config Config

	// Load configuration from file if specified
	if *configFile != "" {
		file, err := os.Open(*configFile)
		if err != nil {
			fmt.Printf("🆘 Error opening config file: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		decoder := json.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			fmt.Printf("🆘 Error parsing config file: %v\n", err)
			os.Exit(1)
		}
	} else {
		// If no config file is provided, use command line arguments
		if flag.NArg() < 1 {
			fmt.Println("🆘 Error: Binary file required")
			flag.Usage()
			os.Exit(1)
		}

		args := flag.Args()
		config = Config{
			Binary:     args[0],
			StartAddr:  *startAddr,
			MemorySize: *memorySize,
			DumpAddrs:  *dumpAddrs,
			CPUType:    *cpuType, // Use the command line CPU type
			CPUSpeed:   *cpuSpeed,
			Verbose:    *verbose, // Use command line verbose flag
		}
	}

	// Set default CPU type if not specified in config file
	if config.CPUType == "" {
		config.CPUType = *cpuType // Use command line CPU type as default
		if config.CPUType == "" {
			config.CPUType = "8008" // Fallback to 8008 if still empty
		}
	}

	// Set default CPU speed if not specified
	if config.CPUSpeed == 0 {
		config.CPUSpeed = *cpuSpeed // Use command line CPU speed as default
		if config.CPUSpeed == 0 {
			config.CPUSpeed = 1000000 // Fallback to 1MHz if still 0
		}
	}

	// Set verbose flag if not specified in config file
	if !config.Verbose {
		config.Verbose = *verbose // Use command line verbose flag as default
	}

	// Display current configuration
	fmt.Println("\n📋 Current Configuration:")
	fmt.Printf("  Binary:      %s\n", config.Binary)
	fmt.Printf("  Start Addr:  %s\n", config.StartAddr)
	fmt.Printf("  Memory Size: %d bytes\n", config.MemorySize)
	fmt.Printf("  CPU Type:    %s\n", config.CPUType)
	fmt.Printf("  CPU Speed:   %d Hz\n", config.CPUSpeed)
	if config.DumpAddrs != "" {
		fmt.Printf("  Dump Addrs:  %s\n", config.DumpAddrs)
	}
	fmt.Printf("  Verbose:     %v\n", config.Verbose)
	fmt.Println()

	// Parse start address
	startAddress, err := parseHexAddr(config.StartAddr, 0x8000)
	if err != nil {
		fmt.Printf("🆘 Error parsing start address: %v\n", err)
		os.Exit(1)
	}

	// Create CPU instance
	var processor cpu.ICPU
	switch config.CPUType {
	case "8008":
		processor = cpu.NewIntel8008(int(config.MemorySize), config.CPUSpeed)
	default:
		fmt.Printf("🆘 Unsupported CPU type: %s\n", config.CPUType)
		fmt.Println("  Available CPU types: 8008")
		os.Exit(1)
	}

	// Set verbose mode on CPU if enabled
	if config.Verbose {
		processor.SetVerbose(true)
	}

	// Load program
	program, err := os.ReadFile(config.Binary)
	if err != nil {
		fmt.Printf("🆘 Error reading binary file: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("✅ Binary loaded successfully: %s (%d bytes)\n", config.Binary, len(program))

	// Copy program to memory
	for i, b := range program {
		processor.Write(uint16(startAddress)+uint16(i), b)
	}
	processor.SetPC(uint16(startAddress))

	if *debug {
		// Run in debug mode
		fmt.Println("\n▶️🔍 Entering debug mode...")
		dbg := debugger.New(processor)

		startTime := time.Now()
		dbg.Run()
		endTime := time.Now()

		// Calculate execution statistics
		duration := endTime.Sub(startTime)

		fmt.Println("⏹️ Emulation finished.")
		fmt.Printf("  Execution completed in %v\n", duration)
		fmt.Printf("  Total cycles:  %d\n", processor.GetCycles())

	} else {
		// Run in normal mode
		fmt.Println("\n▶️  Executing program...")

		processor.Run()

		// Calculate execution statistics
		duration := processor.GetElapsedTime()
		cyclesPerSecond := float64(processor.GetCycles()) / duration.Seconds()

		fmt.Println("⏹️  Emulation finished.")
		fmt.Printf("  Execution completed in %v\n", duration)
		fmt.Printf("  Total cycles:  %d\n", processor.GetCycles())
		fmt.Printf("  Average speed: %.2f Hz (%.2f%% of target)\n",
			cyclesPerSecond,
			(cyclesPerSecond/float64(config.CPUSpeed))*100)
	}

	// Dump specified memory addresses
	if config.DumpAddrs != "" {
		addresses, err := parseAddressSpec(config.DumpAddrs)
		if err != nil {
			fmt.Printf("🆘 Error parsing address specification: %v\n", err)
			os.Exit(1)
		}

		fmt.Println("\n📝 Memory dump:")
		for _, addr := range addresses {
			fmt.Printf("  $%04X: $%02X\n", addr, processor.Read(addr))
		}
	}

	fmt.Println("\n✅ Emulator exiting...")
}

// parseHexAddr parses a hex address string
func parseHexAddr(addr string, defaultAddr uint16) (uint16, error) {
	if addr == "" {
		return defaultAddr, nil
	}
	addr = strings.TrimPrefix(addr, "0x")
	value, err := strconv.ParseUint(addr, 16, 16)
	if err != nil {
		return 0, err
	}
	return uint16(value), nil
}

// parseAddressSpec parses a memory address specification
func parseAddressSpec(spec string) ([]uint16, error) {
	var addresses []uint16
	parts := strings.Split(spec, ",")

	for _, part := range parts {
		if strings.Contains(part, "-") {
			// Handle range
			rangeParts := strings.Split(part, "-")
			if len(rangeParts) != 2 {
				return nil, fmt.Errorf("🆘 invalid range format: %s", part)
			}

			start, err := parseHexAddr(rangeParts[0], 0)
			if err != nil {
				return nil, err
			}

			end, err := parseHexAddr(rangeParts[1], 0)
			if err != nil {
				return nil, err
			}

			if start > end {
				return nil, fmt.Errorf("🆘 invalid range: start > end")
			}

			for addr := start; addr <= end; addr++ {
				addresses = append(addresses, addr)
			}
		} else {
			// Handle single address
			addr, err := parseHexAddr(part, 0)
			if err != nil {
				return nil, err
			}
			addresses = append(addresses, addr)
		}
	}

	return addresses, nil
}
