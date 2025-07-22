package cpu

import (
	"fmt"
	"os"
)

// CPU8008 represents the 8008 processor
type Intel8008 struct {
	CPU
	A     uint8 // Accumulator, A register, 0
	B     uint8 // B register, 1
	C     uint8 // C register, 2
	D     uint8 // D register, 3
	E     uint8 // E register, 4
	H     uint8 // High-order word, H register, 5
	L     uint8 // Low-order word, L register, 6
	Flags struct {
		Carry  bool // Carry Flag (C)
		Zero   bool // Zero Flag (Z)
		Sign   bool // Sign Flag (S)
		Parity bool // Parity Flag (P)
	}
}

// NewCPU creates a new 8008 CPU instance
func NewIntel8008(memorySize int, speed uint) *Intel8008 {
	return &Intel8008{
		CPU: *NewCPU("Intel8008", memorySize, speed, Intel8008Instructions),
	}
}

// GetA returns the accumulator
func (c *Intel8008) GetA() uint8 {
	return c.A
}

// SetA sets the accumulator
func (c *Intel8008) SetA(value uint8) {
	c.A = value
}

// GetB returns the B register
func (c *Intel8008) GetB() uint8 {
	return c.B
}

// SetB sets the B register
func (c *Intel8008) SetB(value uint8) {
	c.B = value
}

// GetC returns the C register
func (c *Intel8008) GetC() uint8 {
	return c.C
}

// SetC sets the C register
func (c *Intel8008) SetC(value uint8) {
	c.C = value
}

// GetD returns the D register
func (c *Intel8008) GetD() uint8 {
	return c.D
}

// SetD sets the D register
func (c *Intel8008) SetD(value uint8) {
	c.D = value
}

// GetE returns the E register
func (c *Intel8008) GetE() uint8 {
	return c.E
}

// SetE sets the E register
func (c *Intel8008) SetE(value uint8) {
	c.E = value
}

// GetH returns the H register
func (c *Intel8008) GetH() uint8 {
	return c.H
}

// SetH sets the H register
func (c *Intel8008) SetH(value uint8) {
	c.H = value
}

// GetL returns the L register
func (c *Intel8008) GetL() uint8 {
	return c.L
}

// SetL sets the L register
func (c *Intel8008) SetL(value uint8) {
	c.L = value
}

// Run executes the program starting at the current PC
func (c *Intel8008) Run() {
	// Start timing
	c.CPU.Run()

	for {
		// Get the opcode before executing
		opcode := c.Memory[c.PC]
		instruction, ok := c.Instructions[opcode]
		if !ok {
			fmt.Printf("🆘 Unknown opcode: $%02X at $%04X\n", opcode, c.PC)
			os.Exit(1) // Exit on unknown opcode
		}
		if instruction.Mnemonic == "HLT" {
			c.WaitForCycles(instruction.Cycles)
			break
		}

		// Execute one instruction
		c.ExecuteInstruction()
	}
	c.CPU.Stop()
}

// ExecuteInstruction executes a single instruction
func (c *Intel8008) ExecuteInstruction() {
	// Get the opcode
	opcode := c.Memory[c.PC]

	// Get the instruction
	instruction, ok := c.Instructions[opcode]
	if !ok {
		fmt.Printf("🆘 Unknown opcode: $%02X at $%04X\n", opcode, c.PC)
		os.Exit(1) // Exit on unknown opcode
	}

	// Only print verbose output if enabled
	if c.IsVerbose() {
		fmt.Printf("PC: %04X, OP: %02X, MN: %s, A:%02X B:%02X C:%02X D:%02X E:%02X H:%02X L:%02X | Flags(CZSP): %d%d%d%d\n",
			c.PC, opcode, instruction.Mnemonic, c.A, c.B, c.C, c.D, c.E, c.H, c.L,
			boolToInt(c.Flags.Carry), boolToInt(c.Flags.Zero), boolToInt(c.Flags.Sign), boolToInt(c.Flags.Parity))
	}

	// Execute the instruction based on its addressing mode
	switch instruction.Mode {
	case Immediate:
		value := c.Memory[c.PC+1]
		c.PC += 2
		c.executeImmediate(instruction, value)
	case Absolute:
		low := uint16(c.Memory[c.PC+1])
		high := uint16(c.Memory[c.PC+2])
		addr := low | (high << 8)
		c.PC += 3
		c.executeAbsolute(instruction, addr)
	case Relative:
		offset := int8(c.Memory[c.PC+1])
		c.PC += 2
		c.executeRelative(instruction, offset)
	case Implied:
		c.PC++
		c.executeImplied(instruction)
	}

	// Wait for the appropriate amount of time
	c.WaitForCycles(instruction.Cycles)
}

// executeImmediate executes an immediate mode instruction
func (c *Intel8008) executeImmediate(instruction Instruction, value byte) {
	switch instruction.Mnemonic {
	case "LAI":
		c.A = value
	case "LBI":
		c.B = value
	case "LCI":
		c.C = value
	case "LDI":
		c.D = value
	case "LEI":
		c.E = value
	case "LHI":
		c.H = value
	case "LLI":
		c.L = value
	case "LMI":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = value
	case "ADI":
		result := uint16(c.A) + uint16(value)
		c.A = byte(result)
		c.Flags.Carry = result > 0xFF
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "ACI":
		result := uint16(c.A) + uint16(value)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.Flags.Carry = result > 0xFF
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "SUI":
		result := uint16(c.A) - uint16(value)
		c.A = byte(result)
		c.Flags.Carry = result > 0xFF
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "SBI":
		result := uint16(c.A) - uint16(value)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.Flags.Carry = result > 0xFF
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "NDI":
		c.A &= value
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "XRI":
		c.A ^= value
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "ORI":
		c.A |= value
		c.Flags.Zero = c.A == 0
		c.Flags.Sign = (c.A & 0x80) != 0
		c.Flags.Parity = c.calculateParity(c.A)
	case "CPI":
		c.updateCompareFlags(c.A, value)
	default:
		fmt.Printf("🆘 Immediate instruction not implemented: %s\n", instruction.Mnemonic)
		os.Exit(1)
	}
}

// executeAbsolute executes an absolute mode instruction
func (c *Intel8008) executeAbsolute(instruction Instruction, addr uint16) {
	switch instruction.Mnemonic {
	case "JMP":
		c.PC = addr
	case "CAL":
		// Save return address on stack
		c.SP--
		c.Memory[c.SP] = byte(c.PC >> 8)
		c.SP--
		c.Memory[c.SP] = byte(c.PC)
		c.PC = addr
	case "JFC":
		if !c.Flags.Carry {
			c.PC = addr
		}
	case "JFZ":
		if !c.Flags.Zero {
			c.PC = addr
		}
	case "JFS":
		if !c.Flags.Sign {
			c.PC = addr
		}
	case "JFP":
		if !c.Flags.Parity {
			c.PC = addr
		}
	case "JTC":
		if c.Flags.Carry {
			c.PC = addr
		}
	case "JTZ":
		if c.Flags.Zero {
			c.PC = addr
		}
	case "JTS":
		if c.Flags.Sign {
			c.PC = addr
		}
	case "JTP":
		if c.Flags.Parity {
			c.PC = addr
		}
	case "CFC":
		if !c.Flags.Carry {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CFZ":
		if !c.Flags.Zero {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CFS":
		if !c.Flags.Sign {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CFP":
		if !c.Flags.Parity {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CTC":
		if c.Flags.Carry {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CTZ":
		if c.Flags.Zero {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CTS":
		if c.Flags.Sign {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	case "CTP":
		if c.Flags.Parity {
			c.SP--
			c.Memory[c.SP] = byte(c.PC >> 8)
			c.SP--
			c.Memory[c.SP] = byte(c.PC)
			c.PC = addr
		}
	default:
		fmt.Printf("🆘 Absolute instruction not implemented: %s\n", instruction.Mnemonic)
		os.Exit(1)
	}
}

// executeRelative executes a relative mode instruction
func (c *Intel8008) executeRelative(instruction Instruction, offset int8) {
	switch instruction.Mnemonic {
	default:
		fmt.Printf("🆘 Relative instruction not implemented: %s\n", instruction.Mnemonic)
		os.Exit(1)
	}
}

// executeImplied executes an implied mode instruction
func (c *Intel8008) executeImplied(instruction Instruction) {
	switch instruction.Mnemonic {
	// Register-to-register transfers
	case "LAB":
		c.A = c.B
	case "LAC":
		c.A = c.C
	case "LAD":
		c.A = c.D
	case "LAE":
		c.A = c.E
	case "LAH":
		c.A = c.H
	case "LAL":
		c.A = c.L
	case "LAM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.A = c.Memory[addr]

	case "LBA":
		c.B = c.A
	case "LBC":
		c.B = c.C
	case "LBD":
		c.B = c.D
	case "LBE":
		c.B = c.E
	case "LBH":
		c.B = c.H
	case "LBL":
		c.B = c.L
	case "LBM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.B = c.Memory[addr]

	case "LCA":
		c.C = c.A
	case "LCB":
		c.C = c.B
	case "LCD":
		c.C = c.D
	case "LCE":
		c.C = c.E
	case "LCH":
		c.C = c.H
	case "LCL":
		c.C = c.L
	case "LCM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.C = c.Memory[addr]

	case "LDA":
		c.D = c.A
	case "LDB":
		c.D = c.B
	case "LDC":
		c.D = c.C
	case "LDE":
		c.D = c.E
	case "LDH":
		c.D = c.H
	case "LDL":
		c.D = c.L
	case "LDM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.D = c.Memory[addr]

	case "LEA":
		c.E = c.A
	case "LEB":
		c.E = c.B
	case "LEC":
		c.E = c.C
	case "LED":
		c.E = c.D
	case "LEH":
		c.E = c.H
	case "LEL":
		c.E = c.L
	case "LEM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.E = c.Memory[addr]

	case "LHA":
		c.H = c.A
	case "LHB":
		c.H = c.B
	case "LHC":
		c.H = c.C
	case "LHD":
		c.H = c.D
	case "LHE":
		c.H = c.E
	case "LHL":
		c.H = c.L
	case "LHM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.H = c.Memory[addr]

	case "LLA":
		c.L = c.A
	case "LLB":
		c.L = c.B
	case "LLC":
		c.L = c.C
	case "LLD":
		c.L = c.D
	case "LLE":
		c.L = c.E
	case "LLH":
		c.L = c.H
	case "LLM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.L = c.Memory[addr]

	// Memory operations
	case "LMA":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.A
	case "LMB":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.B
	case "LMC":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.C
	case "LMD":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.D
	case "LME":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.E
	case "LMH":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.H
	case "LML":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.Memory[addr] = c.L

	// Increment/Decrement
	case "INB":
		c.B++
		c.updateFlags(c.B)
	case "INC":
		c.C++
		c.updateFlags(c.C)
	case "IND":
		c.D++
		c.updateFlags(c.D)
	case "INE":
		c.E++
		c.updateFlags(c.E)
	case "INH":
		c.H++
		c.updateFlags(c.H)
	case "INL":
		c.L++
		c.updateFlags(c.L)

	case "DCB":
		c.B--
		c.updateFlags(c.B)
	case "DCC":
		c.C--
		c.updateFlags(c.C)
	case "DCD":
		c.D--
		c.updateFlags(c.D)
	case "DCE":
		c.E--
		c.updateFlags(c.E)
	case "DCH":
		c.H--
		c.updateFlags(c.H)
	case "DCL":
		c.L--
		c.updateFlags(c.L)

	// ALU operations
	case "ADA":
		result := uint16(c.A) + uint16(c.A)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADB":
		result := uint16(c.A) + uint16(c.B)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADC":
		result := uint16(c.A) + uint16(c.C)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADD":
		result := uint16(c.A) + uint16(c.D)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADE":
		result := uint16(c.A) + uint16(c.E)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADH":
		result := uint16(c.A) + uint16(c.H)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADL":
		result := uint16(c.A) + uint16(c.L)
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ADM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		result := uint16(c.A) + uint16(c.Memory[addr])
		c.A = byte(result)
		c.updateFlagsWithCarry(result)

	case "ACA":
		result := uint16(c.A) + uint16(c.A)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACB":
		result := uint16(c.A) + uint16(c.B)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACC":
		result := uint16(c.A) + uint16(c.C)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACD":
		result := uint16(c.A) + uint16(c.D)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACE":
		result := uint16(c.A) + uint16(c.E)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACH":
		result := uint16(c.A) + uint16(c.H)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACL":
		result := uint16(c.A) + uint16(c.L)
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)
	case "ACM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		result := uint16(c.A) + uint16(c.Memory[addr])
		if c.Flags.Carry {
			result++
		}
		c.A = byte(result)
		c.updateFlagsWithCarry(result)

	case "SUA":
		result := uint16(c.A) - uint16(c.A)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUB":
		result := uint16(c.A) - uint16(c.B)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUC":
		result := uint16(c.A) - uint16(c.C)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUD":
		result := uint16(c.A) - uint16(c.D)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUE":
		result := uint16(c.A) - uint16(c.E)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUH":
		result := uint16(c.A) - uint16(c.H)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUL":
		result := uint16(c.A) - uint16(c.L)
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SUM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		result := uint16(c.A) - uint16(c.Memory[addr])
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)

	case "SBA":
		result := uint16(c.A) - uint16(c.A)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBB":
		result := uint16(c.A) - uint16(c.B)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBC":
		result := uint16(c.A) - uint16(c.C)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBD":
		result := uint16(c.A) - uint16(c.D)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBE":
		result := uint16(c.A) - uint16(c.E)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBH":
		result := uint16(c.A) - uint16(c.H)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBL":
		result := uint16(c.A) - uint16(c.L)
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)
	case "SBM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		result := uint16(c.A) - uint16(c.Memory[addr])
		if c.Flags.Carry {
			result--
		}
		c.A = byte(result)
		c.updateFlagsWithBorrow(result)

	case "NDA":
		c.A &= c.A
		c.updateFlags(c.A)
	case "NDB":
		c.A &= c.B
		c.updateFlags(c.A)
	case "NDC":
		c.A &= c.C
		c.updateFlags(c.A)
	case "NDD":
		c.A &= c.D
		c.updateFlags(c.A)
	case "NDE":
		c.A &= c.E
		c.updateFlags(c.A)
	case "NDH":
		c.A &= c.H
		c.updateFlags(c.A)
	case "NDL":
		c.A &= c.L
		c.updateFlags(c.A)
	case "NDM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.A &= c.Memory[addr]
		c.updateFlags(c.A)

	case "XRA":
		c.A ^= c.A
		c.updateFlags(c.A)
	case "XRB":
		c.A ^= c.B
		c.updateFlags(c.A)
	case "XRC":
		c.A ^= c.C
		c.updateFlags(c.A)
	case "XRD":
		c.A ^= c.D
		c.updateFlags(c.A)
	case "XRE":
		c.A ^= c.E
		c.updateFlags(c.A)
	case "XRH":
		c.A ^= c.H
		c.updateFlags(c.A)
	case "XRL":
		c.A ^= c.L
		c.updateFlags(c.A)
	case "XRM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.A ^= c.Memory[addr]
		c.updateFlags(c.A)

	case "ORA":
		c.A |= c.A
		c.updateFlags(c.A)
	case "ORB":
		c.A |= c.B
		c.updateFlags(c.A)
	case "ORC":
		c.A |= c.C
		c.updateFlags(c.A)
	case "ORD":
		c.A |= c.D
		c.updateFlags(c.A)
	case "ORE":
		c.A |= c.E
		c.updateFlags(c.A)
	case "ORH":
		c.A |= c.H
		c.updateFlags(c.A)
	case "ORL":
		c.A |= c.L
		c.updateFlags(c.A)
	case "ORM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.A |= c.Memory[addr]
		c.updateFlags(c.A)

	case "CPA":
		c.updateCompareFlags(c.A, c.A)
	case "CPB":
		c.updateCompareFlags(c.A, c.B)
	case "CPC":
		c.updateCompareFlags(c.A, c.C)
	case "CPD":
		c.updateCompareFlags(c.A, c.D)
	case "CPE":
		c.updateCompareFlags(c.A, c.E)
	case "CPH":
		c.updateCompareFlags(c.A, c.H)
	case "CPL":
		c.updateCompareFlags(c.A, c.L)
	case "CPM":
		addr := uint16(c.H)<<8 | uint16(c.L)
		c.updateCompareFlags(c.A, c.Memory[addr])

	// Rotate instructions
	case "RLC":
		carry := (c.A & 0x80) != 0
		c.A = (c.A << 1) | c.A>>7
		c.Flags.Carry = carry
	case "RRC":
		carry := (c.A & 0x01) != 0
		c.A = (c.A >> 1) | c.A<<7
		c.Flags.Carry = carry
	case "RAL":
		oldCarry := c.Flags.Carry
		c.Flags.Carry = (c.A & 0x80) != 0
		c.A = (c.A << 1)
		if oldCarry {
			c.A |= 0x01
		}
	case "RAR":
		oldCarry := c.Flags.Carry
		c.Flags.Carry = (c.A & 0x01) != 0
		c.A = (c.A >> 1)
		if oldCarry {
			c.A |= 0x80
		}

	// Return instruction
	case "RET":
		c.PC = uint16(c.Memory[c.SP]) | uint16(c.Memory[c.SP+1])<<8
		c.SP += 2

	case "HLT":
		os.Exit(0)

	default:
		fmt.Printf("🆘 Implied instruction not implemented: %s\n", instruction.Mnemonic)
		os.Exit(1)
	}
}

// Helper functions for flag updates
func (c *Intel8008) updateFlags(value byte) {
	c.Flags.Zero = value == 0
	c.Flags.Sign = (value & 0x80) != 0
	c.Flags.Parity = c.calculateParity(value)
}

func (c *Intel8008) updateFlagsWithCarry(result uint16) {
	c.Flags.Carry = result > 0xFF
	c.updateFlags(byte(result))
}

func (c *Intel8008) updateFlagsWithBorrow(result uint16) {
	c.Flags.Carry = result > 0xFF
	c.updateFlags(byte(result))
}

func (c *Intel8008) updateCompareFlags(a, value byte) {
	result := a - value
	c.Flags.Carry = a < value
	c.Flags.Zero = result == 0
	c.Flags.Sign = (result & 0x80) != 0
	c.Flags.Parity = c.calculateParity(result)
}

func (c *Intel8008) calculateParity(value byte) bool {
	parity := true
	for i := 0; i < 8; i++ {
		if (value & (1 << i)) != 0 {
			parity = !parity
		}
	}
	return parity
}

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
