//go:build unit

package tracker_test

import (
	"io"
	"net"
	"net/url"
	"testing"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
	"github.com/azamaulanaaa/gotor/src/tracker"
	"github.com/azamaulanaaa/gotor/test"
)

func TestEncodeRequest(t *testing.T) {
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
			peerID,
			32,
			12,
			0,
			tracker.EventCompleted,
			net.IPv4(123, 123, 123, 123),
			80,
		},
		urlQuery,
	}

	var out string
	{
		rawOut, err := tracker.EncodeHTTPRequest(testsData.request)
		test.Ok(t, err)
		outBytes, err := io.ReadAll(rawOut)
		test.Ok(t, err)
		out = string(outBytes)
	}

	test.Equals(t, testsData.UrlQuery.Encode(), out)
}
