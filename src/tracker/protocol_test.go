package tracker

import (
	"reflect"
	"testing"
)

func TestNewProtocol(t *testing.T) {
    type Input struct {
        value   string
    }

    type Output struct {
        protocol    Protocol
        err         error
    }

    type TestCase struct {
        input   Input
        output  Output
    }

    testCases := []TestCase{
        { Input{ "http", }, Output{ ProtocolHTTP, nil, }, },
        { Input{ "https", }, Output{ ProtocolHTTPS, nil, }, },
        { Input{ "udp", }, Output{ ProtocolUDP, nil, }, },
        { Input{ "", }, Output{ nil, ErrorProtocolUndefined, }, },
        { Input{ "akdfjas;", }, Output{ nil, ErrorProtocolUndefined, }, },
    }

    for _, testCase := range testCases {
        var output Output
        output.protocol, output.err = NewProtocol(testCase.input.value)
        if reflect.DeepEqual(testCase.output, output) == false {
            t.Errorf("\ngive\t: %v\nexpect\t: %v\ngot\t: %v",testCase.input, testCase.output, output)
        }
    }
}

func TestProtocolToString(t *testing.T) {
    type Input struct {
        protocol Protocol
    }

    type Output struct {
        value   string
    }

    type TestCase struct {
        input   Input
        output  Output
    }

    testCases := []TestCase{
        { Input{ ProtocolHTTP }, Output { "http" }, },
        { Input{ ProtocolHTTPS }, Output { "https" }, },
        { Input{ ProtocolUDP }, Output { "udp" }, },
    }

    for _, testCase := range testCases {
        var output Output
        output.value = testCase.input.protocol.String()

        if reflect.DeepEqual(testCase.output, output) == false {
            t.Errorf("\ngive\t: %v\nexpect\t: %v\ngot\t: %v",testCase.input, testCase.output, output)
        }
    }
}
