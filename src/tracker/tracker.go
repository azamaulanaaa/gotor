package tracker

import (
    "errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"

	"github.com/azamaulanaaa/gotor/src"
)

var (
    ErrorTrackerInvalid = errors.New("value is not a valid tracker")
    ErrorProtocolNotSuppported = errors.New("protocol not supported yet")
)

type tracker_impl struct {
    protocol    Protocol
    host        string
    port        uint16
    path        string
    httpClient  *http.Client
}    

func NewTracker(value string) (src.Tracker, error) {
    var err error

    tracker := tracker_impl{
        httpClient:     http.DefaultClient,
    } 

    var match []string
    {
        regex_pattern := regexp.MustCompile(`^([\w]+):\/\/([\w\d.-]+)(:[\d]+)?(\/(.*)+)?$`)
        matches := regex_pattern.FindAllStringSubmatch(value, 1)
        if len(matches) != 1 {
            return nil, ErrorTrackerInvalid
        }
        match = matches[0]
    }


    tracker.protocol, err = NewProtocol(match[1])
    if err != nil {
        return nil, ErrorTrackerInvalid
    }
    
    tracker.host = match[2]

    portInt64, _ := strconv.ParseInt(match[3][1:], 10, 16)
    tracker.port = uint16(portInt64)

    tracker.path = match[4]

    return &tracker, nil
}

func (tracker tracker_impl) String() string {
    return fmt.Sprintf(
        "%s://%s:%d%s",
        tracker.protocol.String(),
        tracker.host,
        tracker.port,
        tracker.path,
    )
}

func (tracker tracker_impl) Do(request src.TrackerRequest) (src.TrackerResponse, error) {
    switch tracker.protocol {
    case ProtocolHTTP:
       return tracker.doHTTP(request)
    case ProtocolHTTPS:
       return tracker.doHTTPS(request)
    case ProtocolUDP:
        return tracker.doUDP(request)
    default:
        return nil, ErrorProtocolUndefined
    }
}

func (tracker tracker_impl) doHTTP(request src.TrackerRequest) (src.TrackerResponse, error) {
    var err error

    httpReq, err := http.NewRequest(
        "GET",
        tracker.String(),
        nil,
    )
    if err != nil {
        return nil, err
    }

    httpReq.URL.RawQuery = requestQuery(request)

    httpRes, err := tracker.httpClient.Do(httpReq)
    if err != nil {
        return nil, err
    }
    defer httpRes.Body.Close()

    body, _ := io.ReadAll(httpRes.Body)

    return decodeResponse(string(body))
}

func (tracker tracker_impl) doHTTPS(request src.TrackerRequest) (src.TrackerResponse, error) {
    return nil, ErrorProtocolNotSuppported
}

func (tracker tracker_impl) doUDP(request src.TrackerRequest) (src.TrackerResponse, error) {
    return nil, ErrorProtocolNotSuppported
}

