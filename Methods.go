package main

import (
	"fmt"
	"time"
)

// struct will send the copy so changes made inside the function will not efffect the original struct
type Author struct {
	Name string
	roll int
	Book string
	year int
}

func (a Author) change() {
	a.roll = 590
	fmt.Println(a.roll)
}

// for this reason we use pointers inorder make it call by reference
func (b *Author) pchanges(c int) {
	(*b).year = c
}

// sending non-struct type as receiver type
type float float64

func main() {

	res := Author{"Durga", 580, "The Loyal Journey", 2022}
	fmt.Println(res)
	res.change()
	fmt.Println(res)
	j := &res
	j.pchanges(2011)
	fmt.Println(res)
	now := time.Now()
	fmt.Println(int(now.Month()))

	fmt.Println(now)
}
