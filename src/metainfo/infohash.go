package metainfo

import (
	"crypto/sha1"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

func InfoHash(metainfo src.Metainfo) (src.Hash, error) {
    rawData := Raw(metainfo).(bencode.Dictionary)
    infoBencode, err := bencode.Encode(rawData["info"])
    if err != nil {
        return src.Hash{}, err
    }

    hasher := sha1.New()
    hasher.Write([]byte(infoBencode))
    hashSlice := hasher.Sum(nil)
    
    var hash src.Hash
    copy(hash[:], hashSlice)

    return hash, nil
}
