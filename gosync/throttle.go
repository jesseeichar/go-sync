package gosync

import "time"

type Throttle interface {
	ExecuteNext(func())
}
type ThrottleFactory struct {}

func (ThrottleFactory) Create() Throttle {
	return sleepingThrottle{}
}

type sleepingThrottle struct {}

func (sleepingThrottle) ExecuteNext(action func()) {
	time.Sleep(100 * time.Millisecond)
	action()
}
