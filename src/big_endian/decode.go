package bigendian

import (
	"bytes"
	"encoding/binary"
)

func Decode(b []byte, out interface{}) error {
    var err error
    
    err = binary.Read(bytes.NewReader(b), binary.BigEndian, out)

    return err
}
