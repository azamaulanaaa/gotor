package tracker

import (
	"net"
	"time"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

type Tracker interface {
    String()            string
    Do(Request)  (Response, error)
}

type Request interface {
    InfoHash()      hash.Hash
    PeerID()        peer.PeerID
    IP()            (net.IP, bool)
    Port()          (uint16, bool)
    Uploaded()      uint64
    Downloaded()    uint64
    Left()          uint64
    Event()         (Event, bool)
}

type Response interface {
    Interval()      time.Duration
    Peers()         []peer.Peer
}

type Event      string
