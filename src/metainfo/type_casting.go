package metainfo

import (
	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/hash"
)

type Metainfo interface {
	Announce() string
	Info() Info
	Raw() bencode.Dictionary
	InfoHash() (hash.Hash, error)
}
type Info interface {
	PieceLength() uint32
	Pieces() []hash.Hash
	Length() (uint64, bool)
	Files() ([]File, bool)
	Name() (string, bool)
	Private() (bool, bool)
}
type File interface {
	Length() uint64
	Path() string
}
