package cpu

import (
	"time"
)

// AddressingMode represents the addressing mode of an instruction
type AddressingMode string

const (
	Immediate AddressingMode = "IMMEDIATE" // #$value
	Absolute  AddressingMode = "ABSOLUTE"  // $addr
	Relative  AddressingMode = "RELATIVE"  // label
	Implied   AddressingMode = "IMPLIED"   // no operand
)

// Instruction represents a CPU instruction with all its properties
type Instruction struct {
	Opcode      byte           // The opcode byte value
	Mnemonic    string         // The instruction mnemonic
	Mode        AddressingMode // The addressing mode
	Size        int            // Size in bytes
	Cycles      int            // Number of cycles
	Description string         // Description of the instruction
}

// CPU interface defines the methods that any CPU implementation must provide
type ICPU interface {
	// Base operations
	GetName() string
	GetInstructions() map[byte]Instruction

	// Core CPU operations
	Run()
	ExecuteInstruction()
	GetElapsedTime() time.Duration
	GetCyclesPerSecond() float64

	// Memory operations
	Read(addr uint16) byte
	Write(addr uint16, value byte)

	// Register operations
	GetPC() uint16
	SetPC(addr uint16)
	GetSP() uint8
	SetSP(value uint8)
	GetA() uint8
	SetA(value uint8)
	GetX() uint8
	SetX(value uint8)
	GetY() uint8
	SetY(value uint8)

	// Stack operations
	Push(value byte)
	Pull() byte
	Push16(value uint16)
	Pull16() uint16

	// Timing operations
	WaitForCycles(cycles int)
	GetCycles() int
	AddCycles(cycles int)

	// Verbose operations
	SetVerbose(verbose bool)
	IsVerbose() bool
}

// BaseCPU provides common functionality for all CPU implementations
type CPU struct {
	Name         string // CPU name
	Instructions map[byte]Instruction
	PC           uint16    // Program Counter
	SP           uint8     // Stack Pointer
	Memory       []uint8   // Memory
	Cycles       int       // Cycle counter
	Speed        uint      // CPU speed in Hz
	startTime    time.Time // Start time for timing
	stopTime     time.Time // Stop time for timing
	running      bool      // Whether the CPU is currently running
	verbose      bool      // Enable verbose output
}

func (c CPU) GetName() string {
	return c.Name
}

func (c CPU) GetInstructions() map[byte]Instruction {
	return c.Instructions
}

// NewBaseCPU creates a new base CPU instance
func NewCPU(name string, memorySize int, speed uint, instructions map[byte]Instruction) *CPU {
	return &CPU{
		Name:         name,
		Instructions: instructions,
		Memory:       make([]uint8, memorySize),
		SP:           0xFF, // Initialize stack pointer to top of stack
		Speed:        speed,
	}
}

// Read reads a byte from memory
func (c *CPU) Read(addr uint16) byte {
	return c.Memory[addr]
}

// Write writes a byte to memory
func (c *CPU) Write(addr uint16, value byte) {
	c.Memory[addr] = value
}

// GetPC returns the program counter
func (c *CPU) GetPC() uint16 {
	return c.PC
}

// SetPC sets the program counter
func (c *CPU) SetPC(addr uint16) {
	c.PC = addr
}

// GetSP returns the stack pointer
func (c *CPU) GetSP() uint8 {
	return c.SP
}

// SetSP sets the stack pointer
func (c *CPU) SetSP(value uint8) {
	c.SP = value
}

// GetA returns the accumulator
func (c *CPU) GetA() uint8 {
	return 0 // To be implemented by specific CPU
}

// SetA sets the accumulator
func (c *CPU) SetA(value uint8) {
	// To be implemented by specific CPU
}

// GetX returns the X register
func (c *CPU) GetX() uint8 {
	return 0 // To be implemented by specific CPU
}

// SetX sets the X register
func (c *CPU) SetX(value uint8) {
	// To be implemented by specific CPU
}

// GetY returns the Y register
func (c *CPU) GetY() uint8 {
	return 0 // To be implemented by specific CPU
}

// SetY sets the Y register
func (c *CPU) SetY(value uint8) {
	// To be implemented by specific CPU
}

// Push pushes a byte onto the stack
func (c *CPU) Push(value byte) {
	c.Memory[0x0100+uint16(c.SP)] = value
	c.SP--
}

// Pull pulls a value from the stack
func (c *CPU) Pull() byte {
	c.SP++
	return c.Memory[0x0100+uint16(c.SP)]
}

// Push16 pushes a 16-bit value onto the stack
func (c *CPU) Push16(value uint16) {
	high := byte(value >> 8)
	low := byte(value & 0xFF)
	c.Push(high)
	c.Push(low)
}

// Pull16 pulls a 16-bit value from the stack
func (c *CPU) Pull16() uint16 {
	low := uint16(c.Pull())
	high := uint16(c.Pull())
	return (high << 8) | low
}

// Run starts the CPU execution
func (c *CPU) Run() {
	c.startTime = time.Now()
	c.running = true
	c.Cycles = 0
}

// Stops the CPU execution
func (c *CPU) Stop() {
	if c.running {
		c.stopTime = time.Now()
		c.running = false
	}
}

// WaitForCycles waits for the appropriate amount of time based on CPU speed
func (c *CPU) WaitForCycles(cycles int) {
	c.Cycles += cycles
	if !c.running || c.Speed == 0 {
		return // No timing if not running or speed is 0
	}

	// Calculate how long we should have taken so far
	targetElapsed := time.Duration(float64(c.Cycles) * float64(time.Second) / float64(c.Speed))
	actualElapsed := time.Since(c.startTime)

	if actualElapsed < targetElapsed {
		delta := targetElapsed - actualElapsed
		if delta > time.Millisecond {
			time.Sleep(delta)
		} else {
			for targetElapsed-time.Since(c.startTime) > 0 {
				for start := time.Now().UnixNano(); time.Now().UnixNano() == start; {
				}
			}
		}
	}
}

// GetElapsedTime returns the elapsed time since the CPU started running
func (c *CPU) GetElapsedTime() time.Duration {
	if !c.running {
		return c.stopTime.Sub(c.startTime)
	}
	return time.Since(c.startTime)
}

// GetCyclesPerSecond returns the actual cycles per second achieved
func (c *CPU) GetCyclesPerSecond() float64 {
	elapsed := c.GetElapsedTime().Seconds()
	if elapsed == 0 {
		return 0
	}
	return float64(c.Cycles) / elapsed
}

// GetCycles returns the current cycle count
func (c *CPU) GetCycles() int {
	return c.Cycles
}

// AddCycles adds cycles to the cycle counter
func (c *CPU) AddCycles(cycles int) {
	c.Cycles += cycles
}

// SetVerbose enables or disables verbose output
func (c *CPU) SetVerbose(verbose bool) {
	c.verbose = verbose
}

// IsVerbose returns whether verbose output is enabled
func (c *CPU) IsVerbose() bool {
	return c.verbose
}

// GetOpcodes returns a map of mnemonic to opcode for the instruction set
func (c *CPU) GetOpcodes() map[string]byte {
	opcodes := make(map[string]byte)
	for opcode, instr := range c.GetInstructions() {
		opcodes[instr.Mnemonic] = opcode
	}
	return opcodes
}

// GetInstrSizes returns a map of mnemonic to instruction size for the instruction set
func (c *CPU) GetInstrSizes() map[string]int {
	sizes := make(map[string]int)
	for _, instr := range c.GetInstructions() {
		sizes[instr.Mnemonic] = instr.Size
	}
	return sizes
}
