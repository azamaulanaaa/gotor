package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

type request_impl struct {
    infohash    hash.Hash
    peerID      peer.PeerID
    ip          interface{}
    port        interface{}
    uploaded    uint64
    downloaded  uint64
    left        uint64
    event       interface{}
}

func NewRequest(infohash hash.Hash, peerID peer.PeerID, uploaded uint64, downloaded uint64, left uint64) Request {
    return &request_impl{
        infohash: infohash,
        peerID: peerID,
        uploaded: uploaded,
        downloaded: downloaded,
        left: left,
    }
}

func (req *request_impl) InfoHash() hash.Hash{
    return req.infohash
}

func (req *request_impl) PeerID() peer.PeerID {
    return req.peerID
}

func (req *request_impl) SetIP(ip net.IP) {
    req.ip = ip
}
func (req *request_impl) IP() (net.IP, bool) {
    if ip, ok := req.ip.(net.IP); ok {
        return ip, true
    }

    return nil, false
}

func (req *request_impl) SetPort(port uint16) {
    req.port = port
}
func (req *request_impl) Port() (uint16, bool) {
    if port, ok := req.port.(uint16); ok {
        return port, true
    }

    return 0, false
}

func (req *request_impl) Uploaded() uint64 {
    return req.uploaded
}

func (req *request_impl) Downloaded() uint64 {
    return req.downloaded
}

func (req *request_impl) Left() uint64 {
    return req.left
}

func (req *request_impl) SetEvent(event Event) {
    req.event = event
}
func (req *request_impl) Event() (Event, bool) {
    if event, ok := req.event.(Event); ok {
        return event, true
    }

    return Event(""), false
}

