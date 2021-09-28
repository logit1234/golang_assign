package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Enter The Float Number:")
	var fnum float64
	fmt.Scanln(&fnum)

	var y int = int(fnum)
	fmt.Printf("%d", y)

}
