package tracker

import (
	"bytes"
	"fmt"
	"io"
	"net/url"
	"strconv"
	"strings"

	"github.com/azamaulanaaa/gotor/src/bencode"
	bigendian "github.com/azamaulanaaa/gotor/src/big_endian"
	"github.com/azamaulanaaa/gotor/src/peer"
)

func EncodeHTTPRequest(req Request) (io.Reader, error) {
	urlQuery := url.Values{}

	urlQuery.Add("info_hash", string(req.Infohash[:]))
	urlQuery.Add("peer_id", string(req.PeerID[:]))
	urlQuery.Add("downloaded", strconv.FormatInt(req.Downloaded, 10))
	urlQuery.Add("left", strconv.FormatInt(req.Left, 10))
	urlQuery.Add("uploaded", strconv.FormatInt(req.Uploaded, 10))
	urlQuery.Add("event", req.Event.String())
	urlQuery.Add("ip", req.IP.String())
	urlQuery.Add("port", strconv.FormatUint(uint64(req.Port), 10))

	r := strings.NewReader(urlQuery.Encode())

	return r, nil
}

func EncodeHTTPResponse(res Response, failure_reason error) (io.Reader, error) {
	var rawResponse bencode.Dictionary
	if failure_reason != nil {
		rawResponse["failure reason"] = bencode.String(failure_reason.Error())
	} else {
		rawResponse["interval"] = bencode.Integer(res.Interval.Seconds())
		rawResponse["peers"] = ""
		for _, thePeer := range res.Peers {
			rawPeer, err := peer.Encode(thePeer)
			if err != nil {
				return nil, err
			}

			rawResponse["peers"] = fmt.Sprintf("%s%s", rawResponse["peers"], rawPeer)
		}
	}

	ResStr, err := bencode.Encode(rawResponse)
	if err != nil {
		return nil, err
	}

	r := strings.NewReader(ResStr)

	return r, nil
}

func EncodeUDPRequest(rawRequest interface{}) (io.Reader, error) {
	switch request := rawRequest.(type) {
	case UDPConnectRequest:
		return encodeUDPConnectRequest(request)
	case UDPAnnounceRequest:
		return encodeUDPAnnounceRequest(request)
	}

	return nil, ErrorInvalidRequest
}

func EncodeUDPResponse(rawResponse interface{}) (io.Reader, error) {
	switch response := rawResponse.(type) {
	case UDPConnectResponse:
		return encodeUDPConnectResponse(response)
	case UDPAnnounceResponse:
		return encodeUDPAnnounceResponse(response)
	case UDPErrorsResponse:
		return encodeUDPErrorResponse(response)
	}

	return nil, ErrorInvalidResponse
}

func encodeUDPRequestHeader(header UDPRequestHeader, action udpAction) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		data, err := bigendian.Encode(header.ConnectionID)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(action)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(header.TransactionID)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff, nil
}

func encodeUDPConnectRequest(req UDPConnectRequest) (io.Reader, error) {
	return encodeUDPRequestHeader(req.UDPRequestHeader, udpActionConnect)
}

func encodeUDPAnnounceRequest(req UDPAnnounceRequest) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		header, err := encodeUDPRequestHeader(req.UDPRequestHeader, udpActionAnnounce)
		if err != nil {
			return nil, err
		}

		io.Copy(buff, header)
	}

	buff.Write(req.Request.Infohash[:])
	buff.Write(req.Request.PeerID[:])

	{
		data, err := bigendian.Encode(req.Downloaded)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(req.Left)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(req.Uploaded)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(req.Event)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		thePeer := peer.Peer{
			IP:   req.IP,
			Port: req.Port,
		}

		data, err := peer.Encode(thePeer)
		if err != nil {
			return nil, err
		}

		buff.Write(data[:len(data)-2])
	}

	{
		data, err := bigendian.Encode(req.Key)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(req.NumWant)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		thePeer := peer.Peer{
			IP:   req.IP,
			Port: req.Port,
		}

		data, err := peer.Encode(thePeer)
		if err != nil {
			return nil, err
		}

		buff.Write(data[len(data)-2:])
	}

	{
		data, err := bigendian.Encode(req.Extensions)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff, nil
}

func encodeUDPResponseHeader(header UDPResponseHeader, action udpAction) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		data, err := bigendian.Encode(action)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(header.TransactionID)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff, nil
}

func encodeUDPConnectResponse(res UDPConnectResponse) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		data, err := encodeUDPResponseHeader(res.UDPResponseHeader, udpActionConnect)
		if err != nil {
			return nil, err
		}

		io.Copy(buff, data)
	}

	{
		data, err := bigendian.Encode(res.ConnectionID)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff, nil
}

func encodeUDPAnnounceResponse(res UDPAnnounceResponse) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		data, err := encodeUDPResponseHeader(res.UDPResponseHeader, udpActionAnnounce)
		if err != nil {
			return nil, err
		}

		io.Copy(buff, data)
	}

	{
		data, err := bigendian.Encode(int32(res.Interval.Seconds()))
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(res.Leechers)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	{
		data, err := bigendian.Encode(res.Seeders)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	for _, thePeer := range res.Peers {
		data, err := peer.Encode(thePeer)
		if err != nil {
			return nil, err
		}

		buff.Write(data)
	}

	return buff, nil
}

func encodeUDPErrorResponse(res UDPErrorsResponse) (io.Reader, error) {
	buff := &bytes.Buffer{}

	{
		data, err := encodeUDPResponseHeader(res.UDPResponseHeader, udpActionError)
		if err != nil {
			return nil, err
		}

		io.Copy(buff, data)
	}

	buff.Write([]byte(res.Message))

	return buff, nil
}
