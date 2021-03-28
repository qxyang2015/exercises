package main

import (
	"fmt"
)

func main() {
	fmt.Println("start")
	s := "abcABC123"
	//for idx, v := range s {
	//	fmt.Println(idx, v, s[idx], reflect.TypeOf(v), reflect.TypeOf(s[idx]), string(v))
	//}
	for _, v := range []byte(s) {
		fmt.Println(v)
	}
	fmt.Println("done!")
}
