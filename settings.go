package settings

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strings"
)

//Load settings from file
func Load(file string, s interface{}) error {
	var err error
	if reflect.ValueOf(s).Type().Kind() != reflect.Ptr {
		return errors.New("Target is not a pointer")
	} else if s == nil {
		return errors.New("Target is nil")
	}
	testReflect(s)

	if file != "" {
		_, err = os.Stat(file)
		if !os.IsNotExist(err) {
			switch filepath.Ext(file) {
			case ".conf":
				err = loadConf(file, s)
				if err != nil {
					return err
				}
			case ".json":
				err = loadJSON(file, s)
				if err != nil {
					return err
				}
			case ".yml":
				err = loadYaml(file, s)
				if err != nil {
					return err
				}
			}
		}
	}
	err = loadEnvironmentVariables(s)
	if err != nil {
		return err
	}
	err = loadCommandLineArgs(s)
	if err != nil {
		return err
	}
	return nil
}

func loadJSON(file string, settings interface{}) error {
	return nil
}
func loadYaml(file string, settings interface{}) error {
	return nil
}

func loadConf(file string, settings interface{}) error {
	set := map[string]string{}
	str, err := readFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(str, "\n")
	if len(lines) > 0 {
		for _, line := range lines {
			if len(line) > 0 && line[:1] != "#" {
				fields := strings.Fields(line)
				if len(fields) > 1 {
					set[fields[0]] = strings.Join(fields[1:], " ")
				}
			}
		}
	}
	if len(set) == 0 {
		return errors.New("No settings found in: " + file)
	}
	fmt.Println(set)
	settings = &set
	return nil
}

func loadEnvironmentVariables(s interface{}) error {
	return nil
}

func loadCommandLineArgs(s interface{}) error {
	return nil
}

func testReflect(s interface{}) error {
	switch reflect.ValueOf(s).Elem().Type().Kind() {
	case reflect.Struct:
		fmt.Println("I am a struct")
		str := reflect.ValueOf(s).Elem()
		fld := str.FieldByName("Een")
		if fld.IsValid() {
			fld.SetString("Aangepast!!")
		}
		var i int
		for i = 0; i < str.NumField(); i++ {
			fmt.Println("Field:", str.Type().Field(i).Name, "=", str.FieldByName(str.Type().Field(i).Name).String())
		}
	case reflect.Map:
		fmt.Println("I am a Map")
		// v := reflect.ValueOf(s)
		t := reflect.TypeOf(s).Elem().Key()
		v := reflect.TypeOf(s).Elem().Elem()

		fmt.Println("map key type:", t, "value type:", v)
	default:
		return errors.New("Invalid target")
	}
	return nil
}

//Get get a setting
// func (s *Settings) Get(setting string) (string, error) {
// 	if s.Loaded == false {
// 		err := s.Load()
// 		if err != nil {
// 			return "", err
// 		}
// 	}
// 	ret, ok := s.Settings[setting]
// 	if !ok {
// 		return ret, errors.New("Setting " + setting + " doesn't exist")
// 	}
// 	return ret, nil
// 	//return "", errors.New("Setting " + setting + " doesn't exist") //for backwards compatibility (..doesn't end with return statement)
// }

//Set a setting
// func (s *Settings) Set(key, value string) error {
// 	s.Settings[key] = value
// 	return nil
// }

//readFile read file into string
func readFile(path string) (string, error) {
	cont, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(cont), nil
}

//GetSettings return settings
// func GetSettings(filename string) (interface{}, error) {
// 	s := Settings{
// 		File:   filename,
// 		Loaded: false,
// 	}
// 	err := s.Load()
// 	if err != nil {
// 		return s, err
// 	}
// 	return s, nil
// }
