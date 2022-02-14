package tracker

import "strings"

var eventDictionary = map[Event]string{
    EventNone:      "",
    EventCompleted: "completed",
    EventStarted:   "started",
    EventStopped:   "stopped",
}

func NewEvent(value string) (Event, error) {
    value = strings.ToLower(value)

    for event, eventString := range eventDictionary {
        if eventString == value {
            return event, nil
        }
    }

    return EventNone, ErrorInvalidEvent 
}

func (event Event) String() string {
	return eventDictionary[event]
}
