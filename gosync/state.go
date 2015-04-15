package gosync

import (
	"io/ioutil"
	"encoding/json"
	"os"
	"time"
)

type State struct {
	Processed map[string]bool
	PathToInfo map[string]Info
}
type Info struct {
	ModTimeBytes []byte
}

func NewState() State {
	return State {
		Processed: map[string]bool{},
		PathToInfo: map[string]Info{}}
}

func NewInfo(modTime time.Time) Info {
	bytes, err := modTime.MarshalBinary()
	handleError(err)
	return Info{ModTimeBytes: bytes}
}

func (s *Info) ModTime() time.Time {
	if s.ModTimeBytes == nil {
		return time.Now()
	} else {
		var modTime time.Time
		modTime.UnmarshalBinary(s.ModTimeBytes)

		return modTime
	}
}

func LoadState(file string) State {
	bytes, err := ioutil.ReadFile(file)
	handleError(err)

	var state State
	handleError(json.Unmarshal(bytes, &state))
	return state
}

func SaveState(file string, state State) {
	data, err := json.MarshalIndent(state, "", "  ")
	handleError(err)

	ioutil.WriteFile(file, data, (os.FileMode)(660))
}
