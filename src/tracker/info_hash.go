package tracker

import (
	"encoding/hex"
	"errors"
)

var (
    ErrorLengthHash = errors.New("length of hash should be 20 bytes")
    ErrorNotHexString = errors.New("value is not a hex string")
)

type InfoHash interface {
    String() string
    HexString() string
}

type infoHash_impl [20]byte

func NewInfoHashFromSileOfByte(value []byte) (InfoHash, error) {
    if len(value) != 20 {
        return nil, ErrorLengthHash
    }

    var infoHash infoHash_impl

    copy(infoHash[:],value)
    
    return infoHash, nil
}

func NewInfoHashFromHexString(value string) (InfoHash, error) {
    var err error

    hashByte, err := hex.DecodeString(value)
    if err != nil {
        return nil, ErrorNotHexString
    }

    return NewInfoHashFromSileOfByte(hashByte)
}

func (infoHash infoHash_impl) String() string {
    return string(infoHash[:])
}

func (infoHash infoHash_impl) HexString() string {
    return hex.EncodeToString(infoHash[:])
}
