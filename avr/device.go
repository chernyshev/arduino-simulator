package avr

import (
	"debug/macho"
	"fmt"
	"os"
	"strings"
)

const (
	flashSize  = 0x4000
	sramSize   = 0x08FF
	eepromSize = 0x400
)

type CPU struct {
	// ALU – Arithmetic Logic Unit
	SREG byte     // Status Register
	GPRF [32]byte //General Purpose Register File
}

// RAM (from Random-Access Memory)
// ROM (from Read-Only Memory)

// Flash memory in microcontroller-based systems is part of its ROM. The flash memory is where the system's firmware is stored to be executed. For example, think of the famous Blink.ino sketch: when we compile this sketch, we create a binary file that is later stored in the flash memory of an Arduino board. The sketch is then executed when the board is powered on.
// RAM in microcontroller-based systems is where the system's temporary data or run-time data is stored; for example, the variables created by functions of a program. RAM in microcontrollers usually is SRAM; this is a type of RAM that uses a flip-flop to store one bit of data. There is also another type of RAM that can be found in microcontrollers: DRAM.
// EEPROM In microcontroller-based systems, Erasable Programmable Read-Only Memory, or EEPROM, is also part of its ROM; actually, Flash memory is a type of EEPROM. The main difference between Flash memory and EEPROM is how they are managed; EEPROM can be managed at the byte level (write or erased) while Flash can be managed at the block level.
type Memory struct {
	Flash  [flashSize]byte
	SRAM   [sramSize]byte
	EEPROM [eepromSize]byte
}

type Device struct {
	Memory
	macho.Cpu
}

func New() *Device {
	return &Device{}
}

func (d *Device) LoadBinary(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	//if len(data) != 0x8000 {
	//	return errors.New("file size error")
	//}

	for i := range data {
		d.Flash[i] = data[i]
	}

	return nil
}

func (d Device) Run() error {
	offset := 0
	for i := 0; i < len(d.Flash); {
		diff := 4
		if len(d.Flash)-i < 4 {
			diff = len(d.Flash) - i
		}
		instr, err := Decode(d.Flash[i : i+diff])
		if err != nil {
			fmt.Printf(err.Error())
			//	//return err
			i += 2
			offset = i
			continue
		}
		opcode := d.Flash[i : i+int(instr.Size)]
		i += int(instr.Size)
		fmt.Printf("%4x:\t% -12x\t%s\n", offset, opcode, strings.ToLower(instr.Mnemonic))
		offset = i
	}
	return nil
}
