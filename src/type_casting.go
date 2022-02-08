package src

import (
    "net"
)

type PeerID [20]byte
type Hash [20]byte
type Event string

type Peer interface {
    SetPeerID(peerid PeerID)
    GetPeerID() (PeerID, bool)
    SetIP(ip net.IP)
    GetIP() (net.IP, bool)
    SetPort(port uint16)
    GetPort() (uint16, bool)
}

type Tracker interface {
    String() string
    Do(TrackerRequest) (TrackerResponse, error)
}
type TrackerRequest interface {
    SetInfoHash(hash Hash)
    GetInfoHash() (Hash, bool)
    SetPeerID(peerID PeerID)
    GetPeerID() (PeerID, bool)
    SetIP(ip net.IP)
    GetIP() (net.IP, bool)
    SetPort(port uint16)
    GetPort() (uint16, bool)
    SetUploaded(uploaded uint64)
    GetUploaded() (uint64, bool)
    SetDownloaded(downloaded uint64)
    GetDownloaded() (uint64, bool)
    SetLeft(left uint64)
    GetLeft() (uint64, bool)
    SetEvent(event Event)
    GetEvent() (Event, bool)
}
type TrackerResponse struct {
    FailureReason   string
    Interval        uint16
    Peers           []Peer
}

type Metainfo interface {
    Announce() string
    Info() MetainfoInfo
}
type MetainfoInfo interface {
    PieceLength() uint64
    Pieces() []Hash
    Length() (uint64, bool)
    Files() ([]MetainfoFile, bool)
    Name() (string, bool)
    Private() (bool, bool)
}
type MetainfoFile interface {
    Length() uint64
    Path() string
}

type Bitfield interface {
    Set(index uint64, value bool) error
    Get(index uint64) (bool, error)
    Length() uint64
}
