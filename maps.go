package main

import "fmt"

//map is sent as a call by reference

func main() {
	var m map[string]int
	if m == nil {
		fmt.Println(nil)
	}
	//m["a"] = 10               // This is not possible in map since
	//fmt.Println(m)
	k := make(map[string]int)
	if k == nil {
		fmt.Println(nil)
	} else {
		fmt.Println(k)
	}
	k["a"] = 1
	fmt.Println(k)
}
