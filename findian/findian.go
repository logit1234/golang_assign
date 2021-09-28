package main

import (
	"fmt"
	"strings"
	
)

func main() {
	fmt.Printf("Enter The String:")
	var enter string
	fmt.Scanln(&enter)

	fmt.Println(strings.ContainsAny(strings.ToLower(enter), "a"))
	fmt.Println(strings.HasPrefix(strings.ToLower(enter), "i"))
	fmt.Println(strings.HasSuffix(strings.ToLower(enter), "n"))
}
