package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type request struct {
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
    return &request{
        infohash: infohash,
        peerID: peerID,
        uploaded: uploaded,
        downloaded: downloaded,
        left: left,
    }
}

func (req *request) InfoHash() src.Hash{
    return req.infohash
}

func (req *request) PeerID() src.PeerID {
    return req.peerID
}

func (req *request) SetIP(ip net.IP) {
    req.ip = ip
}
func (req *request) IP() (net.IP, bool) {
    if ip, ok := req.ip.(net.IP); ok {
        return ip, true
    }

    return nil, false
}

func (req *request) SetPort(port uint16) {
    req.port = port
}
func (req *request) Port() (uint16, bool) {
    if port, ok := req.port.(uint16); ok {
        return port, true
    }

    return 0, false
}

func (req *request) Uploaded() uint64 {
    return req.uploaded
}

func (req *request) Downloaded() uint64 {
    return req.downloaded
}

func (req *request) Left() uint64 {
    return req.left
}

func (req *request) SetEvent(event src.Event) {
    req.event = event
}
func (req *request) Event() (src.Event, bool) {
    if event, ok := req.event.(src.Event); ok {
        return event, true
    }

    return src.Event(""), false
}

