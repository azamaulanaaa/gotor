package tracker

import (
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/peer"
)

var (
    ErrorInvalidResponse = errors.New("value is not a valid response")
)

func requestQuery(req Request) string {
    queryMap := map[string]string{}

    {
        infohash := req.InfoHash()
        queryMap["info_hash"] = string(infohash[:]) 
    }

    {
        peerID := req.PeerID()
        queryMap["peer_id"] = string(peerID[:])
    }

    if rawIP, ok := req.IP(); ok {
        queryMap["ip"] = rawIP.String()
    }

    if rawPort, ok := req.Port(); ok {
        queryMap["port"] = strconv.FormatUint(uint64(rawPort), 10)
    }

    queryMap["uploaded"] = strconv.FormatUint(req.Uploaded(), 10)
    queryMap["downloaded"] = strconv.FormatUint(req.Downloaded(), 10)
    queryMap["left"] = strconv.FormatUint(req.Left(), 10)

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

func decodeResponse(value string) (Response, error) {
    var err error

    var rawResponse bencode.Dictionary
    {
        var rawData interface{}
        rawData, err = bencode.Decode(strings.NewReader(value))
        if err != nil {
           return nil, err
        }

        var ok bool
        rawResponse, ok = rawData.(bencode.Dictionary)
        if !ok {
            return nil, ErrorInvalidResponse
        }
    }
    
    var res response

    if rawFailureReason, ok := rawResponse["failure reason"].(bencode.String); ok {
        return nil, errors.New(string(rawFailureReason))
    }

    if rawInterval, ok := rawResponse["interval"].(bencode.Integer); ok {
        res.interval = time.Duration(rawInterval) * time.Millisecond
    }

    if rawPeers, ok := rawResponse["peers"].(bencode.String); ok {
        peers := []peer.Peer{}
        rawPeersInByte := []byte(rawPeers)

        numPeers := len(rawPeersInByte) / 6
        for i := 0; i < numPeers; i++ {
            peer, err := peer.NewPeerFromBytes(rawPeersInByte[i:i+6])
            if err == nil {
                peers = append(peers, peer) 
            }
        }

        res.peers = peers
    }
    
    return &res, nil
}
