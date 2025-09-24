# Galoping Top

<img src="assets/galoping-top-logo.jpeg" alt="Galoping Top" width="200" style="background: transparent; border: none; border-radius: 10px;">

A simple 8-bit processor emulator and assembler written in Go by Lukasz Gorgol.

**Project website:** https://galoping.top

**License:** This project is licensed under a custom license that allows free use for personal, educational, and research purposes while restricting commercial use. See [LICENSE](LICENSE.md) file for details.

**Author:** Lukasz Gorgol (Poland)

## Features

- Multiple 8-bit CPU implementations (starting with 8008)
- Basic assembler with CPU selection support
- Support for common addressing modes
- Memory inspection capabilities
- Configurable memory size and program start address
- JSON configuration support for both assembler and emulator
- Interactive debugger with step-by-step execution
- Unified instruction structure with opcodes, cycles, and addressing modes
- Configurable CPU speed with cycle-accurate timing
- Optimized build system for maximum performance
- **Verbose mode** for detailed register/flag output (see below)

## Legal Notice

This software is released under the custom license, which provides strong protection against patent trolls and ensures the software remains free and open source. The author, Lukasz Gorgol, reserves the exclusive right to use this code under any other license terms. No other party may use this code for commercial purposes without explicit written permission from the author.

Any legal disputes arising from this software shall be resolved exclusively in the courts of Poland, with Polish law being the governing law.

## Directory Structure

All program source, binary, and config files are located in the `program/` directory:

```
program/
  intel_8008.asm   # 8008 assembly source
  intel_8008.bin   # 8008 assembled binary
  intel_8008.json  # 8008 config (for both assembler and emulator)
```

## Instruction Structure

The instruction set is implemented using a unified structure that includes:
- Opcode (1 byte)
- Mnemonic (instruction name)
- Addressing mode
- Instruction size in bytes
- Number of cycles
- Description

Example instruction definition:
```go
0xA9: {0xA9, "LDA", Immediate, 2, 2, "Load Accumulator"}
```

## Addressing Modes

The 8008 processor supports several addressing modes that determine how operands are accessed. Here are the addressing modes implemented in our emulator:

### 1. Immediate Addressing Mode
- Symbol: `#$`
- Description: The operand is a constant value that follows the instruction
- Example: `LDA #$1A` (Load the value 0x1A into accumulator)
- Implementation: The value is stored in the byte following the opcode
- Used by: LDA, LDX, ADC, AND, ORA, EOR, CMP
- Cycles: 2

### 2. Absolute Addressing Mode
- Symbol: `$`
- Description: The operand is a 16-bit address that follows the instruction
- Example: `STA $0200` (Store accumulator at memory location 0x0200)
- Implementation: Two bytes follow the opcode (low byte, high byte)
- Used by: STA, STX, JMP, JSR
- Cycles: 4 (STA, STX), 3 (JMP), 6 (JSR)

### 3. Relative Addressing Mode
- Description: Used for branch instructions, specifies a signed offset from the next instruction
- Example: `BCC is_less` (Branch if Carry Clear to label 'is_less')
- Implementation: One byte follows the opcode (signed offset)
- Used by: BCC, BNE, BEQ
- Cycles: 2 (not taken), 3 (taken)

### 4. Implied Addressing Mode
- Description: The instruction operates on an implied register or location
- Example: `DEX` (Decrement X register)
- Implementation: No additional bytes follow the opcode
- Used by: DEX, PHA, PLA, RTS, HLT
- Cycles: 2 (DEX), 3 (PHA), 4 (PLA), 6 (RTS), 1 (HLT)

## Instruction Format

Each instruction consists of:
1. Opcode (1 byte)
2. Operand (0-2 bytes, depending on addressing mode)

### Example Instruction Encoding:
```
LDA #$1A    -> A9 1A    (2 cycles)
STA $0200   -> 8D 00 02 (4 cycles)
BCC is_less -> 90 03    (2/3 cycles)
DEX         -> CA       (2 cycles)
```

## Memory Organization

- Program memory starts at 0x8000
- Stack is located at 0x0100-0x01FF
- Zero page is at 0x0000-0x00FF
- I/O and system memory is at 0x0200-0xFFFF

## Processor Registers

- A: Accumulator (8-bit)
- X: X Index Register (8-bit)
- Y: Y Index Register (8-bit)
- PC: Program Counter (16-bit)
- SP: Stack Pointer (8-bit)
- P: Processor Status Register (8-bit)
  - C: Carry Flag
  - Z: Zero Flag
  - N: Negative Flag

## Example Program

```assembly
; A clean program to test the subroutine logic.
    LDA #$1A      ; Load 26 into A
    JSR check_gte_10 ; Jump to subroutine

    ; Store the results
    STA $0200     ; Store original value at $0200
    STX $0201     ; Store result at $0201
    HLT           ; End of program

check_gte_10:
    CMP #$0A      ; Compare A with 10
    BCC is_less   ; Branch if A < 10
    LDX #$01      ; A >= 10, set X to 1
    RTS           ; Return

is_less:
    LAI #$00      ; A < 10, set A to 0
    RET           ; Return
```

## Build and Run

### Using the Build System

The project includes an optimized build system that produces highly optimized binaries.

```bash
# Make the build script executable (if not already)
chmod +x build.sh

# Build both emulator and assembler (default)
./build.sh

# Build only the emulator
./build.sh emulator

# Build only the assembler
./build.sh assembler

# Clean build artifacts
./build.sh clean

# Build optimized release binaries for multiple platforms
./build.sh release

# Run benchmarks
./build.sh bench

# Run tests
./build.sh test

# Build with profiling enabled
./build.sh profile

# Show help
./build.sh help
```

The optimized binaries will be placed in the `bin/` directory. Release builds for multiple platforms will be placed in the `dist/` directory.

### Manual Build

```bash
# Build the assembler
go build -o asm src/assembler/assembler.go

# Build the emulator
go build -o emu src/emulator/emulator.go

# Assemble a program (default 8008 CPU)
./asm program/intel_8008.asm program/intel_8008.bin

# Assemble a program with specific CPU
./asm -c 8008 program/intel_8008.asm program/intel_8008.bin

# Assemble using JSON config (recommended)
./asm -c program/intel_8008.json

# Run the emulator with default settings (64KB memory, start at 0x8000)
./bin/emulator program/intel_8008.bin

# Run with JSON configuration (recommended)
./bin/emulator -c program/intel_8008.json
```

## Command-line Options

### Assembler Options
- `-c <file>`: Path to JSON configuration file
- `-cpu <type>`: CPU type (default: 8008)

### Emulator Options
- `-c <file>`: Path to JSON configuration file (if provided, no other options should be used)
- `-s <addr>`: Start address for program loading and PC initialization (hex string, e.g., `0x8000`)
- `-d <addrs>`: Memory addresses to dump after execution
  - Single address: `0x0200`
  - Range: `0x0200-0x0205`
  - List: `0x0200,0x0201,0x0202`
  - Mixed: `0x0200,0x0202-0x0205,0x0207`
- `-m <size>`: Memory size in bytes (default: 65536, max: 65536)
- `-cpu <type>`: CPU type (default: 8008)
- `-speed <hz>`: CPU speed in Hz (default: 1000000 for 1MHz)
- `-debug`: Run in debug mode
- `-v`: **Verbose mode** (show PC, registers, and flags for each instruction; otherwise, only shown in debug mode)

## JSON Configuration

The assembler and emulator can both be configured using a single JSON file. The configuration file supports the following fields:

```json
{
    "source": "program/intel_8008.asm", // Assembler: path to source file
    "binary": "program/intel_8008.bin", // Assembler: output binary; Emulator: input binary
    "cpu": "8008",                      // CPU type (default: "8008")
    "start_addr": "0x8000",             // Emulator: start address as hex string (default: "0x8000")
    "memory_size": 65536,               // Emulator: memory size in bytes (default: 65536)
    "dump_addrs": "0x0200-0x0201",      // Emulator: memory addresses to dump
    "verbose": true                     // Emulator: enable verbose output
}
```

Example configuration file (`program/intel_8008.json`):
```json
{
    "source": "program/intel_8008.asm",
    "binary": "program/intel_8008.bin",
    "cpu": "8008",
    "start_addr": "0x8000",
    "memory_size": 65536,
    "dump_addrs": "0x0200-0x0201",
    "verbose": true
}
```

- The assembler uses `source`, `binary`, and `cpu` fields.
- The emulator uses `binary`, `cpu`, `start_addr`, `memory_size`, `dump_addrs`, and `verbose` fields.
- You can use the same config file for both tools.

## Memory Address Specification

The emulator supports flexible memory address specifications for inspecting memory contents after program execution:

- Single address: `0x0200`
- Range of addresses: `0x0200-0x0205`
- Comma-separated list: `0x0200,0x0201,0x0202`
- Mixed format: `0x0200,0x0202-0x0205,0x0207`

If no memory addresses are specified, the emulator will run the program without displaying any memory contents.

## References

- [Addressing Modes - GeeksforGeeks](https://www.geeksforgeeks.org/addressing-modes/)
- [8008 Instruction Set](http://dunfield.classiccmp.org/mod8/8008um.pdf)

## Debugger

The emulator includes an interactive debugger that can be started with the `-debug` flag:

```bash
./emu -debug program/intel_8008.bin
```

### Debugger Commands

- `s` or `step`: Execute one instruction
- `c` or `continue`: Continue execution until HLT
- `r` or `registers`: Show current register values
- `m <addr>`: Show memory at address
- `q` or `quit`: Exit debugger
- `h` or `help`: Show help

## References

- [Intel 8008 User Manual](http://dunfield.classiccmp.org/mod8/8008um.pdf)
- [Intel 8008 Instruction Set Reference](https://en.wikipedia.org/wiki/Intel_8008)

## License

This project is licensed under a custom license that allows free use for personal, educational, and research purposes while restricting commercial use. The license includes strong patent protection provisions to defend against patent trolls and ensures all disputes are handled under Polish jurisdiction. Commercial entities interested in using this software must obtain a separate commercial license. The license maintains full attribution requirements and reserves the right for future relicensing under different terms. For commercial licensing inquiries or questions about usage rights, please contact lgorgol@gmail.com. See the [LICENSE](LICENSE.md) file for complete terms and conditions.
