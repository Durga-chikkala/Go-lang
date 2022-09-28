package main

import "fmt"

func main() {
	type address struct {
		name    string
		city    string
		pincode int
	}

	a := address{"Durga", "hsr layout", 533401}
	fmt.Println(a.name)
	fmt.Println(a.pincode)

}
