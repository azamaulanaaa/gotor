package bitfield

import (
	"errors"

	"github.com/azamaulanaaa/gotor/src"
)

var (
    ErrorOutOfIndex = errors.New("index must be lower than the length")
)

type bitfield []byte

func NewBitfield(length uint64) src.Bitfield {
    return make(bitfield, 0, length / 8)
}

func BitFieldFormBytes(b []byte) src.Bitfield {
    return bitfield(b)
}

func (self bitfield) Set(index uint64, value bool) error {
    if index >= self.Length() {
        return ErrorOutOfIndex
    }

    sectionIndex := index / 8
    bitIndex := index % 8

    var newBit byte
    newBit = 1 << (7 - bitIndex)

    if value {
        self[sectionIndex] = self[sectionIndex] | newBit
    } else {
        self[sectionIndex] = self[sectionIndex] &^ newBit
    }

    return nil
}

func (self bitfield) Get(index uint64) (bool, error) {
    if index >= self.Length() {
        return false, ErrorOutOfIndex
    }

    sectionIndex := index / 8
    bitIndex := index % 8

    bit := self[sectionIndex] >> (7 - bitIndex) & 1
    return bit == 1, nil
}

func (self bitfield) Length() uint64 {
    return uint64(len(self) * 8)
}
