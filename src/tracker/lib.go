package tracker

import (
	"math"
	"time"
)

func UDPTimeout(n uint8) time.Duration {
	return 15 * time.Duration(math.Pow(2, float64(n))) * time.Second
}
