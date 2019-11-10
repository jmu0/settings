package main

import (
	"fmt"

	"github.com/jmu0/settings"
)

type fout struct {
	Fout bool
}

type test struct {
	Een  string `json:"een" yaml:"een"`
	Twee int    `json:"twee" yaml:"twee"`
	Drie bool   `json:"drie" yaml:"drie"`
}

func main() {
	v := test{}
	err := settings.Load("testfile.json", &v)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("struct van json:", v)
	fmt.Println()
	s := map[string]string{}
	err = settings.Load("testfile.yml", &s)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("map van yaml:", s)
}
