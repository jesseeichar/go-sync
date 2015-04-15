package gosync

import "testing"

func Test_Config(t *testing.T) {
	config := LoadConfig("test/config.json")
	defer config.Close()

	config = LoadConfig("test/config1.json")
	defer config.Close()
	if len(config.Mappings) != 2 {
		t.Errorf("Wrong number of mappings %d", len(config.Mappings))
	}
	if len(config.Mappings[0].To) != 2 {
		t.Errorf("Wrong number of To %d", len(config.Mappings))
	}
	// good enough
}
