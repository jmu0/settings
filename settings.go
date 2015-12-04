package settings

import (
	"errors"
	"io/ioutil"
	"strings"
)

type Settings struct {
	File     string
	Settings map[string]string
	Loaded   bool
}

//load settings from file
func (s *Settings) Load() (map[string]string, error) {
	settings := map[string]string{}
	str, err := readFile(s.File)
	if err != nil {
		return settings, errors.New("Can't read settings file: " + s.File)
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
	if len(settings) == 0 {
		return settings, errors.New("No settings found in: " + s.File)
	}
	s.Settings = settings
	s.Loaded = true
	return settings, nil
}

//get a setting
func (s *Settings) Get(setting string) (string, error) {
	if s.Loaded == false {
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
func (s *Settings) Set(key, value string) error {
	s.Settings[key] = value
	return nil
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
		Loaded: false,
	}
	_, err := s.Load()
	if err != nil {
		return s, err
	}
	return s, nil
}
