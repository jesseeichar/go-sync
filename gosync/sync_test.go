package gosync

import (
	"testing"
	"os"
	"path/filepath"
	"io/ioutil"
)

func TestOneOneSync(t *testing.T) {
	os.Remove("test-log.log")
	config := LoadConfig("test/config.json")
	defer config.Close()
	state := State{Processed: map[string]bool{}, PathToInfo: map[string]Info{}}
	s := NewSync(config, state, func() Context {return TestContext{}})
	s.Sync(func(Config, State) bool {
		return false
	})

	var numFiles = 0
	filepath.Walk("test/from", filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
			numFiles += 1
			rel, _ := filepath.Rel("test/from", path)
			for _, toPath := range config.Mappings[0].To {
				to := filepath.Join(toPath, rel)
				toInfo, err := os.Stat(to)
				if os.IsNotExist(err) {
					Fatalf(t, "Expected '%s' to exist but it does not\n", to)
				}
				t.Logf("fromModTime: %v", info.ModTime())
				t.Logf("toModTime: %v", toInfo.ModTime())
				if info.ModTime().Equal(toInfo.ModTime()) {
					Fatalf(t, "ModTime not updated correctly\n")
				}
			}
			return nil
		}));

	if numFiles != 5 {
		Fatalf(t, "Expected 5 files but found %s\n", numFiles)
	}
}

func Fatalf(t *testing.T, pattern string, msg... interface{}) {
	log, _ := ioutil.ReadFile("test-log.log")
	t.Logf("LogData:\n%s\n", log)
	t.Fatalf(pattern, msg);
}

type TestContext struct {}

func (TestContext) ExecuteNext() {
}
func (TestContext) HandleError(err error) {
	panic(err.Error())
}



