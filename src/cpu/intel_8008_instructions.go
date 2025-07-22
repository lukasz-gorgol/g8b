package cpu

var Intel8008Instructions = map[byte]Instruction{

	// Index Register Instructions
	// The load instructions do not affect the flag flip-flops.
	// The increment and decrement instructions affect all flip-flops except the carry.
	// Memory registers are addressed by contents of registers H and L.

	0xC1: {0xC1, "LAB", Implied, 1, 5, "Load register A with content of register B"},
	0xC2: {0xC2, "LAC", Implied, 1, 5, "Load register A with content of register C"},
	0xC3: {0xC3, "LAD", Implied, 1, 5, "Load register A with content of register D"},
	0xC4: {0xC4, "LAE", Implied, 1, 5, "Load register A with content of register E"},
	0xC5: {0xC5, "LAH", Implied, 1, 5, "Load register A with content of register H"},
	0xC6: {0xC6, "LAL", Implied, 1, 5, "Load register A with content of register L"},

	0xC8: {0xC8, "LBA", Implied, 1, 5, "Load register B with content of register A"},
	0xCA: {0xCA, "LBC", Implied, 1, 5, "Load register B with content of register C"},
	0xCB: {0xCB, "LBD", Implied, 1, 5, "Load register B with content of register D"},
	0xCC: {0xCC, "LBE", Implied, 1, 5, "Load register B with content of register E"},
	0xCD: {0xCD, "LBH", Implied, 1, 5, "Load register B with content of register H"},
	0xCE: {0xCE, "LBL", Implied, 1, 5, "Load register B with content of register L"},

	0xD0: {0xD0, "LCA", Implied, 1, 5, "Load register C with content of register A"},
	0xD1: {0xD1, "LCB", Implied, 1, 5, "Load register C with content of register B"},
	0xD3: {0xD3, "LCD", Implied, 1, 5, "Load register C with content of register D"},
	0xD4: {0xD4, "LCE", Implied, 1, 5, "Load register C with content of register E"},
	0xD5: {0xD5, "LCH", Implied, 1, 5, "Load register C with content of register H"},
	0xD6: {0xD6, "LCL", Implied, 1, 5, "Load register C with content of register L"},

	0xD8: {0xD8, "LDA", Implied, 1, 5, "Load register D with content of register A"},
	0xD9: {0xD9, "LDB", Implied, 1, 5, "Load register D with content of register B"},
	0xDA: {0xDA, "LDC", Implied, 1, 5, "Load register D with content of register C"},
	0xDC: {0xDC, "LDE", Implied, 1, 5, "Load register D with content of register E"},
	0xDD: {0xDD, "LDH", Implied, 1, 5, "Load register D with content of register H"},
	0xDE: {0xDE, "LDL", Implied, 1, 5, "Load register D with content of register L"},

	0xE0: {0xE0, "LEA", Implied, 1, 5, "Load register E with content of register A"},
	0xE1: {0xE1, "LEB", Implied, 1, 5, "Load register E with content of register B"},
	0xE2: {0xE2, "LEC", Implied, 1, 5, "Load register E with content of register C"},
	0xE3: {0xE3, "LED", Implied, 1, 5, "Load register E with content of register D"},
	0xE5: {0xE5, "LEH", Implied, 1, 5, "Load register E with content of register H"},
	0xE6: {0xE6, "LEL", Implied, 1, 5, "Load register E with content of register L"},

	0xE8: {0xE8, "LHA", Implied, 1, 5, "Load register H with content of register A"},
	0xE9: {0xE9, "LHB", Implied, 1, 5, "Load register H with content of register B"},
	0xEA: {0xEA, "LHC", Implied, 1, 5, "Load register H with content of register C"},
	0xEB: {0xEB, "LHD", Implied, 1, 5, "Load register H with content of register D"},
	0xEC: {0xEC, "LHE", Implied, 1, 5, "Load register H with content of register E"},
	0xEE: {0xEE, "LHL", Implied, 1, 5, "Load register H with content of register L"},

	0xF0: {0xF0, "LLA", Implied, 1, 5, "Load register L with content of register A"},
	0xF1: {0xF1, "LLB", Implied, 1, 5, "Load register L with content of register B"},
	0xF2: {0xF2, "LLC", Implied, 1, 5, "Load register L with content of register C"},
	0xF3: {0xF3, "LLD", Implied, 1, 5, "Load register L with content of register D"},
	0xF4: {0xF4, "LLE", Implied, 1, 5, "Load register L with content of register E"},
	0xF5: {0xF5, "LLH", Implied, 1, 5, "Load register L with content of register H"},

	0xC7: {0xC7, "LAM", Implied, 1, 8, "Load register A with content of memory register M"},
	0xCF: {0xCF, "LBM", Implied, 1, 8, "Load register B with content of memory register M"},
	0xD7: {0xD7, "LCM", Implied, 1, 8, "Load register C with content of memory register M"},
	0xDF: {0xDF, "LDM", Implied, 1, 8, "Load register D with content of memory register M"},
	0xE7: {0xE7, "LEM", Implied, 1, 8, "Load register E with content of memory register M"},
	0xEF: {0xEF, "LHM", Implied, 1, 8, "Load register H with content of memory register M"},
	0xF7: {0xF7, "LLM", Implied, 1, 8, "Load register L with content of memory register M"},

	0xF8: {0xF8, "LMA", Implied, 1, 7, "Load memory register M with content of register A"},
	0xF9: {0xF9, "LMB", Implied, 1, 7, "Load memory register M with content of register B"},
	0xFA: {0xFA, "LMC", Implied, 1, 7, "Load memory register M with content of register C"},
	0xFB: {0xFB, "LMD", Implied, 1, 7, "Load memory register M with content of register D"},
	0xFC: {0xFC, "LME", Implied, 1, 7, "Load memory register M with content of register E"},
	0xFD: {0xFD, "LMH", Implied, 1, 7, "Load memory register M with content of register H"},
	0xFE: {0xFE, "LML", Implied, 1, 7, "Load memory register M with content of register L"},

	0x06: {0x06, "LAI", Immediate, 2, 8, "Load register A with data"},
	0x0E: {0x0E, "LBI", Immediate, 2, 8, "Load register B with data"},
	0x16: {0x16, "LCI", Immediate, 2, 8, "Load register C with data"},
	0x1E: {0x1E, "LDI", Immediate, 2, 8, "Load register D with data"},
	0x26: {0x26, "LEI", Immediate, 2, 8, "Load register E with data"},
	0x2E: {0x2E, "LHI", Immediate, 2, 8, "Load register H with data"},
	0x36: {0x36, "LLI", Immediate, 2, 8, "Load register L with data"},

	0x3E: {0x3E, "LMI", Immediate, 2, 9, "Load memory register M with data"},

	0x08: {0x08, "INB", Implied, 1, 5, "Increment content of register B"},
	0x10: {0x10, "INC", Implied, 1, 5, "Increment content of register C"},
	0x18: {0x18, "IND", Implied, 1, 5, "Increment content of register D"},
	0x20: {0x20, "INE", Implied, 1, 5, "Increment content of register E"},
	0x28: {0x28, "INH", Implied, 1, 5, "Increment content of register H"},
	0x30: {0x30, "INL", Implied, 1, 5, "Increment content of register L"},

	0x09: {0x09, "DCB", Implied, 1, 5, "Decrement content of register B"},
	0x11: {0x11, "DCC", Implied, 1, 5, "Decrement content of register C"},
	0x19: {0x19, "DCD", Implied, 1, 5, "Decrement content of register D"},
	0x21: {0x21, "DCE", Implied, 1, 5, "Decrement content of register E"},
	0x29: {0x29, "DCH", Implied, 1, 5, "Decrement content of register H"},
	0x31: {0x31, "DCL", Implied, 1, 5, "Decrement content of register L"},

	// Accumulator Group Instructions
	// The result of the ALU instructions affect all flag flip-flops.
	// Rotate instructions affect only the carry flip-flop.

	0x80: {0x80, "ADA", Implied, 1, 5, "Add content of register A to accumulator. Overflow sets carry flag"},
	0x81: {0x81, "ADB", Implied, 1, 5, "Add content of register B to accumulator. Overflow sets carry flag"},
	0x82: {0x82, "ADC", Implied, 1, 5, "Add content of register C to accumulator. Overflow sets carry flag"},
	0x83: {0x83, "ADD", Implied, 1, 5, "Add content of register D to accumulator. Overflow sets carry flag"},
	0x84: {0x84, "ADE", Implied, 1, 5, "Add content of register E to accumulator. Overflow sets carry flag"},
	0x85: {0x85, "ADH", Implied, 1, 5, "Add content of register H to accumulator. Overflow sets carry flag"},
	0x86: {0x86, "ADL", Implied, 1, 5, "Add content of register L to accumulator. Overflow sets carry flag"},

	0x87: {0x87, "ADM", Implied, 1, 8, "Add content of memory register M to accumulator. Overflow sets carry flag"},

	0x04: {0x04, "ADI", Immediate, 2, 8, "Add content of data to accumulator. Overflow sets carry flag"},

	0x88: {0x88, "ACA", Implied, 1, 5, "Add content of register A from accumulator with carry. Overflow sets carry flag"},
	0x89: {0x89, "ACB", Implied, 1, 5, "Add content of register B from accumulator with carry. Overflow sets carry flag"},
	0x8A: {0x8A, "ACC", Implied, 1, 5, "Add content of register C from accumulator with carry. Overflow sets carry flag"},
	0x8B: {0x8B, "ACD", Implied, 1, 5, "Add content of register D from accumulator with carry. Overflow sets carry flag"},
	0x8C: {0x8C, "ACE", Implied, 1, 5, "Add content of register E from accumulator with carry. Overflow sets carry flag"},
	0x8D: {0x8D, "ACH", Implied, 1, 5, "Add content of register H from accumulator with carry. Overflow sets carry flag"},
	0x8E: {0x8E, "ACL", Implied, 1, 5, "Add content of register L from accumulator with carry. Overflow sets carry flag"},

	0x8F: {0x8F, "ACM", Implied, 1, 8, "Add content of memory register M from accumulator with carry. Overflow sets carry flag"},

	0x0C: {0x0C, "ACI", Immediate, 2, 8, "Add content of data from accumulator with carry. Overflow sets carry flag"},

	0x90: {0x90, "SUA", Implied, 1, 5, "Subtract content of register A from accumulator. Underflow sets carry flag"},
	0x91: {0x91, "SUB", Implied, 1, 5, "Subtract content of register B from accumulator. Underflow sets carry flag"},
	0x92: {0x92, "SUC", Implied, 1, 5, "Subtract content of register C from accumulator. Underflow sets carry flag"},
	0x93: {0x93, "SUD", Implied, 1, 5, "Subtract content of register D from accumulator. Underflow sets carry flag"},
	0x94: {0x94, "SUE", Implied, 1, 5, "Subtract content of register E from accumulator. Underflow sets carry flag"},
	0x95: {0x95, "SUH", Implied, 1, 5, "Subtract content of register H from accumulator. Underflow sets carry flag"},
	0x96: {0x96, "SUL", Implied, 1, 5, "Subtract content of register L from accumulator. Underflow sets carry flag"},

	0x97: {0x97, "SUM", Implied, 1, 8, "Subtract content of memory register M from accumulator. Underflow sets carry flag"},

	0x14: {0x14, "SUI", Immediate, 2, 8, "Subtract content of data from accumulator. Underflow sets carry flag"},

	0x98: {0x98, "SBA", Implied, 1, 5, "Subtract content of register A from accumulator with borrow. Underflow sets carry flag"},
	0x99: {0x99, "SBB", Implied, 1, 5, "Subtract content of register B from accumulator with borrow. Underflow sets carry flag"},
	0x9A: {0x9A, "SBC", Implied, 1, 5, "Subtract content of register C from accumulator with borrow. Underflow sets carry flag"},
	0x9B: {0x9B, "SBD", Implied, 1, 5, "Subtract content of register D from accumulator with borrow. Underflow sets carry flag"},
	0x9C: {0x9C, "SBE", Implied, 1, 5, "Subtract content of register E from accumulator with borrow. Underflow sets carry flag"},
	0x9D: {0x9D, "SBH", Implied, 1, 5, "Subtract content of register H from accumulator with borrow. Underflow sets carry flag"},
	0x9E: {0x9E, "SBL", Implied, 1, 5, "Subtract content of register L from accumulator with borrow. Underflow sets carry flag"},

	0x9F: {0x9F, "SBM", Implied, 1, 8, "Subtract content of memory register M from accumulator with borrow. Underflow sets carry flag"},

	0x1C: {0x1C, "SBI", Immediate, 2, 8, "Subtract content of data from accumulator with borrow. Underflow sets carry flag"},

	0xA0: {0xA0, "NDA", Implied, 1, 5, "Compute logical AND of register A content with accumulator"},
	0xA1: {0xA1, "NDB", Implied, 1, 5, "Compute logical AND of register B content with accumulator"},
	0xA2: {0xA2, "NDC", Implied, 1, 5, "Compute logical AND of register C content with accumulator"},
	0xA3: {0xA3, "NDD", Implied, 1, 5, "Compute logical AND of register D content with accumulator"},
	0xA4: {0xA4, "NDE", Implied, 1, 5, "Compute logical AND of register E content with accumulator"},
	0xA5: {0xA5, "NDH", Implied, 1, 5, "Compute logical AND of register H content with accumulator"},
	0xA6: {0xA6, "NDL", Implied, 1, 5, "Compute logical AND of register L content with accumulator"},

	0xA7: {0xA7, "NDM", Implied, 1, 8, "Compute logical AND of memory register M content with accumulator"},

	0x24: {0x24, "NDI", Immediate, 2, 8, "Compute logical AND of data content with accumulator"},

	0xA8: {0xA8, "XRA", Implied, 1, 5, "Compute EXCLUSIVE OR of register A content with accumulator"},
	0xA9: {0xA9, "XRB", Implied, 1, 5, "Compute EXCLUSIVE OR of register B content with accumulator"},
	0xAA: {0xAA, "XRC", Implied, 1, 5, "Compute EXCLUSIVE OR of register C content with accumulator"},
	0xAB: {0xAB, "XRD", Implied, 1, 5, "Compute EXCLUSIVE OR of register D content with accumulator"},
	0xAC: {0xAC, "XRE", Implied, 1, 5, "Compute EXCLUSIVE OR of register E content with accumulator"},
	0xAD: {0xAD, "XRH", Implied, 1, 5, "Compute EXCLUSIVE OR of register H content with accumulator"},
	0xAE: {0xAE, "XRL", Implied, 1, 5, "Compute EXCLUSIVE OR of register L content with accumulator"},

	0xAF: {0xAF, "XRM", Implied, 1, 8, "Compute EXCLUSIVE OR of memory register M content with accumulator"},

	0x2C: {0x2C, "XRI", Immediate, 2, 8, "Compute EXCLUSIVE OR of data content with accumulator"},

	0xB0: {0xB0, "ORA", Implied, 1, 5, "Compute INCLUSIVE OR of register A content with accumulator"},
	0xB1: {0xB1, "ORB", Implied, 1, 5, "Compute INCLUSIVE OR of register B content with accumulator"},
	0xB2: {0xB2, "ORC", Implied, 1, 5, "Compute INCLUSIVE OR of register C content with accumulator"},
	0xB3: {0xB3, "ORD", Implied, 1, 5, "Compute INCLUSIVE OR of register D content with accumulator"},
	0xB4: {0xB4, "ORE", Implied, 1, 5, "Compute INCLUSIVE OR of register E content with accumulator"},
	0xB5: {0xB5, "ORH", Implied, 1, 5, "Compute INCLUSIVE OR of register H content with accumulator"},
	0xB6: {0xB6, "ORL", Implied, 1, 5, "Compute INCLUSIVE OR of register L content with accumulator"},

	0xB7: {0xB7, "ORM", Implied, 1, 8, "Compute INCLUSIVE OR of memory register M content with accumulator"},

	0x34: {0x34, "ORI", Immediate, 2, 8, "Compute INCLUSIVE OR of data content with accumulator"},

	0xB8: {0xB8, "CPA", Implied, 1, 5, "Compare content of register A content with accumulator. Accumulator unchanged"},
	0xB9: {0xB9, "CPB", Implied, 1, 5, "Compare content of register B content with accumulator. Accumulator unchanged"},
	0xBA: {0xBA, "CPC", Implied, 1, 5, "Compare content of register C content with accumulator. Accumulator unchanged"},
	0xBB: {0xBB, "CPD", Implied, 1, 5, "Compare content of register D content with accumulator. Accumulator unchanged"},
	0xBC: {0xBC, "CPE", Implied, 1, 5, "Compare content of register E content with accumulator. Accumulator unchanged"},
	0xBD: {0xBD, "CPH", Implied, 1, 5, "Compare content of register H content with accumulator. Accumulator unchanged"},
	0xBE: {0xBE, "CPL", Implied, 1, 5, "Compare content of register L content with accumulator. Accumulator unchanged"},

	0xBF: {0xBF, "CPM", Implied, 1, 8, "Compare content of memory register M content with accumulator. Accumulator unchanged"},

	0x3C: {0x3C, "CPI", Immediate, 2, 8, "Compare content of data content with accumulator. Accumulator unchanged"},

	0x02: {0x02, "RLC", Implied, 1, 5, "Rotate content of accumulator left"},
	0x0A: {0x0A, "RRC", Implied, 1, 5, "Rotate content of accumulator right"},
	0x12: {0x12, "RAL", Implied, 1, 5, "Rotate content of accumulator left through carry"},
	0x1A: {0x1A, "RAR", Implied, 1, 5, "Rotate content of accumulator right through carry"},

	// Program Counter and Stack Control Instructions

	0x44: {0x44, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x4C: {0x4C, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x54: {0x54, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x5C: {0x5C, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x64: {0x64, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x6C: {0x6C, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x74: {0x74, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},
	0x7C: {0x7C, "JMP", Absolute, 3, 11, "Unconditional jump to memory address B3..B3[6] B2..B2[8]"},

	0x40: {0x40, "JFC", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Carry is false"},
	0x48: {0x48, "JFZ", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Zero is false"},
	0x50: {0x50, "JFS", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Sign is false"},
	0x58: {0x58, "JFP", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Parity is false"},

	0x60: {0x60, "JTC", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Carry is true"},
	0x68: {0x68, "JTZ", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Zero is true"},
	0x70: {0x70, "JTS", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Sign is true"},
	0x78: {0x78, "JTP", Absolute, 3, 11, "Jump to memory address B3..B3[6] B2..B2[8] if flag Parity is true"},

	0x46: {0x46, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x4E: {0x4E, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x56: {0x56, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x5E: {0x5E, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x66: {0x66, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x6E: {0x6E, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x76: {0x76, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},
	0x7E: {0x7E, "CAL", Absolute, 3, 11, "Unconditionaly call memory address B3..B3[6] B2..B2[8]. Save current address in stack"},

	0x42: {0x42, "CFC", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Carry is false"},
	0x4A: {0x4A, "CFZ", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Zero is false"},
	0x52: {0x52, "CFS", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Sign is false"},
	0x5A: {0x5A, "CFP", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Parity is false"},

	0x62: {0x62, "CTC", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Carry is true"},
	0x6A: {0x6A, "CTZ", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Zero is true"},
	0x72: {0x72, "CTS", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Sign is true"},
	0x7A: {0x7A, "CTP", Absolute, 3, 11, "Call memory address B3..B3[6] B2..B2[8] and save current address in stack if flag Parity is true"},

	0x07: {0x07, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x0F: {0x0F, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x17: {0x17, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x1F: {0x1F, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x27: {0x27, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x2F: {0x2F, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x37: {0x37, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},
	0x3F: {0x3F, "RET", Implied, 1, 5, "Unconditionaly return. Down one level in stack"},

	0x03: {0x03, "RFC", Implied, 1, 5, "Return one level is stack if flag Carry is false"},
	0x0B: {0x0B, "RFZ", Implied, 1, 5, "Return one level in stack if flag Zero is false"},
	0x13: {0x13, "RFS", Implied, 1, 5, "Return one level in stack if flag Sign is false"},
	0x1B: {0x1B, "RFP", Implied, 1, 5, "Return one level in stack if flag Parity is false"},

	0x23: {0x23, "RTC", Implied, 1, 5, "Return one level in stack if flag Carry is true"},
	0x2B: {0x2B, "RTZ", Implied, 1, 5, "Return one level in stack if flag Zero is true"},
	0x33: {0x33, "RTS", Implied, 1, 5, "Return one level in stack if flag Sign is true"},
	0x3B: {0x3B, "RTP", Implied, 1, 5, "Return one level in stack if flag Parity is true"},

	0x05: {0x05, "RST", Implied, 1, 5, "Call subroutine at memory address 000000. Up one level in stack"},
	0x0D: {0x0D, "RST", Implied, 1, 5, "Call subroutine at memory address 001000. Up one level in stack"},
	0x15: {0x15, "RST", Implied, 1, 5, "Call subroutine at memory address 010000. Up one level in stack"},
	0x1D: {0x1D, "RST", Implied, 1, 5, "Call subroutine at memory address 011000. Up one level in stack"},
	0x25: {0x25, "RST", Implied, 1, 5, "Call subroutine at memory address 100000. Up one level in stack"},
	0x2D: {0x2D, "RST", Implied, 1, 5, "Call subroutine at memory address 101000. Up one level in stack"},
	0x35: {0x35, "RST", Implied, 1, 5, "Call subroutine at memory address 110000. Up one level in stack"},
	0x3D: {0x3D, "RST", Implied, 1, 5, "Call subroutine at memory address 111000. Up one level in stack"},

	// Input / Output Instructions

	0x41: {0x41, "INP", Implied, 1, 8, "Read content of input port 000 into accumulator"},
	0x43: {0x43, "INP", Implied, 1, 8, "Read content of input port 001 into accumulator"},
	0x45: {0x45, "INP", Implied, 1, 8, "Read content of input port 010 into accumulator"},
	0x47: {0x47, "INP", Implied, 1, 8, "Read content of input port 011 into accumulator"},
	0x49: {0x49, "INP", Implied, 1, 8, "Read content of input port 100 into accumulator"},
	0x4B: {0x4B, "INP", Implied, 1, 8, "Read content of input port 101 into accumulator"},
	0x4D: {0x4D, "INP", Implied, 1, 8, "Read content of input port 110 into accumulator"},
	0x4F: {0x4F, "INP", Implied, 1, 8, "Read content of input port 111 into accumulator"},

	0x51: {0x51, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01000"},
	0x53: {0x53, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01001"},
	0x55: {0x55, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01010"},
	0x57: {0x57, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01011"},
	0x59: {0x59, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01100"},
	0x5B: {0x5B, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01101"},
	0x5D: {0x5D, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01110"},
	0x5F: {0x5F, "OUT", Implied, 1, 6, "Write content of accumulator into output port 01111"},

	0x61: {0x61, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10000"},
	0x63: {0x63, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10001"},
	0x65: {0x65, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10010"},
	0x67: {0x67, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10011"},
	0x69: {0x69, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10100"},
	0x6B: {0x6B, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10101"},
	0x6D: {0x6D, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10110"},
	0x6F: {0x6F, "OUT", Implied, 1, 6, "Write content of accumulator into output port 10111"},

	0x71: {0x71, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11000"},
	0x73: {0x73, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11001"},
	0x75: {0x75, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11010"},
	0x77: {0x77, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11011"},
	0x79: {0x79, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11100"},
	0x7B: {0x7B, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11101"},
	0x7D: {0x7D, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11110"},
	0x7F: {0x7F, "OUT", Implied, 1, 6, "Write content of accumulator into output port 11111"},

	// NOP, No Operation Instructions

	0xC0: {0xC0, "NOP", Implied, 1, 5, "No operation"},
	0xC9: {0xC9, "NOP", Implied, 1, 5, "No operation"},
	0xD2: {0xD2, "NOP", Implied, 1, 5, "No operation"},
	0xDB: {0xDB, "NOP", Implied, 1, 5, "No operation"},
	0xE4: {0xE4, "NOP", Implied, 1, 5, "No operation"},
	0xED: {0xED, "NOP", Implied, 1, 5, "No operation"},
	0xF6: {0xF6, "NOP", Implied, 1, 5, "No operation"},

	// Machine Instructions

	0x00: {0x00, "HLT", Implied, 1, 4, "Enter STOPPED state; remain there until interrupted"},
	0x01: {0x01, "HLT", Implied, 1, 4, "Enter STOPPED state; remain there until interrupted"},
	0xFF: {0xFF, "HLT", Implied, 1, 4, "Enter STOPPED state; remain there until interrupted"},
}
