package messagehandler

import (
	"errors"

	"github.com/azamaulanaaa/gotor/src/bitfield"
	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

type MessageHandler interface {
    SendHandshake(handshake Handshake)  error
    SendMessage(message interface{})    error
    Close()                             error
}

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

type MessageChoke struct {}
type MessageUnChoke struct {}
type MessageInterested struct {}
type MessageNotInterested struct {}
type MessageHave struct {
    Index   uint32
}
type MessageBitfield struct{
    Bitfield bitfield.Bitfield
}
type MessageRequest struct {
    Index   uint32
    Begin   uint32
    Length  uint32
}
type MessagePiece struct {
    Index   uint32
    Begin   uint32
    Piece   []byte
}
type MessageCancel struct {
    Index   uint32
    Begin   uint32
    Length  uint32
}

const (
    MaxLength = uint32(16383)
)

var (
    ErrorMessageTooLong = errors.New("maximum message length is 2^14")
    ErrorMessageInvalid = errors.New("message is invalid")
)
