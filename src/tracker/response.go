package tracker

import (
	"encoding/binary"
	"errors"
	"net"
	"strings"

	"github.com/marksamman/bencode"
)

var (
    ErrorBencodeInvalid = errors.New("value is not a valid bencode")
)

type Response struct {
    Interval    uint16
    Peers       []Peer
    Other       map[string]interface{}
}

func (response *Response) FromBencode(value string) error {
    var err error

    data, err := bencode.Decode(strings.NewReader(value))
    if err != nil {
        return ErrorBencodeInvalid
    }

    response.Interval = uint16(data["interval"].(int64))
    delete(data, "interval")

    if dataPeersStr, ok := data["peers"].(string); ok {
        dataPeers := []byte(dataPeersStr)

        numPeers := len(dataPeers) / 6
        for i := 0; i < numPeers; i++ {
            ip := net.IPv4(dataPeers[i], dataPeers[i+1], dataPeers[+2], dataPeers[i+3])
            port := binary.BigEndian.Uint16(dataPeers[i+4:i+6])

            response.Peers = append(response.Peers, Peer{
                IP: ip,
                Port: port,
            })
        }
    }
    delete(data, "peers")
    
    response.Other = data

    return nil
}
