package gosync

import (
	"time"
)

type Context interface {
	ExecuteNext()
	HandleError(error)
}

type sleepingContext struct {}

func (sleepingContext) ExecuteNext(action func()) {
	time.Sleep(100 * time.Millisecond)
	action()
}
func (sleepingContext) HandleError(err error) {
	time.Sleep(100 * time.Second)
}
