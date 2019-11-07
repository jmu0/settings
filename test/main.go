package main

import (
	"fmt"

	"github.com/jmu0/settings"
)

type fout struct {
	Fout bool
}

type test struct {
	Een  string `json:"een"`
	Twee int    `json:"twee"`
	Drie bool   `json:"drie"`
}

func main() {
	v := test{}
	err := settings.Load("testfile.conf", &v)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("struct:", v)
	fmt.Println()
	s := map[string]string{}
	err = settings.Load("testfile.conf", &s)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	fmt.Println("map:", s)
}
