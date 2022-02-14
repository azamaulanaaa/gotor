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

func EncodeRequest(req Request) (string, error) {
	queryMap := map[string]string{}

	queryMap["info_hash"] = string(req.Infohash[:])
	queryMap["peer_id"] = string(req.PeerID[:])
	queryMap["downloaded"] = strconv.FormatInt(req.Downloaded, 10)
	queryMap["left"] = strconv.FormatInt(req.Left, 10)
	queryMap["uploaded"] = strconv.FormatInt(req.Uploaded, 10)
	queryMap["event"] = req.Event.String()
	queryMap["ip"] = req.IP.String()
	queryMap["port"] = strconv.FormatUint(uint64(req.Port), 10)

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

    return queryStr, nil
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
