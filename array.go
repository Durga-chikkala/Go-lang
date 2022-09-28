package main

import "fmt"

func main() {
	//declaration of array
	var a [10]int
	fmt.Println(a)
	b := [5]int{1, 2, 3, 4}
	fmt.Println(b)
	fmt.Println(b[2])
	for i := 0; i < len(b); i++ {
		fmt.Println(b[i])
	}
	fmt.Println(b)
	c := [][]int{{1, 2}, {3, 4}, {5, 6}}
	fmt.Println(c[0][0])
}
