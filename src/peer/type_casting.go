package peer

import (
    "errors"
    "net"
)

type PeerID     [20]byte

type Peer interface {
    PeerID()    (PeerID, bool)
    IP()        net.IP
    Port()      uint16
}

var (
    ErrorPeerBytesInvalid = errors.New("data byte of peer should be 6 bytes")
)
