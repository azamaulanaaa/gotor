package hash

import (
	"crypto/sha1"
	"encoding/hex"
	"errors"
)

var (
    ErrorLengthHash = errors.New("length of hash should be 20 bytes")
    ErrorNotHexString = errors.New("value is not a hex string")
)

func Calculate(b []byte) Hash {
    var hash Hash

    hasher := sha1.New()
    hasher.Write(b)
    rawHash := hasher.Sum(nil)
   
    copy(hash[:], rawHash)

    return hash
}

func Decode(value string) (Hash, error) {
    var err error

    hashByte, err := hex.DecodeString(value)
    if err != nil {
        return Hash{}, ErrorNotHexString
    }

    if len(hashByte) != 20 {
        return Hash{}, ErrorLengthHash
    }

    var infoHash Hash
    copy(infoHash[:], hashByte)

    return infoHash, nil
}
