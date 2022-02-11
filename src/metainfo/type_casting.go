package metainfo

import (
    "github.com/azamaulanaaa/gotor/src/hash"
)

type Metainfo interface {
    Announce()      string
    Info()          Info
}
type Info interface {
    PieceLength()   uint64
    Pieces()        []hash.Hash
    Length()        (uint64, bool)
    Files()         ([]File, bool)
    Name()          (string, bool)
    Private()       (bool, bool)
}
type File interface {
    Length() uint64
    Path() string
}

