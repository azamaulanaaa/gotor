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
type Hash [20]byte
type Tracker interface {
    String() string
    Do(TrackerRequest) (TrackerResponse, error)
}
type TrackerRequest struct {
    InfoHash        Hash
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
}
type Event string

type Metainfo interface {
    Announce() string
    Info() Info
}
type Info interface {
    PieceLength() uint64
    Pieces() []Hash
    Length() (uint64, bool)
    Files() ([]File, bool)
    Name() (string, bool)
    Private() (bool, bool)
}
type File interface {
    Length() uint64
    Path() string
}
