package gosync

import "testing"

func Test_Config(t *testing.T) {
	LoadConfig("test/config.json")

	// good enough
}
