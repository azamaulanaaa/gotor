package tracker

import (
	"net"

	"github.com/azamaulanaaa/gotor/src"
)

type request struct {
    infohash interface{}
    peerid interface{}
    ip interface{}
    port interface{}
    uploaded interface{}
    downloaded interface{}
    left interface{}
    event interface{}
}

func NewRequest() src.TrackerRequest {
    req := request{}

    return &req
}

func (req *request) SetInfoHash(hash src.Hash) {
    req.infohash = hash
}
func (req *request) GetInfoHash() (src.Hash, bool){
    if hash, ok := req.infohash.(src.Hash); ok {
        return hash, true
    }

    return src.Hash{}, false
}

func (req *request) SetPeerID(peerID src.PeerID) {
    req.peerid = peerID
}
func (req *request) GetPeerID() (src.PeerID, bool) {
    if peerid, ok := req.peerid.(src.PeerID); ok {
        return peerid, true
    }

    return src.PeerID{}, false
}

func (req *request) SetIP(ip net.IP) {
    req.ip = ip
}
func (req *request) GetIP() (net.IP, bool) {
    if ip, ok := req.ip.(net.IP); ok {
        return ip, true
    }

    return nil, false
}

func (req *request) SetPort(port uint16) {
    req.port = port
}
func (req *request) GetPort() (uint16, bool) {
    if port, ok := req.port.(uint16); ok {
        return port, true
    }

    return 0, false
}

func (req *request) SetUploaded(uploaded uint64) {
    req.uploaded = uploaded
}
func (req *request) GetUploaded() (uint64, bool) {
    if uploaded, ok := req.uploaded.(uint64); ok {
        return uploaded, true
    }

    return 0, false
}

func (req *request) SetDownloaded(downloaded uint64) {
    req.downloaded = downloaded
}
func (req *request) GetDownloaded() (uint64, bool) {
    if downloaded, ok := req.downloaded.(uint64); ok {
        return downloaded, true
    }

    return 0, false
}

func (req *request) SetLeft(left uint64) {
    req.left = left
}
func (req *request) GetLeft() (uint64, bool) {
    if left, ok := req.left.(uint64); ok {
        return left, true
    }
    
    return 0, false
}

func (req *request) SetEvent(event src.Event) {
    req.event = event
}
func (req *request) GetEvent() (src.Event, bool) {
    if event, ok := req.event.(src.Event); ok {
        return event, true
    }

    return src.Event(""), false
}

