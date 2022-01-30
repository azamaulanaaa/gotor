package tracker

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strconv"
)

var (
    ErrorTrackerInvalid = errors.New("value is not a valid tracker")
    ErrorProtocolNotSuppported = errors.New("protocol not supported yet")
)

type Client interface {
    FromString(string) error
    String() string
    Do(Request) (Response, error)
}

type client_impl struct {
    Protocol    Protocol
    Host        string
    Port        uint16
    Path        string
    HTTPClient  *http.Client
}    

func NewClient(value string) (Client, error) {
    var err error

    client := client_impl{
        HTTPClient:     http.DefaultClient,
    } 

    err = client.FromString(value) 
    if err != nil {
        return nil, err
    }

    return &client, nil
}

func (client client_impl) String() string {
    return fmt.Sprintf(
        "%s://%s:%d%s",
        client.Protocol.String(),
        client.Host,
        client.Port,
        client.Path,
    )
}

func (client *client_impl) FromString(value string) error {
    var err error

    var match []string
    {
        regex_pattern := regexp.MustCompile(`^([\w]+):\/\/([\w\d.-]+)(:[\d]+)?(\/(.*)+)?$`)
        matches := regex_pattern.FindAllStringSubmatch(value, 1)
        if len(matches) != 1 {
            return ErrorTrackerInvalid
        }
        match = matches[0]
    }


    client.Protocol, err = NewProtocol(match[1])
    if err != nil {
        return ErrorTrackerInvalid
    }
    
    client.Host = match[2]

    portInt64, _ := strconv.ParseInt(match[3][1:], 10, 16)
    client.Port = uint16(portInt64)

    client.Path = match[4]

    return nil
}

func (client client_impl) Do(request Request) (Response, error) {
    switch client.Protocol {
    case ProtocolHTTP:
       return client.doHTTP(request)
    case ProtocolHTTPS:
       return client.doHTTPS(request)
    case ProtocolUDP:
        return client.doUDP(request)
    default:
        return Response{}, ErrorProtocolUndefined
    }
}

func (client client_impl) doHTTP(request Request) (Response, error) {
    var err error

    req, err := http.NewRequest(
        "GET",
        client.String(),
        nil,
    )
    if err != nil {
        return Response{}, err
    }

    req.URL.RawQuery = request.String()

    res, err := client.HTTPClient.Do(req)
    if err != nil {
        return Response{}, err
    }
    defer res.Body.Close()

    body, _ := io.ReadAll(res.Body)

    var response = Response{}
    response.FromBencode(string(body))
    
    return response, nil
}

func (client client_impl) doHTTPS(request Request) (Response, error) {
    return Response{ }, ErrorProtocolNotSuppported
}

func (client client_impl) doUDP(request Request) (Response, error) {
    return Response{ }, ErrorProtocolNotSuppported
}
