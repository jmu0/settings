package main

import (
	"fmt"

	"github.com/jmu0/settings"
)

type test struct {
	Een string `json:"een"`
}

func main() {
	v := test{
		Een: "test",
	}
	settings.Load("testfile.conf", &v)
	//json.Unmarshal([]byte("{\"een\":\"eerste\"}"), &v)
	fmt.Println("struct:", v)
	// var str string
	// fmt.Println(settings.Load("testfile.conf", &str))
	fmt.Println()
	s := map[string]string{"twee": "drie"}
	settings.Load("testfile.conf", &s)
	fmt.Println("map:", s)
}
