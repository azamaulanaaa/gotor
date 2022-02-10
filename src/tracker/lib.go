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
    queryMap := map[string]string{}

    if rawInfoHash, ok := req.InfoHash(); ok {
        queryMap["info_hash"] = string(rawInfoHash[:]) 
    }

    if rawPeerID, ok := req.PeerID(); ok {
        queryMap["peer_id"] = string(rawPeerID[:])
    }

    if rawIP, ok := req.IP(); ok {
        queryMap["ip"] = rawIP.String()
    }

    if rawPort, ok := req.Port(); ok {
        queryMap["port"] = strconv.FormatUint(uint64(rawPort), 10)
    }

    if rawUploaded, ok := req.Uploaded(); ok {
        queryMap["uploaded"] = strconv.FormatUint(rawUploaded, 10)
    }

    if rawDownloaded, ok := req.Downloaded(); ok {
        queryMap["downloaded"] = strconv.FormatUint(rawDownloaded, 10)
    }

    if rawLeft, ok := req.Left(); ok {
        queryMap["left"] = strconv.FormatUint(rawLeft, 10)
    }

    if rawEvent, ok := req.Event(); ok {
        queryMap["event"] = string(rawEvent)
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
        return nil, ErrorPeerBytesInvalid
    }
    ip := net.IPv4(value[0], value[1], value[2], value[3])
    port := binary.BigEndian.Uint16(value[4:6])

    peer := peer{
        ip: ip,
        port: port,
    }

    return &peer, nil
}
