package main

import "fmt"

func main() {
	var input string = "Durga"
	if input == "" {
		fmt.Printf("one for you,one for me")
	} else {
		fmt.Printf("one for %v one for me", input)
	}

}
