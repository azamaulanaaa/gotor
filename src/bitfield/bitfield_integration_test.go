//go:build integration

package bitfield_test

import (
	"context"
	"testing"

	"github.com/azamaulanaaa/gotor/src/bitfield"
	"github.com/azamaulanaaa/gotor/test"
)

type testData struct {
	bytes         []byte
	actual_length uint32
	length        uint32
	index         uint32
	value         bool
}

func TestBitfield(t *testing.T) {
	theTestDatas := []testData{
		{[]byte{0}, 3, 8, 4, false},
		{[]byte{0, 0}, 16, 16, 10, false},
	}

	ctx := context.Background()

	for _, theTestData := range theTestDatas {
		theBitfield := bitfield.NewBitfield(theTestData.actual_length)

		test.Equals(t, theTestData.bytes, theBitfield.AsBytes())
		test.Equals(t, theTestData.length, theBitfield.Length())

		{
			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, theTestData.value, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.index, true)
			test.Ok(t, err)

			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, true, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.index, false)
			test.Ok(t, err)

			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, false, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.length, true)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)

			_, err = theBitfield.Get(ctx, theTestData.length)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)
		}

		{
			err := theBitfield.Set(ctx, theTestData.length+10, true)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)

			_, err = theBitfield.Get(ctx, theTestData.length+10)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)
		}
	}
}

func TestBitfieldFromByte(t *testing.T) {
	theTestDatas := []testData{
		{[]byte{105}, 7, 8, 4, true},
		{[]byte{1, 181, 167}, 20, 24, 5, false},
	}

	ctx := context.Background()

	for _, theTestData := range theTestDatas {
		theBitfield := bitfield.BitFieldFormBytes(theTestData.bytes)

		test.Equals(t, theTestData.bytes, theBitfield.AsBytes())
		test.Equals(t, theTestData.length, theBitfield.Length())

		{
			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, theTestData.value, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.index, true)
			test.Ok(t, err)

			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, true, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.index, false)
			test.Ok(t, err)

			value, err := theBitfield.Get(ctx, theTestData.index)
			test.Ok(t, err)
			test.Equals(t, false, value)
		}

		{
			err := theBitfield.Set(ctx, theTestData.length, true)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)

			_, err = theBitfield.Get(ctx, theTestData.length)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)
		}

		{
			err := theBitfield.Set(ctx, theTestData.length+10, true)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)

			_, err = theBitfield.Get(ctx, theTestData.length+10)
			test.Equals(t, bitfield.ErrorOutOfIndex, err)
		}
	}
}
