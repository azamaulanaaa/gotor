package tracker

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/url"
	"strconv"
	"strings"

	"github.com/azamaulanaaa/gotor/src"
	"github.com/azamaulanaaa/gotor/src/bencode"
)

var (
    ErrorFailureReason = errors.New("error explained in failure reason field")
    ErrorInvalidResponse = errors.New("value is not a valid response")
    ErrorPeerBytesInvalid = errors.New("data byte of peer should be 6 bytes")
)

func requestQuery(req src.TrackerRequest) string {
    queryMap := map[string]string {
        "info_hash": string(req.InfoHash[:]),
        "peer_id": string(req.PeerID[:]),
        "ip": req.IP.String(),
        "port": strconv.FormatUint(uint64(req.Port), 10),
        "uploaded": strconv.FormatUint(req.Uploaded, 10),
        "downloaded": strconv.FormatUint(req.Downloaded, 10),
        "left": strconv.FormatUint(req.Left, 10),
        "event": string(req.Event),
    }
    queryStr := ""
    for key, value := range queryMap {
        junction := ""
        if queryStr != "" {
            junction = "&"
        }
        queryStr = fmt.Sprintf("%s%s%s=%s",
            queryStr,
            junction,
            key,
            url.QueryEscape(value),
        )

    }
    return queryStr
}

func decodeResponse(value string) (src.TrackerResponse, error) {
    var err error

    var rawResponse bencode.Dictionary
    {
        var rawData interface{}
        rawData, err = bencode.Decode(strings.NewReader(value))
        if err != nil {
            return src.TrackerResponse{}, err
        }

        var ok bool
        rawResponse, ok = rawData.(bencode.Dictionary)
        if !ok {
            return src.TrackerResponse{}, ErrorInvalidResponse
        }
    }
    
    var response src.TrackerResponse

    if rawFailureReason, ok := rawResponse["failure reason"].(bencode.String); ok {
        response.FailureReason = string(rawFailureReason)
        return response, ErrorFailureReason
    }

    if rawInterval, ok := rawResponse["interval"].(bencode.Integer); ok {
        response.Interval = uint16(rawInterval)
    }

    if rawPeers, ok := rawResponse["peers"].(bencode.String); ok {
        peers := []src.Peer{}
        rawPeersInByte := []byte(rawPeers)

        numPeers := len(rawPeersInByte) / 6
        for i := 0; i < numPeers; i++ {
            peer, err := decodeBytePeer(rawPeersInByte[i:i+6])
            if err == nil {
                peers = append(peers, peer) 
            }
        }

        response.Peers = peers
    }
    
    return response, nil
}

func decodeBytePeer(value []byte) (src.Peer, error) {
    if len(value) != 6 {
        return src.Peer{ }, ErrorPeerBytesInvalid
    }
    ip := net.IPv4(value[0], value[1], value[2], value[3])
    port := binary.BigEndian.Uint16(value[4:6])

    return src.Peer{
        IP: ip,
        Port: port,
    }, nil
}
