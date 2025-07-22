package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"

	"github.com/lukasz-gorgol/g8b/src/cpu"
)

// Config represents the assembler configuration
type Config struct {
	Source    string `json:"source,omitempty"`
	Binary    string `json:"binary,omitempty"`
	CPUType   string `json:"cpu,omitempty"`
	StartAddr string `json:"start_addr,omitempty"` // Start address as hex string (e.g., "0x8000")
	XXD       bool   `json:"xxd,omitempty"`
}

// assemble performs the two-pass assembly process
func assemble(inputFile *os.File, cpuType string, startAddress uint16) ([]byte, map[string]uint16) {
	labels := make(map[string]uint16)
	var currentAddress uint16 = startAddress
	var binary []byte

	// Get CPU-specific opcodes and instruction sizes
	var opcodes map[string]byte
	var instrSizes map[string]int

	switch cpuType {
	case "8008":
		opcodes = cpu.NewIntel8008(1, 1).GetOpcodes()
		instrSizes = cpu.NewIntel8008(1, 1).GetInstrSizes()
	default:
		fmt.Printf("🆘 Unsupported CPU type: %s\n", cpuType)
		fmt.Println("  Available CPU types: 8008")
		os.Exit(1)
	}

	// First pass: collect labels
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") {
			continue
		}

		if strings.HasSuffix(line, ":") {
			label := strings.TrimSuffix(line, ":")
			labels[label] = currentAddress
			fmt.Printf("Label: %v, address: $%04X\n", label, currentAddress)
			continue
		}

		parts := strings.Fields(line)
		mnemonic := parts[0]
		currentAddress += uint16(instrSizes[mnemonic])
	}

	// Second pass: generate binary
	inputFile.Seek(0, 0)
	scanner = bufio.NewScanner(inputFile)
	currentAddress = 0x8000

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" || strings.HasPrefix(line, ";") || strings.HasSuffix(line, ":") {
			continue
		}

		if idx := strings.Index(line, ";"); idx != -1 {
			line = line[:idx]
		}

		parts := strings.Fields(line)
		mnemonic := parts[0]

		opcode, exists := opcodes[mnemonic]
		if !exists {
			fmt.Printf("🆘 Unknown mnemonic: %s\n", mnemonic)
			os.Exit(1)
		}

		binary = append(binary, opcode)

		if len(parts) > 1 {
			operand := parts[1]
			switch cpuType {
			case "8008":
				handle8008Operand(mnemonic, operand, labels, currentAddress, &binary)
			}
		}
		currentAddress += uint16(instrSizes[mnemonic])
	}

	return binary, labels
}

func handle8008Operand(mnemonic, operand string, labels map[string]uint16, currentAddress uint16, binary *[]byte) {
	if strings.HasPrefix(mnemonic, "J") ||
		strings.HasPrefix(mnemonic, "CA") ||
		strings.HasPrefix(mnemonic, "CF") ||
		strings.HasPrefix(mnemonic, "CT") {
		// Handle absolute instructions
		var address uint16
		if strings.HasPrefix(operand, "$") {
			address64, _ := strconv.ParseUint(operand[1:], 16, 16)
			address = uint16(address64)
		} else if addr, ok := labels[operand]; ok {
			address = addr
		} else {
			fmt.Printf("🆘 Unknown label or address: %s\n", operand)
			os.Exit(1)
		}
		*binary = append(*binary, byte(address&0xFF), byte(address>>8))
	} else if strings.HasSuffix(mnemonic, "I") {
		// Handle immediate instructions
		if strings.HasPrefix(operand, "#$") {
			value, _ := strconv.ParseUint(operand[2:], 16, 16)
			if (mnemonic == "LHI" || mnemonic == "LLI") && value > 0xFF {
				fmt.Printf("🆘 Warning: %s only loads 8 bits, got #$%X (truncated to #$%02X)\n", mnemonic, value, value&0xFF)
				os.Exit(1)
			}
			*binary = append(*binary, byte(value))
		} else {
			fmt.Printf("🆘 Invalid immediate value format for %s: %s\n", mnemonic, operand)
			os.Exit(1)
		}
	} else {
		fmt.Printf("🆘 Error: can't assembly %s: %s\n", mnemonic, operand)
		os.Exit(1)
	}
}

func main() {
	// Define command-line flags
	configFile := flag.String("c", "", "Path to JSON configuration file")
	cpuType := flag.String("cpu", "8008", "CPU type (default: 8008)")
	startAddr := flag.String("s", "0x8000", "Start address for program loading and PC initialization (hex string)")
	xxdFlag := flag.Bool("xxd", false, "Run xxd on output binary after assembly")
	flag.Parse()

	// Parse command-line arguments
	args := flag.Args()
	if len(args) != 2 && *configFile == "" {
		fmt.Println("Usage: go run assembler.go [options] <source.asm> <binary.bin>")
		fmt.Println("\nOptions:")
		fmt.Println("  -c <file>    Path to JSON configuration file")
		fmt.Println("  -s <addr>    Start address (hex string, e.g., 0x8000)")
		fmt.Println("  -cpu <type>  CPU type (default: 8008)")
		fmt.Println("  -xxd         Run xxd on output binary after assembly")
		os.Exit(1)
	}

	fmt.Println("✅ All systems go! Assembling starting...")

	var config Config

	if *configFile != "" {
		// Load configuration from file
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
		if flag.NArg() < 2 {
			fmt.Println("🆘 Error: Assembler and binary files required")
			flag.Usage()
			os.Exit(1)
		}

		args := flag.Args()
		config = Config{
			Source:    args[0],
			Binary:    args[1],
			StartAddr: *startAddr,
			CPUType:   *cpuType,
			XXD:       *xxdFlag,
		}
	}

	// Set default CPU type if not specified in config file
	if config.CPUType == "" {
		config.CPUType = *cpuType // Use command line CPU type as default
		if config.CPUType == "" {
			config.CPUType = "8008" // Fallback to 8008 if still empty
		}
	}

	// If -xxd is set, override config file xxd
	if *xxdFlag {
		config.XXD = true
	}

	// Display current configuration
	fmt.Println("\n📋 Current Configuration:")
	fmt.Printf("  Source File: %s\n", config.Source)
	fmt.Printf("  Binary File: %s\n", config.Binary)
	fmt.Printf("  Start Addr:  %s\n", config.StartAddr)
	fmt.Printf("  CPU Type:    %s\n", config.CPUType)
	fmt.Printf("  XXD:         %v\n", config.XXD)
	fmt.Println()

	// Parse start address
	startAddress, err := parseHexAddr(config.StartAddr, 0x8000)
	if err != nil {
		fmt.Printf("🆘 Error parsing start address: %v\n", err)
		os.Exit(1)
	}

	// Open source file
	inputFile, err := os.Open(config.Source)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	// Create binary file
	outputFile, err := os.Create(config.Binary)
	if err != nil {
		panic(err)
	}
	defer outputFile.Close()

	// Assemble
	binary, _ := assemble(inputFile, config.CPUType, startAddress)
	outputFile.Write(binary)
	fmt.Printf("\n✅ Assembled successfully to %s using %s CPU\n", config.Binary, config.CPUType)

	// Optionally run xxd
	if config.XXD {
		fmt.Printf("\n▶️  Running xxd on %s...\n\n", config.Binary)
		cmd := exec.Command("xxd", config.Binary)
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		if err := cmd.Run(); err != nil {
			fmt.Printf("🆘 Error running xxd: %v\n", err)
		}
	}
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
