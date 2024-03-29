package peer

import (
	"errors"
	"net"
)

type PeerID [20]byte

type Peer struct {
	PeerID PeerID
	IP     net.IP
	Port   uint16
}

var (
	ErrorPeerBytesInvalid = errors.New("data byte of peer should be either 6 or 18 bytes")
)
