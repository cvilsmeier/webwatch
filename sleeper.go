package webwatch

import (
	"time"
)

type Sleeper interface {
	Sleep(duration time.Duration)
}

func NewSleeper() Sleeper {
	return sleeperImpl{}
}

type sleeperImpl struct {
}

func (this sleeperImpl) Sleep(duration time.Duration) {
	DEBUG.Printf("will sleep %s\n", duration)
	time.Sleep(duration)
}
