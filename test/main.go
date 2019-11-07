package main

import (
	"fmt"

	"github.com/jmu0/settings"
)

type fout struct {
	Fout bool
}

type test struct {
	Een  string      `json:"een"`
	Twee int         `json:"twee"`
	Drie interface{} `json:"drie"`
}

func main() {
	v := test{
		Een: "test",
	}
	err := settings.Load("testfile.conf", &v)
	//json.Unmarshal([]byte("{\"een\":\"eerste\"}"), &v)
	fmt.Println("struct:", v)
	// fmt.Println()
	// s := map[string]string{"twee": "3"}
	// err := settings.Load("testfile.conf", &s)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
	// fmt.Println("map:", s)
}
