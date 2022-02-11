package metainfo

import (
	"github.com/azamaulanaaa/gotor/src/bencode"
)

func Encode(metainfo Metainfo) (string, error) {
    rawData := Raw(metainfo)
    return bencode.Encode(rawData)
}
