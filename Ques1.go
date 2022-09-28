package main

import (
	"fmt"
	"strings"
)

func cap(input string) bool {
	if input == strings.ToUpper(input) {
		return true
	} else {
		return false
	}

}
func check(input string) {
	switch {
	case cap(input) && string(input[len(input)-1]) == "?":
		fmt.Println("Calm down, I know what I'm doing!")
	case cap(input):
		fmt.Println("Whoa, chill out!")
	case input == " ":
		fmt.Println("Fine. Be that way!")
	case string(input[len(input)-1]) == "?":
		fmt.Println("Sure")
	default:
		fmt.Println("Whatever.")
	}
}
func main() {
	var input string = " "
	check(input)
}
