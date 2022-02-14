package tracker

func (event Event) String() string {
	dictionary := map[Event]string{
		EventNone:      "",
		EventCompleted: "completed",
		EventStarted:   "started",
		EventStopped:   "stopped",
	}

	return dictionary[event]
}
