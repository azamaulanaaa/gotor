package messagehandler

import (
	"bytes"
	"io"

	"github.com/azamaulanaaa/gotor/src/big_endian"
	"github.com/azamaulanaaa/gotor/src/bitfield"
)

func decodeHandshake(r io.Reader) (Handshake, error) {
    var err error

    const minLength = 49

    var protocolLength uint8
    {
        buff := make([]byte, 1)
        for {
            _, err = r.Read(buff)
            if err == io.EOF {
                continue
            }
            if err != nil {
                return Handshake{}, err
            }
            break
        }
        err = bigendian.Decode(buff, &protocolLength)
        if err != nil {
            return Handshake{}, err
        }
    }

    {
        buff := make([]byte, int(protocolLength) + minLength)
        for {
            _, err = r.Read(buff)
            if err == io.EOF {
                continue
            }
            if err != nil {
                return Handshake{}, err
            }
            break
        }
        r = bytes.NewReader(buff)
    }

    handshake := Handshake{}

    {
        buff := make([]byte, protocolLength)
        _, err = r.Read(buff)
        if err != nil {
            return Handshake{}, err
        }

        handshake.Protocol = string(buff)
    }

    _, err = r.Read(handshake.Reserved[:])
    if err != nil {
        return Handshake{}, err
    }
    
    _, err = r.Read(handshake.Infohash[:])
    if err != nil {
        return Handshake{}, err
    }

    _, err = r.Read(handshake.PeerID[:])
    if err != nil {
        return Handshake{}, err
    }

    return handshake, nil
}

func decodeMessage(r io.Reader) (interface{}, error) {
    var err error

    var messageLen uint32
    {
        buff := make([]byte, 4)
        for {
            _, err = r.Read(buff)
            if err == io.EOF {
                continue
            }
            if err != nil {
                return nil, err
            }

            break
        }

        err = bigendian.Decode(buff, &messageLen)
        if err != nil {
            return nil, err
        }
    }

    {
        buff := make([]byte, messageLen)
        for {
            _, err = r.Read(buff)
            if err == io.EOF {
                continue
            }
            if err != nil {
                return nil, err
            }

            break
        }

        r = bytes.NewReader(buff)
    }

    var id messageID
    {
        buff := make([]byte, 1)
        _, err = r.Read(buff)
        if err != nil {
            return nil, err
        }

        var rawID uint8
        err = bigendian.Decode(buff, &rawID)
        if err != nil {
            return nil, err
        }

        id = messageID(rawID)
    }

    data := make([]byte, messageLen - 1)
    _, err = r.Read(data)
    if err != nil && err != io.EOF {
        return nil, err
    }

    switch id {
    case messageChoke:
        return decodeChoke(data)
    case messageUnChoke:
        return decodeUnChoke(data)
    case messageInterested:
        return decodeInterested(data)
    case messageNotInterested:
        return decodeNotInterested(data)
    case messageHave:
        return decodeHave(data)
    case messageBitfield:
        return decodeBitfield(data)
    case messageRequest:
        return decodeRequest(data)
    case messagePiece:
        return decodePiece(data)
    case messageCancel:
        return decodeCancel(data)
    }

    return nil, ErrorMessageInvalid
}

func decodeChoke(data []byte) (MessageChoke, error) {
    return MessageChoke{}, nil
}

func decodeUnChoke(data []byte) (MessageUnChoke, error) {
    return MessageUnChoke{}, nil
}

func decodeInterested(data []byte) (MessageInterested, error) {
    return MessageInterested{}, nil
}

func decodeNotInterested(data []byte) (MessageNotInterested, error) {
    return MessageNotInterested{}, nil
}

func decodeHave(data []byte) (MessageHave, error) {
    var err error

    const length = 1
    
    if len(data) != length {
        return MessageHave{}, ErrorMessageInvalid
    }

    var message MessageHave

    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return MessageHave{}, err
    }

    return message, nil
}

func decodeBitfield(data []byte) (MessageBitfield, error) {
    theBitfield := bitfield.BitFieldFormBytes(data)

    return MessageBitfield{
        Bitfield: theBitfield,
    }, nil
}

func decodeRequest(data []byte) (MessageRequest, error) {
    var err error

    const length = 12

    if len(data) != length {
        return MessageRequest{}, ErrorMessageInvalid
    }

    var message MessageRequest

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return MessageRequest{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return MessageRequest{}, err
    }

    err = bigendian.Decode(data[8:12], &message.Length)
    if err != nil {
        return MessageRequest{}, err
    }

    return message, nil
}

func decodePiece(data []byte) (MessagePiece, error) {
    var err error

    const minLength = 8
    if len(data) < minLength {
        return MessagePiece{}, ErrorMessageInvalid
    }

    var message MessagePiece

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return MessagePiece{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return MessagePiece{}, err
    }

    message.Piece = data[8:]

    return message, nil
}

func decodeCancel(data []byte)(MessageCancel, error) {
    var err error

    const length = 12
    if len(data) != length {
        return MessageCancel{}, ErrorMessageInvalid
    }

    var message MessageCancel

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return MessageCancel{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return MessageCancel{}, err
    }

    err = bigendian.Decode(data[8:12], &message.Length)
    if err != nil {
        return MessageCancel{}, err
    }

    return message, nil
}
