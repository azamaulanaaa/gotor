package tracker

import (
	"reflect"
	"testing"
)

func TestNewEventFromString(t *testing.T) {
    type Input struct {
        value   string
    }

    type Output struct {
        event   Event
        err     error
    }

    type TestCase struct {
        input   Input
        output  Output
    }

    testCases :=[]TestCase{
        { Input{ "started", }, Output{ EventStarted, nil, }, },
        { Input{ "stopped", }, Output{ EventStopped, nil, }, },
        { Input{ "completed", }, Output{ EventCompleted, nil, }, },
        { Input{ "", }, Output{ nil, ErrorEventUndefined, }, },
        { Input{ "akdfjas;", }, Output{ nil, ErrorEventUndefined, }, },
    }

    for _, testCase := range testCases {
        var output Output
        output.event, output.err = NewEvent(testCase.input.value)
        if reflect.DeepEqual(testCase.output, output) == false {
            t.Errorf("\ngive\t: %v\nexpect\t: %v\ngot\t: %v",testCase.input, testCase.output, output)
        }
    }
}
