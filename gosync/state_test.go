package gosync

import (
	"testing"
	"time"
	"reflect"
)

func Test_State(t *testing.T) {
	state := State{PathToInfo: map[string]Info{}}
	time1 := time.Now()
	time2 := time1.Add(time.Duration(2000))
	state.PathToInfo["c:/Program Files/file"] = NewInfo(time1)
	state.PathToInfo["c:/Program Files/file2"] = NewInfo(time2)

	SaveState("state.json", state)

	loaded := LoadState("state.json")

	if reflect.DeepEqual(loaded, state) {
		t.Errorf("Wrong state loaded, %s != %s", state, loaded)
	}
}
