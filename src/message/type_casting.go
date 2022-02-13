package message

import (
	"errors"
	"io"

	"github.com/azamaulanaaa/gotor/src/bitfield"
	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

const (
    MaxLength = uint32(16383)
)

var (
    ErrorMessageTooLong = errors.New("maximum message length is 2^14")
    ErrorMessageInvalid = errors.New("message is invalid")
)

type Reserved [8]byte

type Handshake struct {
    Protocol    string
    Reserved    Reserved
    Infohash    hash.Hash
    PeerID      peer.PeerID
}

type KeepAlive struct {}

type messageID uint

const (
    messageChoke messageID = iota
    messageUnChoke
    messageInterested
    messageNotInterested
    messageHave
    messageBitfield
    messageRequest
    messagePiece
    messageCancel
)

type Choke struct {}
type UnChoke struct {}
type Interested struct {}
type NotInterested struct {}
type Have struct {
    Index   uint32
}
type Bitfield struct{
    Bitfield bitfield.Bitfield
}
type Request struct {
    Index   uint32
    Begin   uint32
    Length  uint32
}
type Piece struct {
    Index   uint32
    Begin   uint32
    Piece   io.ReaderAt
}
type Cancel struct {
    Index   uint32
    Begin   uint32
    Length  uint32
}

type MessageHandler func(message interface{}) error
type HandshakeHandler func(handshake Handshake) error
