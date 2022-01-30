package tracker

import (
	"errors"
)

var (
    ErrorProtocolUndefined = errors.New("protocol is not defined")
)

type Protocol interface {
    String() string
}

type protocol_impl uint

const (
    ProtocolHTTP protocol_impl = iota
    ProtocolHTTPS 
    ProtocolUDP 
)

var protocol_list = map[protocol_impl]string {
    ProtocolHTTP: "http",
    ProtocolHTTPS: "https",
    ProtocolUDP: "udp",
}

func NewProtocol(value string) (Protocol, error) {
    for protocol, protocolStr := range protocol_list {
        if protocolStr == value {
            return protocol, nil
        }
    }
    return nil, ErrorProtocolUndefined
}

func (protocol protocol_impl) String() string {
    return protocol_list[protocol]
}
