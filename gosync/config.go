package gosync

import (
	"io/ioutil"
	"encoding/json"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

const (
	OFF = 0
	ERROR = 1
	INFO = 2
	DEBUG = 3)

type Config struct {
	DebugLevel int
	Log string
	logger *log.Logger
	logFile *os.File
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
	if len(config.Log) == 0 {
		if current, err := user.Current(); err != nil {
			handleError(err)
		} else {
			config.Log = filepath.Join(current.HomeDir, ".gosync", "gosync.log")
		}
	}
	handleError(os.MkdirAll(filepath.Dir(config.Log), os.FileMode(0777)))
	logFile, err := os.OpenFile(config.Log, os.O_WRONLY|os.O_CREATE, os.FileMode(0666))
	handleError(err)
	config.logFile = logFile
	config.logger = log.New(config.logFile, "", log.LstdFlags)
	return config
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}

func (c Config) Errorf(pat string, args... interface {}){
	if c.DebugLevel >= ERROR {
		c.logger.Printf("[ERROR] - " + pat, args...)
	}
}
func (c Config) Infof(pat string, args... interface {}){
	if c.DebugLevel >= INFO {
		c.logger.Printf("[INFO] - " + pat, args...)
	}
}
func (c Config) Debugf(pat string, args... interface {}){
	if c.DebugLevel >= DEBUG {
		c.logger.Printf("[DEBUG] - " + pat, args...)
	}
}

func (c Config) Close() {
	c.logFile.Close()
}
