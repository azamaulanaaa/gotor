package message

import (
	"bytes"
	"io"

	"github.com/azamaulanaaa/gotor/src/big_endian"
)

func EncodeHandshake(handshake Handshake) (io.Reader, error) {
    var err error

    encodedProtocolLength, err := bigendian.Encode(uint8(len(handshake.Protocol)))
    if err != nil {
        return nil, err
    }

    r := &bytes.Buffer{}
    for _, theData := range [][]byte{
        encodedProtocolLength,
        []byte(handshake.Protocol),
        handshake.Reserved[:],
        handshake.Infohash[:],
        handshake.PeerID[:],
    } {
        r.Write(theData)
    }

    return r, nil
}

func EncodeMessage(rawMessage interface{}) (io.Reader, error) {
    switch message := rawMessage.(type) {
    case KeepAlive:
        return encodeKeepAlive()
    case Choke:
        return encodeChoke(message)
    case UnChoke:
        return encodeUnChoke(message)
    case Interested:
        return encodeInterested(message)
    case NotInterested:
        return encodeNotInterested(message)
    case Have:
        return encodeHave(message)
    case Bitfield:
        return encodeBitfield(message)
    case Request:
        return encodeRequest(message)
    case Piece:
        return encodePiece(message)
    case Cancel:
        return encodeCancel(message)
    default:
        return nil, ErrorMessageInvalid
    }
}

func encodeKeepAlive() (io.Reader, error) {
    encodedLenMessage, err := bigendian.Encode(uint32(0))
    if err != nil {
        return nil, err
    }

    return bytes.NewReader(encodedLenMessage), nil
}

func finisher(id messageID, data []byte) (io.Reader, error) {
    encodedLenMessage, err := bigendian.Encode(uint32(len(data) + 1))
    if err != nil {
        return nil, err
    }

    r := &bytes.Buffer{}
    for _, theData := range [][]byte{
        encodedLenMessage,
        []byte{byte(id)},
        data,
    }{
        r.Write(theData)
    }
    
    return r, nil
}

func encodeChoke(message Choke) (io.Reader, error) {
    return finisher(messageChoke, nil)
}

func encodeUnChoke(message UnChoke) (io.Reader, error) {
    return finisher(messageChoke, nil)
}

func encodeInterested(message Interested) (io.Reader, error) {
    return finisher(messageInterested, nil)
}

func encodeNotInterested(message NotInterested) (io.Reader, error) {
    return finisher(messageNotInterested, nil)
}

func encodeHave(message Have) (io.Reader, error) {
    var err error

    data := make([]byte, 4)
    data, err = bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    return finisher(messageHave, data)
}

func encodeBitfield(message Bitfield) (io.Reader, error) {
    data := message.Bitfield.AsBytes()

    return finisher(messageBitfield, data)
}

func encodeRequest(message Request) (io.Reader, error) {
    var err error

    if message.Length > MaxLength {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    encodedLength, err := bigendian.Encode(message.Length)
    if err != nil {
        return nil,  err
    }

    data := make([]byte, 0, 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        encodedLength,
    } {
        data = append(data, theData...)
    }

    return finisher(messageRequest, data)
}

func encodePiece(message Piece) (io.Reader, error) {
    var err error

    var piece *bytes.Buffer
    piece = &bytes.Buffer{}

    length, err := io.Copy(piece, newReaderReadAt(message.Piece, 0))
    if err != nil {
        return nil, err
    }

    if int(length) > int(MaxLength) {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    data := make([]byte, length + 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        piece.Bytes(),
    } {
        data = append(data, theData...)
    }

    return finisher(messagePiece, data)
}

func encodeCancel(message Cancel) (io.Reader, error) {
    var err error

    if message.Length > MaxLength {
        return nil, ErrorMessageTooLong
    }

    encodedIndex, err := bigendian.Encode(message.Index)
    if err != nil {
        return nil, err
    }

    encodedBegin, err := bigendian.Encode(message.Begin)
    if err != nil {
        return nil, err
    }

    encodedLength, err := bigendian.Encode(message.Length)
    if err != nil {
        return nil, err
    }

    data := make([]byte, 0, 12)
    for _, theData := range [][]byte{
        encodedIndex,
        encodedBegin,
        encodedLength,
    } {
        data = append(data, theData...)
    }

    return finisher(messageCancel, data)
}
