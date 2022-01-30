package tracker

import (
	"encoding/binary"
	"errors"
	"net"
)

var (
    ErrorPeerBytesInvalid = errors.New("data byte of peer should be 6 bytes")
)

type Peer interface {
    PeerID() string
    IP() net.IP
    Port() uint16
}

type peer_impl struct {
    peerID  string
    ip      net.IP
    port    uint16
}

func NewPeerFromByte(value []byte) (Peer, error) {
    if len(value) != 6 {
        return nil, ErrorPeerBytesInvalid
    }
    ip := net.IPv4(value[0], value[1], value[2], value[3])
    port := binary.BigEndian.Uint16(value[4:6])

    return peer_impl{
        ip: ip,
        port: port,
    }, nil
}

func (peer peer_impl) PeerID() string {
    return peer.peerID
}

func (peer peer_impl) IP() net.IP {
    return peer.ip
}

func (peer peer_impl) Port() uint16 {
    return peer.port
}
