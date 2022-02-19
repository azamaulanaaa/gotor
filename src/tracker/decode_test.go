//go:build unit

package tracker_test

import (
	"net"
	"net/url"
	"strings"
	"testing"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
	"github.com/azamaulanaaa/gotor/src/tracker"
	"github.com/azamaulanaaa/gotor/test"
)

type requestTestData struct {
	request  tracker.Request
	UrlQuery url.Values
}

func TestDecodeRequest(t *testing.T) {
	var infohash hash.Hash
	copy(infohash[:], "ahashshoudlworkswell")

	var peerID peer.PeerID
	copy(peerID[:], "otherhashshouldworks")

	var urlQuery url.Values
	urlQuery, err := url.ParseQuery("info_hash=ahashshoudlworkswell&downloaded=32&ip=123.123.123.123&event=completed&port=80&left=12&uploaded=0&peer_id=otherhashshouldworks")
	test.Ok(t, err)

	testsData := requestTestData{
		tracker.Request{
			infohash,
			32,
			12,
			0,
			tracker.EventCompleted,
			peer.Peer{
				peerID,
				net.IPv4(123, 123, 123, 123),
				80,
			},
		},
		urlQuery,
	}

	out, err := tracker.DecodeHTTPRequest(strings.NewReader(testsData.UrlQuery.Encode()))
	test.Ok(t, err)
	test.Equals(t, testsData.request, out)
}
