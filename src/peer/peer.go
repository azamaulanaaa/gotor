package peer

import (
	"net"

	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
)

func NewPeer(peerID PeerID, ip net.IP, port uint16) Peer {
    return Peer{
        PeerID: peerID,
        IP: ip,
        Port: port,
    }
}

func NewPeerFromBytes(b []byte) (Peer, error) {
    var err error

    if len(b) != 6 {
        return Peer{}, ErrorPeerBytesInvalid
    }

    ip := net.IPv4(b[0], b[1], b[2], b[3])

    var port uint16
    err = bigendian.Decode(b[4:], &port)
    if err != nil {
        return Peer{}, err
    }

    peer := Peer{
        IP: ip,
        Port: port,
    }

    return peer, nil
}
