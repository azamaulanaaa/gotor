package bigendian

import (
	"bytes"
	"encoding/binary"
)

func Encode(value interface{}) ([]byte, error) {
    var err error

    var buff bytes.Buffer
    err = binary.Write(&buff, binary.BigEndian, value)
    if err != nil {
        return nil, ErrorTypeNotSupported
    }

    return buff.Bytes(), nil
}

