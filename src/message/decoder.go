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

func DecodeMessage(r io.Reader) (interface{}, error) {
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

    if messageLen == 0 {
        return decodeKeepAlive()
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

func decodeKeepAlive() (KeepAlive, error) {
    return KeepAlive{}, nil
}

func decodeChoke(data []byte) (Choke, error) {
    return Choke{}, nil
}

func decodeUnChoke(data []byte) (UnChoke, error) {
    return UnChoke{}, nil
}

func decodeInterested(data []byte) (Interested, error) {
    return Interested{}, nil
}

func decodeNotInterested(data []byte) (NotInterested, error) {
    return NotInterested{}, nil
}

func decodeHave(data []byte) (Have, error) {
    var err error

    const length = 1
    
    if len(data) != length {
        return Have{}, ErrorMessageInvalid
    }

    var message Have

    err = bigendian.Decode(data, &message.Index)
    if err != nil {
        return Have{}, err
    }

    return message, nil
}

func decodeBitfield(data []byte) (Bitfield, error) {
    theBitfield := bitfield.BitFieldFormBytes(data)

    return Bitfield{
        Bitfield: theBitfield,
    }, nil
}

func decodeRequest(data []byte) (Reqeust, error) {
    var err error

    const length = 12

    if len(data) != length {
        return Reqeust{}, ErrorMessageInvalid
    }

    var message Reqeust

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return Reqeust{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return Reqeust{}, err
    }

    err = bigendian.Decode(data[8:12], &message.Length)
    if err != nil {
        return Reqeust{}, err
    }

    return message, nil
}

func decodePiece(data []byte) (Piece, error) {
    var err error

    const minLength = 8
    if len(data) < minLength {
        return Piece{}, ErrorMessageInvalid
    }

    var message Piece

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return Piece{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return Piece{}, err
    }

    message.Piece = data[8:]

    return message, nil
}

func decodeCancel(data []byte)(Cancel, error) {
    var err error

    const length = 12
    if len(data) != length {
        return Cancel{}, ErrorMessageInvalid
    }

    var message Cancel

    err = bigendian.Decode(data[0:4], &message.Index)
    if err != nil {
        return Cancel{}, err
    }

    err = bigendian.Decode(data[4:8], &message.Begin)
    if err != nil {
        return Cancel{}, err
    }

    err = bigendian.Decode(data[8:12], &message.Length)
    if err != nil {
        return Cancel{}, err
    }

    return message, nil
}
