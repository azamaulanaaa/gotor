package peer

import (
	"net"

	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
)

type peer_impl struct {
    peerID  interface{}
    ip      net.IP
    port    uint16
}

func NewPeerFromBytes(b []byte) (Peer, error) {
    var err error

    if len(b) != 6 {
        return nil, ErrorPeerBytesInvalid
    }

    ip := net.IPv4(b[0], b[1], b[2], b[3])

    var port uint16
    err = bigendian.Decode(b[4:], &port)
    if err != nil {
        return nil, err
    }

    peer := peer_impl{
        ip: ip,
        port: port,
    }

    return &peer, nil
}

func (peer *peer_impl) PeerID() (PeerID, bool) {
    if peerID, ok := peer.peerID.(PeerID); ok {
        return peerID, true
    }

    return PeerID{}, false
}

func (self *peer_impl) IP() net.IP {
    return self.ip
}

func (self *peer_impl) Port() uint16 {
    return self.port
}

