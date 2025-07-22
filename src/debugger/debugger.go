package debugger

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/lukasz-gorgol/g8b/src/cpu"
)

// Debugger represents the debugger state
type Debugger struct {
	cpu         cpu.ICPU
	breakpoints map[uint16]bool
	running     bool
	stepMode    bool
	lastPC      uint16
}

// New creates a new debugger instance
func New(cpu cpu.ICPU) *Debugger {
	return &Debugger{
		cpu:         cpu,
		breakpoints: make(map[uint16]bool),
		running:     true,
		stepMode:    false,
		lastPC:      cpu.GetPC(),
	}
}

// Run starts the debugger's main loop
func (d *Debugger) Run() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Galoping Top Debugger - Type 'help' for commands")

	for d.running {
		fmt.Print("(debug) ")
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}

		cmd := strings.ToLower(parts[0])
		args := parts[1:]

		switch cmd {
		case "help", "h":
			d.printHelp()
		case "break", "b":
			d.handleBreakpoint(args)
		case "run", "r":
			d.run()
		case "step", "s":
			d.step()
		case "continue", "c":
			d.continueExecution()
		case "registers", "reg":
			d.printRegisters()
		case "memory", "m":
			d.printMemory(args)
		case "disassemble", "d":
			d.disassemble(args)
		case "watch", "w":
			d.handleWatch(args)
		case "quit", "q":
			d.running = false
		default:
			fmt.Printf("Unknown command: %s\n", cmd)
		}
	}
}

// printHelp displays available commands
func (d *Debugger) printHelp() {
	fmt.Println("Available commands:")
	fmt.Println("  help, h              - Show this help")
	fmt.Println("  break, b <addr>      - Set breakpoint at address")
	fmt.Println("  run, r               - Run until breakpoint or end")
	fmt.Println("  step, s              - Execute one instruction")
	fmt.Println("  continue, c          - Continue execution")
	fmt.Println("  registers, reg       - Show CPU registers")
	fmt.Println("  memory, m <addr>     - Show memory at address")
	fmt.Println("  disassemble, d <addr>- Disassemble code at address")
	fmt.Println("  watch, w <addr>      - Watch memory location")
	fmt.Println("  quit, q              - Exit debugger")
}

// handleBreakpoint sets or removes a breakpoint
func (d *Debugger) handleBreakpoint(args []string) {
	if len(args) == 0 {
		fmt.Println("Breakpoints:")
		for addr := range d.breakpoints {
			fmt.Printf("  $%04X\n", addr)
		}
		return
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	if d.breakpoints[addr] {
		delete(d.breakpoints, addr)
		fmt.Printf("Removed breakpoint at $%04X\n", addr)
	} else {
		d.breakpoints[addr] = true
		fmt.Printf("Set breakpoint at $%04X\n", addr)
	}
}

// run executes the program until a breakpoint or end
func (d *Debugger) run() {
	d.stepMode = false
	for {
		if d.breakpoints[d.cpu.GetPC()] {
			fmt.Printf("Breakpoint hit at $%04X\n", d.cpu.GetPC())
			return
		}

		opcode := d.cpu.Read(d.cpu.GetPC())
		if opcode == 0x00 {
			fmt.Println("Program halted")
			return
		}

		d.executeInstruction()
	}
}

// step executes a single instruction
func (d *Debugger) step() {
	d.executeInstruction()
}

// continueExecution continues program execution
func (d *Debugger) continueExecution() {
	d.stepMode = false
	d.run()
}

// executeInstruction executes one instruction and prints debug info
func (d *Debugger) executeInstruction() {
	opcode := d.cpu.Read(d.cpu.GetPC())
	instruction, ok := d.cpu.GetInstructions()[opcode]
	var mnemonic string
	if ok {
		mnemonic = instruction.Mnemonic
	} else {
		mnemonic = "???"
	}

	// Print CPU-specific debug info
	switch c := d.cpu.(type) {
	case *cpu.Intel8008:
		fmt.Printf("PC: $%04X | Opcode: $%02X %-3s | A: $%02X B: $%02X C: $%02X | Flags(CZSP): %d%d%d%d\n",
			c.GetPC(), opcode, mnemonic, c.GetA(), c.GetB(), c.GetC(),
			boolToInt(c.Flags.Carry),
			boolToInt(c.Flags.Zero),
			boolToInt(c.Flags.Sign),
			boolToInt(c.Flags.Parity))
	default:
		fmt.Printf("PC: $%04X | Opcode: $%02X %-3s\n",
			d.cpu.GetPC(), opcode, mnemonic)
	}

	d.lastPC = d.cpu.GetPC()
	d.cpu.ExecuteInstruction()

	if d.cpu.GetPC() == d.lastPC {
		fmt.Println("Program halted")
		d.running = false
	}
}

// printRegisters displays CPU register values
func (d *Debugger) printRegisters() {
	fmt.Printf("PC: $%04X\n", d.cpu.GetPC())
	fmt.Printf("SP: $%02X\n", d.cpu.GetSP())

	// Print CPU-specific registers
	switch c := d.cpu.(type) {
	case *cpu.Intel8008:
		fmt.Printf("A:  $%02X\n", c.GetA())
		fmt.Printf("B:  $%02X\n", c.GetB())
		fmt.Printf("C:  $%02X\n", c.GetC())
		fmt.Printf("D:  $%02X\n", c.GetD())
		fmt.Printf("E:  $%02X\n", c.GetE())
		fmt.Printf("H:  $%02X\n", c.GetH())
		fmt.Printf("L:  $%02X\n", c.GetL())
		fmt.Printf("Flags: C:%d Z:%d S:%d P:%d\n",
			boolToInt(c.Flags.Carry),
			boolToInt(c.Flags.Zero),
			boolToInt(c.Flags.Sign),
			boolToInt(c.Flags.Parity))
	default:
		fmt.Printf("A:  $%02X\n", d.cpu.GetA())
	}
}

// printMemory displays memory contents
func (d *Debugger) printMemory(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: memory <address>")
		return
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	fmt.Printf("Memory at $%04X:\n", addr)
	for i := 0; i < 16; i++ {
		fmt.Printf("$%04X: $%02X\n", addr+uint16(i), d.cpu.Read(addr+uint16(i)))
	}
}

// disassemble displays disassembled code
func (d *Debugger) disassemble(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: disassemble <address>")
		return
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	fmt.Printf("Disassembly at $%04X:\n", addr)
	for i := 0; i < 10; i++ {
		opcode := d.cpu.Read(addr)
		instruction, ok := d.cpu.GetInstructions()[opcode]
		var mnemonic string
		if ok {
			mnemonic = instruction.Mnemonic
		} else {
			mnemonic = "???"
		}
		fmt.Printf("$%04X: $%02X %s\n", addr, opcode, mnemonic)
		addr++
	}
}

// handleWatch sets or removes a memory watch
func (d *Debugger) handleWatch(args []string) {
	if len(args) == 0 {
		fmt.Println("Usage: watch <address>")
		return
	}

	addr, err := parseAddress(args[0])
	if err != nil {
		fmt.Printf("Invalid address: %v\n", err)
		return
	}

	fmt.Printf("Watching memory at $%04X: $%02X\n", addr, d.cpu.Read(addr))
}

// parseAddress parses a hex address string
func parseAddress(s string) (uint16, error) {
	s = strings.TrimPrefix(s, "0x")
	value, err := strconv.ParseUint(s, 16, 16)
	if err != nil {
		return 0, err
	}
	return uint16(value), nil
}

// boolToInt converts a boolean to an integer (0 or 1)
func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
