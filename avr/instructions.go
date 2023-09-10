package avr

// TODO generate this file automatically
type Op uint16

const (
	_ Op = iota
	ADC
	ADD
	ADIW
	AND
	ANDI
	ASR
	BCLR
	BLD
	BRBC
	BRBS
	BRCC
	BRCS
	BREAK
	BREQ
	BRGE
	BRHC
	BRHS
	BRID
	BRIE
	BRLO
	BRLT
	BRMI
	BRNE
	BRPL
	BRSH
	BRTC
	BRTS
	BRVC
	BRVS
	BSET
	BST
	CALL
	CBI
	CBR
	CLC
	CLH
	CLI
	CLN
	CLR
	CLS
	CLT
	CLV
	CLZ
	COM
	CP
	CPC
	CPI
	CPSE
	DEC
	DES
	EICALL
	EIJMP
	ELPM
	EOR
	FMUL
	FMULS
	FMULSU
	ICALL
	IJMP
	IN
	INC
	JMP
	LAC
	LAS
	LAT
	LDI
	LDS
	LPM
	LSL
	LSR
	MOV
	MOVW
	MUL
	MULS
	MULSU
	NEG
	NOP
	OR
	ORI
	OUT
	POP
	PUSH
	RCALL
	RET
	RETI
	RJMP
	ROL
	ROR
	SBC
	SBCI
	SBI
	SBIC
	SBIS
	SBIW
	SBR
	SBRC
	SEI
	SEZ
	SLEEP
	LD
	SPM
	ST
	STS
	SUB
	SUBI
	SWAP
	TST
	WDR
	XCH
)

type Inst struct {
	//mask     uint32
	//value    uint32
	Op       Op
	Size     uint8
	Mnemonic string
	Operands []uint32
}

type OperandType = uint

const (
	UnknownOperandType OperandType = iota
	ConstantAddress
)

type Operand struct {
	Mask uint32
	Type OperandType
}

var instrFormats = []struct {
	Op       Op
	Mnemonic string
	mask     uint16
	value    uint16
	Size     uint8
	Operands []Operand
}{
	// ADC
	{Op: ADC, Mnemonic: "ADC", mask: 0xFC00, value: 0x1C00, Size: 2},
	// ADD – Add without Carry
	{Op: ADD, Mnemonic: "ADD", mask: 0xFC00, value: 0xC00, Size: 2},
	// ADIW – Add Immediate to Word
	{Op: ADIW, Mnemonic: "ADIW", mask: 0xFF00, value: 0x9600, Size: 2},
	// AND
	{Op: AND, Mnemonic: "AND", mask: 0xFC00, value: 0x2000, Size: 2},
	// ANDI – Logical AND with Immediate
	{Op: ANDI, Mnemonic: "ANDI", mask: 0xF000, value: 0x7000, Size: 2},
	// BREQ – Branch if Equal
	{Op: BREQ, Mnemonic: "BREQ", mask: 0xFC07, value: 0xF001, Size: 2},
	// BRNE
	{Op: BRNE, Mnemonic: "BRNE", mask: 0xFC07, value: 0xF401, Size: 2},
	// BRCC – Branch if Carry Cleared
	{Op: BRCC, Mnemonic: "BRCC", mask: 0xFC07, value: 0xF400, Size: 2},
	// BRCS – Branch if Carry Set
	{Op: BRCS, Mnemonic: "BRCS", mask: 0xFC07, value: 0xF000, Size: 2},
	// CALL
	{Op: CALL, Mnemonic: "CALL", mask: 0xFE0E, value: 0x940E, Size: 4},
	// CLI – Clear Global Interrupt Enable Bit
	{Op: CLI, Mnemonic: "CLI", mask: 0xFFFF, value: 0x94F8, Size: 2},
	// CLR
	{Op: CLR, Mnemonic: "CLR", mask: 0xFC00, value: 0x2400, Size: 2}, // as EOR
	// CP – Compare
	{Op: CP, Mnemonic: "CP", mask: 0xFC00, value: 0x1400, Size: 2},
	// CPSE – Compare Skip if Equal
	{Op: CPSE, Mnemonic: "CPSE", mask: 0xFC00, value: 0x1000, Size: 2},
	// COM – One’s Complement
	{Op: COM, Mnemonic: "COM", mask: 0xFE0F, value: 0x9400, Size: 2},
	// CPC
	{Op: CPC, Mnemonic: "CPC", mask: 0xFC00, value: 0x0400, Size: 2},
	// CPI – Compare with Immediate
	{Op: CPI, Mnemonic: "CPI", mask: 0xF000, value: 0x3000, Size: 2},
	// DEC – Decrement
	{Op: DEC, Mnemonic: "DEC", mask: 0xFE0F, value: 0x940A, Size: 2},
	// EOR – Exclusive OR
	{Op: EOR, Mnemonic: "EOR", mask: 0xFC00, value: 0x2400, Size: 2}, // as CLR
	// IN - Load an I/O Location to Register
	{Op: IN, Mnemonic: "IN", mask: 0xF800, value: 0xB000, Size: 2},
	// LD – Load Indirect from Data Space to Register using X
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x900C, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x900D, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x900E, Size: 2},
	// LD (LDD) – Load Indirect from Data Space to Register using Y
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x8008, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x9009, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x900A, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xD208, value: 0x8008, Size: 2},
	// LD (LDD) – Load Indirect From Data Space to Register using Z
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x8000, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x9001, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xFE0F, value: 0x9002, Size: 2},
	{Op: LD, Mnemonic: "LD", mask: 0xD208, value: 0x8000, Size: 2},
	// LDI
	{Op: LDI, Mnemonic: "LDI", mask: 0xF000, value: 0xE000, Size: 2},
	// LDS – Load Direct from Data Space
	{Op: LDS, Mnemonic: "LDS", mask: 0xFE0F, value: 0x9000, Size: 4},
	// LDS (AVRrc) – Load Direct from Data Space
	{Op: LDS, Mnemonic: "LDS", mask: 0xF800, value: 0xA000, Size: 2},
	// LPM – Load Program Memory
	{Op: LPM, Mnemonic: "LPM", mask: 0xFFFF, value: 0x95C8, Size: 2},
	{Op: LPM, Mnemonic: "LPM", mask: 0xFE0F, value: 0x9004, Size: 2},
	{Op: LPM, Mnemonic: "LPM", mask: 0xFE0F, value: 0x9005, Size: 2},
	// MOV – Copy Register
	{Op: MOV, Mnemonic: "MOV", mask: 0xFC00, value: 0x2C00, Size: 2},
	// MOVW – Copy Register Word
	{Op: MOVW, Mnemonic: "MOVW", mask: 0xFF00, value: 0x100, Size: 2},
	// MULS
	{Op: MULS, Mnemonic: "MULS", mask: 0xFF00, value: 0x200, Size: 2},
	// MULSU
	{Op: MULSU, Mnemonic: "MULSU", mask: 0xFF88, value: 0x300, Size: 2},
	// NOP
	{Op: NOP, Mnemonic: "NOP", mask: 0xFFFF, value: 0x0000, Size: 2},
	// OR – Logical OR
	{Op: OR, Mnemonic: "OR", mask: 0xFC00, value: 0x2800, Size: 2},
	// ORI
	{Op: ORI, Mnemonic: "ORI", mask: 0xF000, value: 0x6000, Size: 2},
	// OUT
	{Op: OUT, Mnemonic: "OUT", mask: 0xF800, value: 0xB800, Size: 2},
	// POP
	{Op: POP, Mnemonic: "POP", mask: 0xFE0F, value: 0x900F, Size: 2},
	// PUSH
	{Op: PUSH, Mnemonic: "PUSH", mask: 0xFE0F, value: 0x920F, Size: 2},
	// RET
	{Op: RET, Mnemonic: "RET", mask: 0xFFFF, value: 0x9508, Size: 2},
	// RETI – Return from Interrupt
	{Op: RETI, Mnemonic: "RETI", mask: 0xFFFF, value: 0x9518, Size: 2},
	// RJMP
	{Op: RJMP, Mnemonic: "RJMP", mask: 0xF000, value: 0xC000, Size: 2},
	// SBC
	{Op: SBC, Mnemonic: "SBC", mask: 0xFC00, value: 0x800, Size: 2},
	// SBI
	{Op: SBI, Mnemonic: "SBI", mask: 0xFF00, value: 0x9A00, Size: 2},
	// SBIS – Skip if Bit in I/O Register is Set
	{Op: SBIS, Mnemonic: "SBIS", mask: 0xFF00, value: 0x9B00, Size: 2},
	// SBIW
	{Op: SBIW, Mnemonic: "SBIW", mask: 0xFF00, value: 0x9700, Size: 2},
	// SBCI – Subtract Immediate with Carry SBI – Set Bit in I/O Register
	{Op: SBCI, Mnemonic: "SBCI", mask: 0xF000, value: 0x4000, Size: 2},
	// SEI – Set Global Interrupt Enable Bit
	{Op: SEI, Mnemonic: "SEI", mask: 0xFFFF, value: 0x9478, Size: 2},
	// ST – Store Indirect From Register to Data Space using Index X
	{Op: ST, Mnemonic: "ST", mask: 0xFE0F, value: 0x920C, Size: 2},
	{Op: ST, Mnemonic: "ST", mask: 0xFE0F, value: 0x920D, Size: 2},
	{Op: ST, Mnemonic: "ST", mask: 0xFE0F, value: 0x920E, Size: 2},
	// STS
	{Op: STS, Mnemonic: "STS", mask: 0xFE00, value: 0x9200, Size: 4},
	// SUB – Subtract Without Carry
	{Op: SUB, Mnemonic: "SUB", mask: 0xFC00, value: 0x1800, Size: 2},
	// SUBI – Subtract Immediate
	{Op: SUBI, Mnemonic: "SUBI", mask: 0xF000, value: 0x5000, Size: 2},
	// JMP
	{Op: JMP, Mnemonic: "JMP", mask: 0xFE0E, value: 0x940C, Size: 4, Operands: []Operand{{Mask: 0x1F1FFFF, Type: ConstantAddress}}},
}
