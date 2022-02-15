package tracker

import (
	"errors"
	"net"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/peer"
)

func DecodeRequest(value string) (Request, error) {
    urlQuery, err := url.ParseQuery(value)
    if err != nil {
        return Request{}, err
    }

    var req Request

    copy(req.Infohash[:], urlQuery.Get("info_hash"))
    copy(req.PeerID[:], urlQuery.Get("peer_id"))

    if urlQuery.Has("downloaded") {
        value := urlQuery.Get("downloaded")
        downloaded, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Downloaded = downloaded
    }

    if urlQuery.Has("left") {
        value := urlQuery.Get("left")
        left, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Left = left
    }

    if urlQuery.Has("uploaded") {
        value := urlQuery.Get("uploaded")
        uploaded, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            return Request{}, err
        }

        req.Uploaded = uploaded
    }

    {
        event, err := NewEvent(urlQuery.Get("event"))
        if err != nil {
            return Request{}, err
        }
        
        req.Event = event
    }

    req.IP = net.ParseIP(urlQuery.Get("ip"))

    if urlQuery.Has("port") {
        value = urlQuery.Get("port")
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
			peer, err := peer.Decode(rawPeersInByte[i : i+6])
			if err == nil {
				peers = append(peers, peer)
			}
		}

		res.Peers = peers
	}

	return res, nil
}
