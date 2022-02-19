package peer

import (
	"net"

	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
)

func Decode(b []byte) (Peer, error) {
	var err error

	var peer Peer

	switch len(b) {
	case 6:
		peer.IP = net.IP(b[:4])

		err = bigendian.Decode(b[4:], &peer.Port)
		if err != nil {
			return Peer{}, err
		}
	case 18:
		peer.IP = net.IP(b[:16])

		err = bigendian.Decode(b[16:], &peer.Port)
		if err != nil {
			return Peer{}, err
		}
	default:
		return Peer{}, ErrorPeerBytesInvalid
	}

	return peer, nil
}
