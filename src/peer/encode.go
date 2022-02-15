package peer

import (
	"bytes"

	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
)

func Encode(peer Peer) ([]byte, error) {
	buff := &bytes.Buffer{}
	{
		data := peer.IP.To4()
		if data == nil {
			data = peer.IP.To16()
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(peer.Port)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff.Bytes(), nil
}
