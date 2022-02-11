package tracker

import (
	"time"

	"github.com/azamaulanaaa/gotor/src/peer"
)

type response struct {
    interval    time.Duration
    peers       []peer.Peer
}

func (res *response) Interval() time.Duration {
    return res.interval
}

func (res *response) Peers() []peer.Peer {
    return res.peers
}
