package connection

import (
	"bytes"
	"encoding/binary"
    "errors"
	"io"

	"github.com/azamaulanaaa/gotor/src"
)

var (
    ErrorTypeNotSupported = errors.New("type not supported yet")
)

func parseHandshake(conn io.Reader) (src.Handshake, error) {
    var lenpstr uint
    {
        buf := make([]byte, 1)
        _, err := conn.Read(buf)
        if err != nil {
            return src.Handshake{}, err
        }
        lenpstr = uint(buf[0])
    }

    pstr := make([]byte, lenpstr)
    {
        _, err := conn.Read(pstr)
        if err != nil {
            return src.Handshake{}, err
        }
    }

    reserved := src.Reserved{}
    {
        _, err := conn.Read(reserved[:])
        if err != nil {
            return src.Handshake{}, err
        }
    }

    infoHash := src.Hash{}
    {
        _, err := conn.Read(infoHash[:])
        if err != nil {
            return src.Handshake{}, err
        }
    }

    peerID := src.PeerID{}
    {
        _, err := conn.Read(peerID[:])
        if err != nil {
            return src.Handshake{}, err
        }
    }

    return src.Handshake{
        Protocol: pstr,
        Reserved: reserved,
        InfoHash: infoHash,
        PeerID: peerID,
    }, nil
}

func parseMessage(conn io.Reader) (src.Message, error) {
    var lenMessage uint32
    {
        buf := make([]byte, 4)
        _, err := conn.Read(buf)
        if err != nil {
            return src.Message{}, err
        }
        lenMessage = binary.BigEndian.Uint32(buf) - 1
    }

    var messageID src.MessageID
    {
        buf := make([]byte, 1)
        _, err := conn.Read(buf)
        if err != nil {
            return src.Message{}, err
        }
        messageID = src.MessageID(buf[0])
    }

    payload := make([]byte, lenMessage)
    _, err := conn.Read(payload)
    if err != nil {
        return src.Message{}, err
    }

    return src.Message{
        MessageID: messageID,
        Payload: payload,
    }, nil
}

func EncodeBigEndian(value interface{}) ([]byte, error) {
    var err error

    var buff bytes.Buffer
    err = binary.Write(&buff, binary.BigEndian, value)
    if err != nil {
        return nil, ErrorTypeNotSupported
    }

    return buff.Bytes(), nil
}

func DecodeBigEndian(b []byte) (interface{}, error) {
    var err error
    
    var out interface{}
    err = binary.Read(bytes.NewBuffer(b), binary.BigEndian, &out)

    return out, err
}
