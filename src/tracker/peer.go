package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type peer struct {
    peerid  interface{}
    ip      net.IP
    port    uint16
}

func NewPeer() src.Peer {
    return &peer{}
}

func (self *peer) PeerID() (src.PeerID, bool) {
    if peerid, ok := self.peerid.(src.PeerID); ok {
        return peerid, true
    }

    return src.PeerID{}, false
}

func (self *peer) IP() net.IP {
    return self.ip
}

func (self *peer) Port() uint16 {
    return self.port
}

