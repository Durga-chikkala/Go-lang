package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	response, err := http.Get("https://durgachikkala-programmer.github.io/study/")
	if err != nil {
		panic(err)
	}
	fmt.Printf("%T", response)
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))

}
