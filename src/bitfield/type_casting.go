package bitfield

import "context"

type Bitfield interface {
	Set(ctx context.Context, index uint32, value bool) error
	Get(ctx context.Context, index uint32) (bool, error)
	Length() uint32
	AsBytes() []byte
}
