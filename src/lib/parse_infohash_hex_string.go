package lib

import (
	"encoding/hex"
    "errors"

    "github.com/azamaulanaaa/gotor/src"
)

var (
    ErrorLengthHash = errors.New("length of hash should be 20 bytes")
    ErrorNotHexString = errors.New("value is not a hex string")
)

func ParseInfoHashHexString(value string) (src.InfoHash, error) {
    var err error

    hashByte, err := hex.DecodeString(value)
    if err != nil {
        return src.InfoHash{}, ErrorNotHexString
    }

    if len(hashByte) != 20 {
        return src.InfoHash{}, ErrorLengthHash
    }

    var infoHash src.InfoHash
    copy(infoHash[:], hashByte)

    return infoHash, nil
}
