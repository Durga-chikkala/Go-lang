package main

import (
	"Project1/Handler"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("hii")
	http.HandleFunc("/Durga", Handler.FileHandler)
	http.ListenAndServe(":8000", nil)

}
