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
	"github.com/marksamman/bencode"
)

var (
    ErrorFailureReason = errors.New("error explained in failure reason field")
    ErrorBencodeInvalid = errors.New("value is not a valid bencode")
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

    data, err := bencode.Decode(strings.NewReader(value))
    if err != nil {
        return src.TrackerResponse{}, ErrorBencodeInvalid
    }
    var response src.TrackerResponse

    if value, ok := data["failure reason"].(string); ok {
        response.FailureReason = value
        return response, ErrorFailureReason
    }

    response.Interval = uint16(data["interval"].(int64))
    delete(data, "interval")

    if dataPeersStr, ok := data["peers"].(string); ok {
        dataPeers := []byte(dataPeersStr)

        numPeers := len(dataPeers) / 6
        for i := 0; i < numPeers; i++ {
            peer, err := decodeBytePeer(dataPeers[i:i+6])
            if err == nil {
                response.Peers = append(response.Peers, peer) 
            }
        }
    }
    delete(data, "peers")
    
    response.Other = data

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
