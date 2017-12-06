package webwatch

import (
	"time"
)

// A Sleeper sleeps.
type Sleeper interface {

	// Sleep sleeps the specified duration.
	Sleep(duration time.Duration)
}

// NewSleeper creates a new Sleeper.
func NewSleeper() Sleeper {
	return sleeperImpl{}
}

type sleeperImpl struct {
}

func (si sleeperImpl) Sleep(duration time.Duration) {
	DEBUG.Printf("will sleep %s\n", duration)
	time.Sleep(duration)
}
