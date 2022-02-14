package tracker

import (
	"errors"
	"net"
	"strconv"
	"strings"
	"time"

	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/peer"
)

func DecodeRequest(value string) (Request, error) {
    queryMap := map[string]string{}

    {
        splitedValue := strings.Split(value, "&")

        for _, v := range splitedValue {
            pair := strings.Split(v, "=") 
            
            queryMap[pair[0]] = ""
            if len(pair) > 1 {
                queryMap[pair[0]] = pair[1]
            }
        }
    }

    var req Request

    if value, ok := queryMap["info_hash"]; ok {
        copy(req.Infohash[:], []byte(value))
    }

    if value, ok := queryMap["peer_id"]; ok {
        copy(req.PeerID[:], []byte(value))
    }

    if value, ok := queryMap["downloaded"]; ok {
        downloaded, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Downloaded = downloaded
    }

    if value, ok := queryMap["left"]; ok {
        left, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Left = left
    }

    if value, ok := queryMap["uploaded"]; ok {
        uploaded, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Downloaded = uploaded
    }

    if value, ok := queryMap["event"]; ok {
        event, err := NewEvent(value)
        if err != nil {
            return Request{}, err
        }
        
        req.Event = event
    }

    if value, ok := queryMap["ip"]; ok {
        req.IP = net.ParseIP(value)
    }

    if value, ok := queryMap["port"]; ok {
        port, err := strconv.ParseUint(value, 10, 16)
        if err != nil {
            return Request{}, err
        }

        req.Port = uint16(port)
    }

    return req, nil
}

func DecodeResponse(value string) (Response, error) {
	var err error

	var rawResponse bencode.Dictionary
	{
		var rawData interface{}
		rawData, err = bencode.Decode(strings.NewReader(value))
		if err != nil {
			return Response{}, err
		}

		var ok bool
		rawResponse, ok = rawData.(bencode.Dictionary)
		if !ok {
			return Response{}, ErrorInvalidResponse
		}
	}

	var res Response

	if rawFailureReason, ok := rawResponse["failure reason"].(bencode.String); ok {
		return Response{}, errors.New(string(rawFailureReason))
	}

	if rawInterval, ok := rawResponse["interval"].(bencode.Integer); ok {
		res.Interval = time.Duration(rawInterval) * time.Second
	}

	if rawPeers, ok := rawResponse["peers"].(bencode.String); ok {
		peers := []peer.Peer{}
		rawPeersInByte := []byte(rawPeers)

		numPeers := len(rawPeersInByte) / 6
		for i := 0; i < numPeers; i++ {
			peer, err := peer.NewPeerFromBytes(rawPeersInByte[i : i+6])
			if err == nil {
				peers = append(peers, peer)
			}
		}

		res.Peers = peers
	}

	return res, nil
}
