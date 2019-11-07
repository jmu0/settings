package settings

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
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

func loadConf(file string, s interface{}) error {
	var err error
	var str string
	str, err = readFile(file)
	if err != nil {
		return err
	}
	lines := strings.Split(str, "\n")
	if len(lines) > 0 {
		for _, line := range lines {
			if len(line) > 0 && line[:1] != "#" {
				fields := strings.Fields(line)
				if len(fields) > 1 {
					err = set(fields[0], strings.Join(fields[1:], " "), s)
					if err != nil {
						return err
					}
				}
			}
		}
	} else {
		return errors.New("No settings found in: " + file)
	}
	return nil
}

func loadEnvironmentVariables(s interface{}) error {
	var prg, key, value string
	var spl []string
	var err error
	prg = os.Args[0]
	prg = strings.Split(prg, "/")[len(strings.Split(prg, "/"))-1]
	for _, v := range os.Environ() {
		// fmt.Println(v[:len(prg)])
		if len(v) > len(prg) && v[:len(prg)] == strings.ToUpper(prg) {
			spl = strings.Split(v, "=")
			if len(spl) > 1 {
				key = strings.ToLower(strings.Replace(spl[0], strings.ToUpper(prg)+"_", "", 1))
				value = strings.Join(spl[1:], "=")
				err = set(key, value, s)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func loadCommandLineArgs(s interface{}) error {
	//TODO: load args
	var spl []string
	var key, value string
	var err error
	for _, e := range os.Args[1:] {
		if e[:2] == "--" {
			e = e[2:]
			spl = strings.Split(e, "=")
			if len(spl) > 1 {
				key = strings.ToLower(spl[0])
				value = strings.Join(spl[1:], "=")
				err = set(key, value, s)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func set(key, value string, s interface{}) error {
	//TODO: handle more types
	switch reflect.ValueOf(s).Elem().Type().Kind() {
	case reflect.Struct:
		st := reflect.TypeOf(s).Elem()
		var i int
		for i = 0; i < st.NumField(); i++ {
			if st.Field(i).Name == key || st.Field(i).Tag.Get("json") == key {
				fld := reflect.ValueOf(s).Elem().FieldByName(st.Field(i).Name)
				if fld.IsValid() {
					if fld.Kind() == reflect.String {
						fld.SetString(value)
					} else if fld.Kind() == reflect.Interface {
						fld.Set(reflect.ValueOf(value))
					} else if fld.Kind() == reflect.Int {
						intVal, err := strconv.ParseInt(value, 10, 64)
						if err != nil {
							return errors.New("Invalid int value: " + value)
						}
						fld.SetInt(intVal)
					} else if fld.Kind() == reflect.Bool {
						boolVal, err := strconv.ParseBool(value)
						if err != nil {
							return errors.New("Invalid bool value: " + value)
						}
						fld.SetBool(boolVal)
					} else {
						return errors.New("Target struct has invalid field type: " + fld.Kind().String())
					}
					return nil
				}
			}
		}
		return nil
	case reflect.Map:
		keyType := reflect.TypeOf(s).Elem().Key()
		valueType := reflect.TypeOf(s).Elem().Elem()
		if keyType.Kind() == reflect.Int {
			if valueType.Kind() == reflect.String || valueType.Kind() == reflect.Interface {
				intKey, err := strconv.Atoi(key)
				if err != nil {
					intKey = 0
				}
				reflect.ValueOf(s).Elem().SetMapIndex(reflect.ValueOf(intKey), reflect.ValueOf(value))
			} else {
				return errors.New("Invalid value type")
			}
		} else if keyType.Kind() == reflect.String {
			if valueType.Kind() == reflect.String || valueType.Kind() == reflect.Interface {
				reflect.ValueOf(s).Elem().SetMapIndex(reflect.ValueOf(key), reflect.ValueOf(value))
			} else {
				return errors.New("Invalid value type")
			}
		} else {
			return errors.New("Invalid key type")
		}
		return nil
	}
	return errors.New("Invalid target")
}

//Get gets setting from file/env/args
func Get(filename, key string, s interface{}) error {
	//TODO: load settings and get key
	return nil
}

//readFile read file into string
func readFile(path string) (string, error) {
	cont, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(cont), nil
}
