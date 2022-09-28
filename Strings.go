package main

import (
	"fmt"
	"strings"
)

func main() {
	a := "durga chikkala"
	fmt.Println(strings.ToUpper(a))
	b := "DURGA CHIIKALA"
	fmt.Println(strings.ToLower(b))
	fmt.Println(strings.ToTitle(a))
	fmt.Println(a)
	fmt.Println(strings.Split(a, " "))
	fmt.Println(strings.Replace(a, "a", "k", 2))
	c := "     Durga chikkala         "
	fmt.Println(strings.TrimSpace(c))
	d := "!! DUrga chikkala !!"
	fmt.Println(strings.Trim(d, " !"))

	req := "a"
	m := 0
	for _, v := range d {
		if string(v) == req {
			m = m + 1
		}
	}
	fmt.Println(m)
	fmt.Println(strings.Count("Durga Chikkala", "a"))
	e := "hii man how are you"
	fmt.Println(strings.Contains(e, "hii"))
	for i, v := range a {
		fmt.Println(i, string(v))
	}
}
