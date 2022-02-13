package tracker

import (
	"regexp"
)

func NewTracker(host string) (Tracker, error) {
	var match []string
	{
		regex_pattern := regexp.MustCompile(`^([\w]+):\/\/([\w\d.-]+)(:[\d]+)?(\/(.*)+)?$`)
		matches := regex_pattern.FindAllStringSubmatch(host, 1)
		if len(matches) != 1 {
			return nil, ErrorTrackerInvalid
		}
		match = matches[0]
	}
	protocol := match[1]

	switch protocol {
	case "http", "https":
		return newHttpTracker(host), nil
	default:
		return nil, ErrorProtocolNotSuppported
	}
}
