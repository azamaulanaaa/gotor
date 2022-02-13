package bitfield

import (
	"context"
	"errors"
)

var (
	ErrorOutOfIndex = errors.New("index must be lower than the length")
)

type bitfield_impl []byte

func NewBitfield(length uint32) Bitfield {
	sectionLength := length / 8
	if length%8 > 0 {
		sectionLength++
	}

	return make(bitfield_impl, sectionLength)
}

func BitFieldFormBytes(b []byte) Bitfield {
	return bitfield_impl(b)
}

func (bitfield bitfield_impl) Set(ctx context.Context, index uint32, value bool) error {
	if index >= bitfield.Length() {
		return ErrorOutOfIndex
	}

	sectionIndex := index / 8
	bitIndex := index % 8

	var newBit byte
	newBit = 1 << (7 - bitIndex)

	if value {
		bitfield[sectionIndex] = bitfield[sectionIndex] | newBit
	} else {
		bitfield[sectionIndex] = bitfield[sectionIndex] &^ newBit
	}

	return nil
}

func (bitfield bitfield_impl) Get(ctx context.Context, index uint32) (bool, error) {
	if index >= bitfield.Length() {
		return false, ErrorOutOfIndex
	}

	sectionIndex := index / 8
	bitIndex := index % 8

	bit := bitfield[sectionIndex] >> (7 - bitIndex) & 1
	return bit == 1, nil
}

func (bitfield bitfield_impl) Length() uint32 {
	return uint32(len(bitfield) * 8)
}

func (bitfield bitfield_impl) AsBytes() []byte {
	return []byte(bitfield)
}
