package gosync

import (
	"testing"
	"time"
	"reflect"
)

func Test_State(t *testing.T) {
	state := NewState()
	time1 := time.Now()
	time2 := time1.Add(time.Duration(2000))
	state.PathToInfo["c:/Program Files/file"] = NewInfo(time1)
	state.PathToInfo["c:/Program Files/file2"] = NewInfo(time2)
	state.Processed["c:/Program Files"] = true

	SaveState("state.json", state)

	loaded := LoadState("state.json")

	if !reflect.DeepEqual(loaded, state) {
		t.Errorf("Wrong PathToInfo loaded, %s != %s", state.PathToInfo, loaded.PathToInfo)
	}
	if len(loaded.PathToInfo) != 2 {
	    t.Errorf("Expected 2 PathToInfo but found %d", len(loaded.PathToInfo))
	}
}
