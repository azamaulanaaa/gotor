package metainfo

import (
	"github.com/azamaulanaaa/gotor/src/bencode"
)

func Encode(metainfo Metainfo) (string, error) {
	rawData := metainfo.Raw()
	return bencode.Encode(rawData)
}
