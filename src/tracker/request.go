package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type Request struct {
    infohash    src.Hash
    peerID      src.PeerID
    ip          interface{}
    port        interface{}
    uploaded    uint64
    downloaded  uint64
    left        uint64
    event       interface{}
}

func NewRequest(infohash src.Hash, peerID src.PeerID, uploaded uint64, downloaded uint64, left uint64) src.TrackerRequest {
    return &Request{
        infohash: infohash,
        peerID: peerID,
        uploaded: uploaded,
        downloaded: downloaded,
        left: left,
    }
}

func (req *Request) InfoHash() src.Hash{
    return req.infohash
}

func (req *Request) PeerID() src.PeerID {
    return req.peerID
}

func (req *Request) SetIP(ip net.IP) {
    req.ip = ip
}
func (req *Request) IP() (net.IP, bool) {
    if ip, ok := req.ip.(net.IP); ok {
        return ip, true
    }

    return nil, false
}

func (req *Request) SetPort(port uint16) {
    req.port = port
}
func (req *Request) Port() (uint16, bool) {
    if port, ok := req.port.(uint16); ok {
        return port, true
    }

    return 0, false
}

func (req *Request) Uploaded() uint64 {
    return req.uploaded
}

func (req *Request) Downloaded() uint64 {
    return req.downloaded
}

func (req *Request) Left() uint64 {
    return req.left
}

func (req *Request) SetEvent(event src.Event) {
    req.event = event
}
func (req *Request) Event() (src.Event, bool) {
    if event, ok := req.event.(src.Event); ok {
        return event, true
    }

    return src.Event(""), false
}

