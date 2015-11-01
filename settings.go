package settings

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Settings struct {
	File     string
	Settings map[string]string
	loaded   bool
}

//load settings from file
func (s *Settings) Load() (map[string]string, error) {
	settings := map[string]string{}
	str, err := readFile(s.File)
	if err != nil {
		return settings, errors.New("Settings file does not exist: " + s.File)
	}
	lines := strings.Split(str, "\n")
	if len(lines) > 0 {
		for _, line := range lines {
			if len(line) > 0 && line[:1] != "#" {
				fields := strings.Fields(line)
				if len(fields) > 1 {
					settings[fields[0]] = strings.Join(fields[1:], " ")
				}
			}
		}
	}
	s.Settings = settings
	s.loaded = true
	return settings, nil
}

//get a setting
func (s *Settings) Get(setting string) (string, error) {
	if s.loaded == false {
		_, err := s.Load()
		if err != nil {
			return "", err
		}
	}
	if ret, ok := s.Settings[setting]; !ok {
		return ret, errors.New("Setting " + setting + " doesn't exist")
	} else {
		return ret, nil
	}
	return "", errors.New("Setting " + setting + " doesn't exist") //for backwards compatibility (..doesn't end with return statement)
}

//read file into string
func readFile(path string) (string, error) {
	cont, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(cont), nil
}

//return settings
func GetSettings(filename string) (Settings, error) {
	s := Settings{
		File:   filename,
		loaded: false,
	}
	_, err := s.Load()
	if err != nil {
		return s, err
	}
	return s, nil
}
