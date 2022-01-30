package tracker

import (
	"net"

    "github.com/azamaulanaaa/gotor/src/lib"
)

type Request struct {
    InfoHash        InfoHash    `url:"info_hash"`
    PeerId          string      `url:"peer_id"`
    IP              net.IP      `url:"ip"`
    Port            uint16      `url:"port"` 
    Uploaded        uint64      `url:"uploaded"`
    Downloaded      uint64      `url:"downloaded"`
    Left            uint64      `url:"left"`
    Event           Event       `url:"event"`
}

func (request Request) String() string {
    queryString, _ := lib.QueryString(request, "url")
    return queryString
}

