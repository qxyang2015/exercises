package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	cnt := 10
	for mask := 0; mask < 1<<cnt; mask++ {
		for i := 0; i < cnt; i++ {
			fmt.Printf("%b >> %v  = %b\n", mask, i, mask>>i)
		}
	}
	fmt.Println("done!")
}
