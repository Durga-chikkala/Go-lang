package main

import "fmt"

func main() {
	fmt.Println("hii")
	var fruitList = []string{"apple"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "mango", "banana")
	fruitList = append(fruitList[1:])

	fmt.Println(fruitList)
	c := make([]int, 4)
	c[0] = 1
	c[1] = 2
	c[2] = 3
	c[3] = 4
	c = append(c, 5, 6)
	fmt.Println(c)
	index := 4
	c = append(c[:index], c[index+1:]...)
	fmt.Println(c)
	fmt.Println(len(c))
}
