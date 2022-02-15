package tracker

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/azamaulanaaa/gotor/src/bencode"
	"github.com/azamaulanaaa/gotor/src/peer"
)

func EncodeRequest(req Request) (string, error) {
	urlQuery := url.Values{}

	urlQuery.Add("info_hash", string(req.Infohash[:]))
	urlQuery.Add("peer_id", string(req.PeerID[:]))
	urlQuery.Add("downloaded", strconv.FormatInt(req.Downloaded, 10))
	urlQuery.Add("left", strconv.FormatInt(req.Left, 10))
	urlQuery.Add("uploaded", strconv.FormatInt(req.Uploaded, 10))
	urlQuery.Add("event", req.Event.String())
	urlQuery.Add("ip", req.IP.String())
	urlQuery.Add("port", strconv.FormatUint(uint64(req.Port), 10))

	return urlQuery.Encode(), nil
}

func EncodeResponse(res Response, failure_reason error) (string, error) {
	var rawResponse bencode.Dictionary
	if failure_reason != nil {
		rawResponse["failure reason"] = bencode.String(failure_reason.Error())
		return bencode.Encode(rawResponse)
	}

	rawResponse["interval"] = bencode.Integer(res.Interval.Seconds())
	rawResponse["peers"] = ""
	for _, thePeer := range res.Peers {
		rawPeer, err := peer.Encode(thePeer)
		if err != nil {
			return "", err
		}

		rawResponse["peers"] = fmt.Sprintf("%s%s", rawResponse["peers"], rawPeer)
	}

	return bencode.Encode(rawResponse)
}
