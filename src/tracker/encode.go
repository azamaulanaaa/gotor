package tracker

import (
	"net/url"
	"strconv"
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
