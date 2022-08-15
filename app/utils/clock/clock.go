package clock

import "time"

type Clock interface {
	Now() time.Time
}

type ClockImpl struct{}

func New() Clock {
	return &ClockImpl{}
}

func (s *ClockImpl) Now() time.Time {
	return time.Now()
}
