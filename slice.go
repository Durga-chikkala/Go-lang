package main

import "fmt"

// Slice is more powerful,flexible,convenient than array
// In array we fix the size of the array whereas in slice we don't fix the size of the array
// Slice contains three things 1.Pointer 2.length 3.Capacity
// the zero value of slice is nil
func sprint(slice []int) {
	fmt.Println("slice is=", slice, "length=", len(slice), "capacity=", cap(slice))
	fmt.Printf("%p\n", &slice)

}

func change(b []int) {
	b[3] = 40
	fmt.Println(b)
}

func main() {
	//var slice []int
	//fmt.Println(slice)
	//if slice == nil {
	//	fmt.Println("nill")
	//}
	//slice = append(slice, 1)
	//sprint(slice)
	//slice = append(slice, 2)
	//slice = append(slice, 3)
	//sprint(slice)
	//b := slice[1:]
	//sprint(b)
	//b = append(b, 4)
	//sprint(b)
	//sprint(slice)
	//x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//fmt.Println(x)
	//s := x[0:3]
	//fmt.Println(s)
	//s = s[3:]
	//
	//sprint(s)
	//s = s[:2]5
	//sprint(s)
	//s = s[:5]
	//sprint(s)
	//a := 10
	//b := 20
	//swap(&a, &b)
	//fmt.Println(a, b)
	//s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//s = s[0:2]
	//sprint(s)
	//s = s[2:]
	//sprint(s)
	//s = s[3:2]
	//sprint(s)
	s := make([]int, 5)
	s[0] = 10
	s[1] = 20
	change(s)
	fmt.Println(s)

}
