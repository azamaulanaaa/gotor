package src

import (
    "net"
)

type PeerID     [20]byte
type Hash       [20]byte
type Event      string

type Peer interface {
    PeerID()    (PeerID, bool)
    IP()        net.IP
    Port()      uint16
}

type Tracker interface {
    String()            string
    Do(TrackerRequest)  (TrackerResponse, error)
}
type TrackerRequest interface {
    InfoHash()      Hash
    PeerID()        PeerID
    IP()            (net.IP, bool)
    Port()          (uint16, bool)
    Uploaded()      uint64
    Downloaded()    uint64
    Left()          uint64
    Event()         (Event, bool)
}
type TrackerResponse interface {
    Interval()      uint16
    Peers()         []Peer
}

type Metainfo interface {
    Announce()      string
    Info()          MetainfoInfo
}
type MetainfoInfo interface {
    PieceLength()   uint32
    Pieces()        []Hash
    Length()        (uint64, bool)
    Files()         ([]MetainfoFile, bool)
    Name()          (string, bool)
    Private()       (bool, bool)
}
type MetainfoFile interface {
    Length() uint64
    Path() string
}

type Bitfield interface {
    Set(index uint32, value bool)   error
    Get(index uint32)               (bool, error)
    Length()                        uint32
    AsBytes()                       []byte
}
