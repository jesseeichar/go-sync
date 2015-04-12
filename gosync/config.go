package gosync

import (
	"io/ioutil"
	"encoding/json"
)

type Config struct {
	Mappings []Mapping
}

type Mapping struct {
	From string
	To []string
}

func LoadConfig(file string) Config {
	data,err := ioutil.ReadFile(file)

	handleError(err)
	var config Config
	handleError(json.Unmarshal(data, &config))

	return config
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

