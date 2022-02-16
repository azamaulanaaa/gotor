package tracker

import (
	"errors"
	"net"
	"time"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

var (
	ErrorTrackerInvalid        = errors.New("value is not a valid tracker")
	ErrorProtocolNotSuppported = errors.New("protocol not supported yet")
	ErrorUnableToConnect       = errors.New("unable connect to tracker server")
	ErrorInvalidRequest        = errors.New("value is not a valid request")
	ErrorInvalidResponse       = errors.New("value is not a valid response")
	ErrorInvalidEvent          = errors.New("value is not a valid event")
)

type Event int32

const (
	EventNone Event = iota
	EventCompleted
	EventStarted
	EventStopped
)

type Request struct {
	Infohash   hash.Hash
	PeerID     peer.PeerID
	Downloaded int64
	Left       int64
	Uploaded   int64
	Event      Event
	IP         net.IP
	Port       uint16
}

type Response struct {
	Interval time.Duration
	Peers    []peer.Peer
}

type udpAction int32

const (
	udpActionConnect udpAction = iota
	udpActionAnnounce
	udpActionScrape
	udpActionErrors
)

type UDPRequestHeader struct {
	ConnectionID  int64
	TransactionID int32
}

type UDPResponseHeader struct {
	TransactionID int32
}

type UDPConnectRequest struct {
	UDPRequestHeader
}

type UDPConnectResponse struct {
	UDPResponseHeader
	ConnectionID int64
}

type UDPAnnounceRequest struct {
	UDPRequestHeader
	Request
	Key        uint32
	NumWant    int32
	Extensions uint16
}

type UDPAnnounceResponse struct {
	UDPResponseHeader
	Response
	Leechers int32
	Seeders  int32
}

type UDPErrorsResponse struct {
	UDPResponseHeader
	Message string
}
