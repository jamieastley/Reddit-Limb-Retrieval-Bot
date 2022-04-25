package repository

import "time"

type Clock interface {
	NowUTC() time.Time
}

type clock struct{}

func NewClock() Clock {
	return &clock{}
}

func (c *clock) NowUTC() time.Time {
	return time.Now().UTC()
}
