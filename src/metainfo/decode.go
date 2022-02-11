package metainfo

import (
	"errors"
	"io"

	"github.com/azamaulanaaa/gotor/src/bencode"
)

var (
    ErrorInvalidMetainfo = errors.New("invalid metainfo data")
)

func Decode(r io.Reader) (Metainfo, error) {
    rawMetainfo, err := bencode.Decode(r)
    if err != nil {
        return nil, err
    }

    if metainfo, ok := rawMetainfo.(bencode.Dictionary); ok {
        return metainfo_impl(metainfo), nil
    }

    return nil, ErrorInvalidMetainfo
}
