package connection

import "github.com/azamaulanaaa/gotor/src"

const (
    MessageChoke src.MessageID = iota
    MessageUnChoke
    MessageInterested
    MessageNotInterested
    MessageHave
    MessageBitfield
    MessageRequest
    MessagePiece
    MessageCancel
)

