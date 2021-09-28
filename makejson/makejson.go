package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	mapj := make(map[string]string)

	for i := 0; i <2; i++ {
		arr := [2]string{"Name :","Address :"}
		fmt.Printf(arr[i])
		var name string
		fmt.Scan(&name)
		mapj[arr[i]] = name
	}
	jmap, _ := json.Marshal(mapj)
	fmt.Printf(string(jmap))
}
