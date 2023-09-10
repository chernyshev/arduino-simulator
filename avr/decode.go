package avr

import (
	"encoding/binary"
	"errors"
)

func Decode(src []byte) (inst Inst, err error) {
	if len(src) < 4 {
		return Inst{}, errors.New("too short")
	}

	instrOpcode := binary.LittleEndian.Uint16(src[0:2])
	dataOpcode := binary.LittleEndian.Uint16(src[2:4])

	for i := range instrFormats {
		f := &instrFormats[i]
		if instrOpcode&f.mask != f.value {
			continue
		}

		var opcode uint32
		switch f.Size {
		case 2:
			opcode = uint32(instrOpcode)
		case 4:
			opcode = (uint32(instrOpcode) << 16) | uint32(dataOpcode)
		default:
			return Inst{}, errors.New("invalid opcode size")
		}

		var operands []uint32
		for j := range f.Operands {
			operands = append(operands, extractOperand(opcode, f.Operands[j].Mask))
		}

		inst = Inst{
			Op:       f.Op,
			Size:     f.Size,
			Mnemonic: f.Mnemonic,
			Operands: operands,
		}
		return inst, nil
	}

	return Inst{
		Size:     2,
		Mnemonic: ".word",
	}, nil
}

func extractOperand(data uint32, mask uint32) uint32 {
	significantBytes := data & mask

	var operand uint32

	for m := mask; m != 0; {
		// check highest bit
		if (m&0x80000000)>>31 == 1 {
			operand = (operand << 1) | ((significantBytes & 0x80000000) >> 31)
		}
		m = m << 1
		significantBytes = significantBytes << 1
	}

	return operand
}
