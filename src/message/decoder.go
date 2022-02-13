package message

import (
	"bytes"
	"io"

	"github.com/azamaulanaaa/gotor/src/big_endian"
	"github.com/azamaulanaaa/gotor/src/bitfield"
)

func DecodeHandshake(r io.Reader) (Handshake, error) {
    var err error

    const minLength = 49

    var protocolLength uint8
    {
        buff, err := readUntil(r, 1)
        if err != nil {
            return Handshake{}, err
        }
        err = bigendian.Decode(buff, &protocolLength)
        if err != nil {
            return Handshake{}, err
        }
    }

    handshake := Handshake{}

    {
        buff, err := readUntil(r, int64(protocolLength))
        if err != nil {
            return Handshake{}, err
        }

        handshake.Protocol = string(buff)
    }

    var buff []byte

    buff, err = readUntil(r, int64(len(handshake.Reserved)))
    if err != nil {
        return Handshake{}, err
    }
    copy(handshake.Reserved[:], buff)
    
    buff, err = readUntil(r, int64(len(handshake.Infohash)))
    if err != nil {
        return Handshake{}, err
    }
    copy(handshake.Infohash[:], buff)

    buff, err = readUntil(r, int64(len(handshake.PeerID)))
    if err != nil {
        return Handshake{}, err
    }
    copy(handshake.PeerID[:], buff)

    return handshake, nil
}

func DecodeMessage(r io.Reader) (interface{}, error) {
    var messageLen uint32
    {
        buff, err := readUntil(r, 4)
        if err != nil {
            return nil, err
        }

        err = bigendian.Decode(buff, &messageLen)
        if err != nil {
            return nil, err
        }
    }

    if messageLen == 0 {
        return decodeKeepAlive()
    }

    var id messageID
    {
        buff, err := readUntil(r, 1)
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

    r = newLimitReader(r, int64(messageLen) - 1)

    switch id {
    case messageChoke:
        return decodeChoke(r)
    case messageUnChoke:
        return decodeUnChoke(r)
    case messageInterested:
        return decodeInterested(r)
    case messageNotInterested:
        return decodeNotInterested(r)
    case messageHave:
        return decodeHave(r)
    case messageBitfield:
        return decodeBitfield(r)
    case messageRequest:
        return decodeRequest(r)
    case messagePiece:
        return decodePiece(r)
    case messageCancel:
        return decodeCancel(r)
    }

    return nil, ErrorMessageInvalid
}

func decodeKeepAlive() (KeepAlive, error) {
    return KeepAlive{}, nil
}

func decodeChoke(r io.Reader) (Choke, error) {
    return Choke{}, nil
}

func decodeUnChoke(r io.Reader) (UnChoke, error) {
    return UnChoke{}, nil
}

func decodeInterested(r io.Reader) (Interested, error) {
    return Interested{}, nil
}

func decodeNotInterested(r io.Reader) (NotInterested, error) {
    return NotInterested{}, nil
}

func decodeHave(r io.Reader) (Have, error) {
    var err error

    const length = 1
   
    data, err := readUntil(r, length)
    if err != nil {
        return Have{}, err
    }

    var message Have

    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return Have{}, err
    }

    return message, nil
}

func decodeBitfield(r io.Reader) (Bitfield, error) {
    data, err := io.ReadAll(r)
    if err != nil {
        return Bitfield{}, err
    }

    theBitfield := bitfield.BitFieldFormBytes(data)

    return Bitfield{
        Bitfield: theBitfield,
    }, nil
}

func decodeRequest(r io.Reader) (Request, error) {
    var err error

    var message Request
    var data []byte

    data, err = readUntil(r, 4)
    if err != nil {
        return Request{}, err
    }
    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return Request{}, err
    }

    data, err = readUntil(r, 4)
    if err != nil {
        return Request{}, err
    }
    err = bigendian.Decode(data, &message.Begin)
    if err != nil {
        return Request{}, err
    }

    data, err = readUntil(r, 4)
    if err != nil {
        return Request{}, err
    }
    err = bigendian.Decode(data, &message.Length)
    if err != nil {
        return Request{}, err
    }

    return message, nil
}

func decodePiece(r io.Reader) (Piece, error) {
    var err error

    var message Piece
    var data []byte

    data, err = readUntil(r, 4)
    if err != nil {
        return Piece{}, err
    }
    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return Piece{}, err
    }

    data, err = readUntil(r, 4)
    if err != nil {
        return Piece{}, err
    }
    err = bigendian.Decode(data, &message.Begin)
    if err != nil {
        return Piece{}, err
    }

    data, err = io.ReadAll(r)
    if err != nil {
        return Piece{}, err
    }
    buff := bytes.NewReader(data)

    message.Piece = buff
    

    return message, nil
}

func decodeCancel(r io.Reader)(Cancel, error) {
    var err error

    var message Cancel
    var data []byte

    data, err = readUntil(r, 4)
    if err != nil {
        return Cancel{}, err
    }
    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return Cancel{}, err
    }

    data, err = readUntil(r, 4)
    if err != nil {
        return Cancel{}, err
    }
    err = bigendian.Decode(data, &message.Begin)
    if err != nil {
        return Cancel{}, err
    }

    data, err = readUntil(r, 4)
    if err != nil {
        return Cancel{}, err
    }
    err = bigendian.Decode(data, &message.Length)
    if err != nil {
        return Cancel{}, err
    }

    return message, nil
}
