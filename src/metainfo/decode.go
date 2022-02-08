package metainfo

import (
	"errors"
	"io"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

var (
    ErrorInvalidMetainfo = errors.New("invalid metainfo data")
)

func Decode(r io.Reader) (src.Metainfo, error) {
    rawMetainfo, err := bencode.Decode(r)
    if err != nil {
        return nil, err
    }

    if theMetainfo, ok := rawMetainfo.(bencode.Dictionary); ok {
        return metainfo(theMetainfo), nil
    }

    return nil, ErrorInvalidMetainfo
}
