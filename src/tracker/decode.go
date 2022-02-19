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
	copy(req.Me.PeerID[:], urlQuery.Get("peer_id"))

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

	req.Me.IP = net.ParseIP(urlQuery.Get("ip"))

	if urlQuery.Has("port") {
		value := urlQuery.Get("port")
		port, err := strconv.ParseUint(value, 10, 16)
		if err != nil {
			return Request{}, err
		}

		req.Me.Port = uint16(port)
	}

	return req, nil
}

func DecodeHTTPResponse(r io.Reader, iplen IPLen) (Response, error) {
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

		numPeers := len(rawPeersInByte) / int(iplen)
		for i := 0; i < numPeers; i++ {
			peer, err := peer.Decode(rawPeersInByte[i : i+int(iplen)])
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

func DecodeUDPResponse(r io.Reader, iplen IPLen) (interface{}, error) {
	header, action, err := decodeUDPResponseHeader(r)
	if err != nil {
		return nil, ErrorInvalidEvent
	}

	switch action {
	case udpActionConnect:
		return decodeUDPConnectResponse(r, header)
	case udpActionAnnounce:
		return decodeUDPAnnounceResponse(r, header, iplen)
	case udpActionErrors:
		return decodeUDPErrorsResponse(r, header)
	}

	return nil, ErrorInvalidResponse
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

		req.Me.PeerID = peerID
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
		_, err := r.Read(req.Me.IP)
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

		err = bigendian.Decode(buff, &req.Me.Port)
		if err != nil {
			return UDPAnnounceRequest{}, err
		}
	}

	return req, nil
}

func decodeUDPResponseHeader(r io.Reader) (UDPResponseHeader, udpAction, error) {
	header := UDPResponseHeader{}

	var action udpAction
	{
		data := make([]byte, 4)
		_, err := r.Read(data)
		if err != nil {
			return UDPResponseHeader{}, 0, err
		}

		err = bigendian.Decode(data, &action)
		if err != nil {
			return UDPResponseHeader{}, 0, err
		}
	}

	{
		data := make([]byte, 4)
		_, err := r.Read(data)
		if err != nil {
			return UDPResponseHeader{}, 0, err
		}

		err = bigendian.Decode(data, &header.TransactionID)
		if err != nil {
			return UDPResponseHeader{}, 0, err
		}
	}

	return header, action, nil
}

func decodeUDPConnectResponse(r io.Reader, header UDPResponseHeader) (UDPConnectResponse, error) {
	res := UDPConnectResponse{
		UDPResponseHeader: header,
	}

	{
		data := make([]byte, 8)
		_, err := r.Read(data)
		if err != nil {
			return UDPConnectResponse{}, err
		}

		err = bigendian.Decode(data, &res.ConnectionID)
		if err != nil {
			return UDPConnectResponse{}, err
		}
	}

	return res, nil
}

func decodeUDPAnnounceResponse(r io.Reader, header UDPResponseHeader, iplen IPLen) (UDPAnnounceResponse, error) {
	res := UDPAnnounceResponse{
		UDPResponseHeader: header,
	}

	{
		data := make([]byte, 4)
		_, err := r.Read(data)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}

		var interval32 int32

		err = bigendian.Decode(data, &interval32)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}

		res.Interval = time.Duration(interval32) * time.Second
	}

	{
		data := make([]byte, 4)
		_, err := r.Read(data)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}

		err = bigendian.Decode(data, &res.Leechers)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}
	}

	{
		data := make([]byte, 4)
		_, err := r.Read(data)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}

		err = bigendian.Decode(data, &res.Seeders)
		if err != nil {
			return UDPAnnounceResponse{}, err
		}
	}

	{
		res.Peers = []peer.Peer{}

		for {
			data := make([]byte, iplen)
			n, err := r.Read(data)
			if (err != nil && n == 0) || n != int(iplen) {
				break
			}
			if err != nil {
				return UDPAnnounceResponse{}, err
			}

			thePeer, err := peer.Decode(data)
			if err != nil {
				return UDPAnnounceResponse{}, ErrorInvalidResponse
			}

			res.Peers = append(res.Peers, thePeer)
		}
	}

	return res, nil
}

func decodeUDPErrorsResponse(r io.Reader, header UDPResponseHeader) (UDPErrorsResponse, error) {
	res := UDPErrorsResponse{
		UDPResponseHeader: header,
	}

	{
		message, err := io.ReadAll(r)
		if err != nil {
			return UDPErrorsResponse{}, err
		}

		res.Message = string(message)
	}

	return res, nil
}
