package main

import (
	"errors"
	"fmt"
	"io"
)

type s struct {
	a int
	b int
}

func (t s) Read(p []byte) (i int, err error) {
	fmt.Println("hii")
	if t.a == 0 {
		fmt.Println("by")
		return 1, nil
	}
	return 0, errors.New("err")
}

func main() {
	a := s{0, 2}
	k, err := io.ReadAll(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(k)
	//http.HandleFunc("/", Handlerrr)
	//http.ListenAndServe(":7000", nil)
}
