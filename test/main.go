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
	var err error
	// v := test{}
	// err := settings.Load("testfile.json", &v)
	// if err != nil {
	// 	fmt.Println("ERROR json:", err)
	// }
	// fmt.Println("struct van json:", v)
	// fmt.Println()
	// s := map[string]string{}
	// err = settings.Load("testfile.yml", &s)
	// if err != nil {
	// 	fmt.Println("ERROR yml:", err)
	// }
	// fmt.Println("map van yaml:", s)
	var g string
	err = settings.Get("testfile.conf", "een", &g)
	if err != nil {
		fmt.Println("ERROR Get string:", err)
	} else {

		fmt.Println("Test Get string:", g)
	}
	var i int
	err = settings.Get("testfile.conf", "twee", &i)
	if err != nil {
		fmt.Println("ERROR Get:", err)
	} else {

		fmt.Println("Test Get int:", i)
	}
}
