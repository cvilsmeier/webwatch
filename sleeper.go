package main

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
	time.Sleep(duration)
}
