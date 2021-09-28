package main

import (
	"fmt"
	"sort"
	"os"
	"bufio"
)
func main() {
	intSort := make([]int, 3)

	i:=0
	for i < len(intSort){
		fmt.Printf("Enter the Intergers:")
		
		//var num int
		//fmt.Scan(&num)
		 
		inputReader := bufio.NewReader(os.Stdin)
		num, _, _ := inputReader.ReadRune()

		if num == 'X' {
			
			fmt.Printf("Thanks")
			break
		}   		
        intSort[i] = int(num-48)	
		i++
	}
	sort.Ints(intSort)
	fmt.Printf("%d",intSort)
	
}