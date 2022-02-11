package metainfo

import (
	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/hash"
)

func InfoHash(metainfo Metainfo) (hash.Hash, error) {
    rawData := Raw(metainfo).(bencode.Dictionary)
    infoBencode, err := bencode.Encode(rawData["info"])
    if err != nil {
        return hash.Hash{}, err
    }

    hash := hash.Calculate([]byte(infoBencode))

    return hash, nil
}
