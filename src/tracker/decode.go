package tracker

import (
	"errors"
	"io"
	"net"
	"net/url"
	"strconv"
	"time"

	"github.com/azamaulanaaa/gotor/src/bencode"
	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
	"github.com/azamaulanaaa/gotor/src/hash"
	"github.com/azamaulanaaa/gotor/src/peer"
)

func DecodeHTTPRequest(r io.Reader) (Request, error) {
	var urlQuery url.Values
	{
		value, err := io.ReadAll(r)
		if err != nil {
			return Request{}, err
		}

		urlQuery, err = url.ParseQuery(string(value))
		if err != nil {
			return Request{}, err
		}
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
		value := urlQuery.Get("port")
		port, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return Request{}, err
		}

		req.Port = uint16(port)
	}

	return req, nil
}

func DecodeHTTPResponse(r io.Reader) (Response, error) {
	var rawResponse bencode.Dictionary
	{
		var rawData interface{}
		rawData, err := bencode.Decode(r)
		if err != nil {
			return Response{}, err
		}

		var ok bool
		rawResponse, ok = rawData.(bencode.Dictionary)
		if !ok {
			return Response{}, err
		}
	}

	if rawFailureReason, ok := rawResponse["failure reason"].(bencode.String); ok {
		return Response{}, errors.New(string(rawFailureReason))
	}

	var res Response

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

func DecodeUDPRequest(r io.Reader) (interface{}, error) {
	header, action, err := decodeUDPRequestHeader(r)
	if err != nil {
		return nil, err
	}

	switch action {
	case udpActionConnect:
		return decodeUDPConnectRequest(r, header)
	case udpActionAnnounce:
		return decodeUDPAnnounceRequest(r, header)
	}

	return nil, ErrorInvalidRequest
}

func decodeUDPRequestHeader(r io.Reader) (UDPRequestHeader, udpAction, error) {
	header := UDPRequestHeader{}

	{
		buff := make([]byte, 8)
		_, err := r.Read(buff)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}

		err = bigendian.Decode(buff, &header.ConnectionID)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}
	}

	var action udpAction

	{
		buff := make([]byte, 4)
		_, err := r.Read(buff)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}

		err = bigendian.Decode(buff, &action)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}
	}

	{
		buff := make([]byte, 4)
		_, err := r.Read(buff)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}

		err = bigendian.Decode(buff, &header.TransactionID)
		if err != nil {
			return UDPRequestHeader{}, 0, err
		}
	}

	return header, action, nil
}

func decodeUDPConnectRequest(r io.Reader, header UDPRequestHeader) (UDPConnectRequest, error) {
	req := UDPConnectRequest{
		header,
	}

	return req, nil
}

func decodeUDPAnnounceRequest(r io.Reader, header UDPRequestHeader) (UDPAnnounceRequest, error) {
	req := UDPAnnounceRequest{
		UDPRequestHeader: header,
	}

	{
		var infohash hash.Hash
		_, err := r.Read(infohash[:])
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		req.Infohash = infohash
	}

	{
		var peerID peer.PeerID
		_, err := r.Read(peerID[:])
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		req.PeerID = peerID
	}

	{
		buff := make([]byte, 8)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Downloaded)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 8)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Left)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 8)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Uploaded)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 4)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Event)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		_, err := r.Read(req.IP)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 4)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Key)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 4)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.NumWant)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	{
		buff := make([]byte, 2)
		_, err := r.Read(buff)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}

		err = bigendian.Decode(buff, &req.Port)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	return req, nil
}
