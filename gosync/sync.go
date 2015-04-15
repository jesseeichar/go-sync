package gosync

import (
	"path/filepath"
	"os"
	"io"
)

type Sync struct {
	config Config
	state  State
	contextFactory func() Context
}

func NewSync(config Config, state State, contextFactory func() Context) Sync {
	return Sync{config:config, state:state, contextFactory:contextFactory}
}

func (s Sync) Sync(continueFunc func(Config, State) bool) {

	s.config.Infof("Starting Sync\n")
	process := []Mapping{}
	for _, mapping := range s.config.Mappings {
		if !s.state.Processed[mapping.From] {
			process = append(process, mapping)
		}
	}

	s.multiSync(process)

	for continueFunc(s.config, s.state) {
		s.state.Processed = map[string]bool{}
		s.multiSync(s.config.Mappings)
	}
}

func (s Sync) multiSync(process []Mapping) {
	s.config.Infof("Syncing: %s\n", process)
	for _, mapping := range process {
		s.singleSync(mapping)
	}
	s.config.Infof("Finished Syncing: %s\n", process)
}

func (s Sync) singleSync(process Mapping) {
	s.config.Debugf("Single Sync %v", process)
	filepath.Walk(process.From, filepath.WalkFunc(func(path string, info os.FileInfo, err error) error {
			s.config.Debugf("Start single file copy '%v'", path)
			if err != nil {
				s.config.Errorf("Error walking file tree: '%s'. path: '%s'", process.From, path)
				s.contextFactory().HandleError(err)
				return nil
			}
			rel, err := filepath.Rel(process.From, path)
			if err != nil {
				s.config.Errorf("Error making relative path in file tree: '%s'. path: '%s'", process.From, path)
				s.contextFactory().HandleError(err)
			} else {
				for _, to := range process.To {
					fromFile, err := os.Open(path)
					if err == nil {
						toPath := filepath.Join(to, rel)
						err := os.MkdirAll(filepath.Dir(toPath), os.FileMode(0700))
						if err != nil {
							s.config.Errorf("Error making parent directories of to file: %v",  err);
							s.contextFactory().HandleError(err)
							return nil;
						}

						toInfo := s.state.PathToInfo[path]
						notExists := os.IsNotExist(err)
						upToDate := toInfo.ModTime().Equal(toInfo.ModTime()) || toInfo.ModTime().After(toInfo.ModTime())
						if (notExists || !upToDate) {
							toFile, err := os.Create(toPath)
							if err != nil {
								s.config.Errorf("Error creating to file: %q, error: %v", toPath, err)
							}
							s.config.Debugf("Copying file '%v' to '%v'", path, toPath)
							_, err = io.Copy(fromFile, toFile)
							s.state.PathToInfo[path] = toInfo
							s.config.Debugf("Done copying file '%v' to '%v'", path, toFile)
						} else {
							s.config.Debugf("Not Copying file as it is up-to-date: '%v' to '%v'\n\t%v is after %v", path, toPath,
								toInfo.ModTime(), info.ModTime())
						}
					}
					if err != nil {
						s.config.Errorf("Error copying a file, path '%s', to '%s', error: '%v'", path, filepath.Join(to, rel), err)
						s.contextFactory().HandleError(err)
					}
				}
			}
			return nil
		}))
	s.state.Processed[process.From] = true
}
