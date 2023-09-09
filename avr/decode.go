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
	//_ = binary.BigEndian.Uint16(src[2:4])

	for i := range instrFormats {
		f := &instrFormats[i]
		if instrOpcode&f.mask != f.value {
			continue
		}

		inst = Inst{
			Op:       f.Op,
			Size:     f.Size,
			Mnemonic: f.Mnemonic,
		}
		return inst, nil
	}

	return Inst{
		Size:     2,
		Mnemonic: ".word",
	}, nil
}
