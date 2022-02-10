package tracker

import (
	"time"

	"github.com/azamaulanaaa/gotor/src"
)

type response struct {
    interval    time.Duration
    peers       []src.Peer
}

func (res *response) Interval() time.Duration {
    return res.interval
}

func (res *response) Peers() []src.Peer {
    return res.peers
}
