package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Hello welcome")
	content := "This is Durga Dude"
	file, err := os.Create("./details.txt")
	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	defer file.Close()
	byte, err := ioutil.ReadFile("./details.txt")
	fmt.Println(string(byte))

}
