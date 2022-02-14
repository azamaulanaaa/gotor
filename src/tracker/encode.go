package tracker

import (
	"fmt"
	"net/url"
	"strconv"
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
