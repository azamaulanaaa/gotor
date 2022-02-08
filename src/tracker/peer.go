package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type peer struct {
    peerid interface{}
    ip interface{}
    port interface{}
}

func NewPeer() src.Peer {
    thePeer := peer{}
    return &thePeer
}

func (self *peer) SetPeerID(peerid src.PeerID){
    self.peerid = peerid
}
func (self *peer) GetPeerID() (src.PeerID, bool) {
    if peerid, ok := self.peerid.(src.PeerID); ok {
        return peerid, true
    }

    return src.PeerID{}, false
}

func (self *peer) SetIP(ip net.IP) {
    self.ip = ip
}
func (self *peer) GetIP() (net.IP, bool) {
    if ip, ok := self.ip.(net.IP); ok {
        return ip, true
    }

    return nil, false
}

func (self *peer) SetPort(port uint16) {
    self.port = port
}
func (self *peer) GetPort() (uint16, bool) {
    if port, ok := self.port.(uint16); ok {
        return port, true
    }

    return 0, false
}

