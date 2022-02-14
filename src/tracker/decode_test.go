//go:build unit

package tracker_test

import (
	"net"
	"testing"

	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
	"github.com/azamaulanaaa/gotor/src/tracker"
	"github.com/azamaulanaaa/gotor/test"
)

type requestTestData struct {
	Request        tracker.Request
	EncodedRequest string
}

func TestDecodeRequest(t *testing.T) {
	var infohash hash.Hash
	copy(infohash[:], "ahashshoudlworkswell")

	var peerID peer.PeerID
	copy(peerID[:], "otherhashshouldworks")

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
		"info_hash=ahashshoudlworkswell&downloaded=32&ip=123.123.123.123&event=completed&port=80&left=12&completed=0&peer_id=otherhashshouldworks",
	}

	out, err := tracker.DecodeRequest(testsData.EncodedRequest)
	test.Ok(t, err)
	test.Equals(t, testsData.Request, out)
}
