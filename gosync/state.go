package gosync

import (
	"time"
	"io/ioutil"
	"encoding/json"
	"os"
)

type State struct {
	PathToInfo map[string]Info
}
type Info struct {
	LastModified []byte
}

func NewInfo(lastModified time.Time) Info {
	bytes, err := lastModified.MarshalBinary()
	handleError(err)
	return Info{LastModified: bytes}
}

func (s *Info) LastModifiedTime() time.Time {
	var lastModified time.Time
	lastModified.UnmarshalBinary(s.LastModified)

	return lastModified
}

func LoadState(file string) State {
	bytes, err := ioutil.ReadFile(file)
	handleError(err)

	var state map[string]Info
	handleError(json.Unmarshal(bytes, &state))
	return State{PathToInfo: state}
}

func SaveState(file string, state State) {
	data, err := json.MarshalIndent(state, "", "  ")
	handleError(err)

	ioutil.WriteFile(file, data, (os.FileMode)(660))
}
