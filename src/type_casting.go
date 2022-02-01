package src

import (
    "net"
)

type PeerID [20]byte
type Peer struct {
    PeerID  PeerID
    IP      net.IP
    Port    uint16
}
type InfoHash [20]byte
type Tracker interface {
    String() string
    Do(TrackerRequest) (TrackerResponse, error)
}
type TrackerRequest struct {
    InfoHash        InfoHash
    PeerID          PeerID
    IP              net.IP
    Port            uint16      
    Uploaded        uint64
    Downloaded      uint64
    Left            uint64
    Event           Event
}
type TrackerResponse struct {
    FailureReason   string
    Interval        uint16
    Peers           []Peer
    Other           map[string]interface{}
}
type Event string