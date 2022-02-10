package tracker

import "github.com/azamaulanaaa/gotor/src"

type response struct {
    interval    uint16
    peers       []src.Peer
}

func (res *response) Interval() uint16 {
    return res.interval
}

func (res *response) Peers() []src.Peer {
    return res.peers
}
