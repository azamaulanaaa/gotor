package metainfo

import (
	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

func Encode(metainfo src.Metainfo) (string, error) {
    rawData := Raw(metainfo)
    return bencode.Encode(rawData)
}
