package tracker

import (
	"errors"
)

var (
    ErrorEventUndefined = errors.New("event is not defined")
)

type Event interface {
    String() string
}

type event_impl uint

const (
    EventStarted event_impl = iota
    EventStopped
    EventCompleted
)

var event_list = map[event_impl]string{
    EventStarted: "started",
    EventStopped: "stopped", 
    EventCompleted: "completed",
}

func NewEvent(value string) (Event, error) {
    for event, eventString := range event_list {
        if value == eventString {
            return event, nil
        }
    }

    return nil, ErrorEventUndefined
}

func (event event_impl) String() string {
    eventString, ok := event_list[event]
    if ok {
        return eventString
    }

    return ""
}
