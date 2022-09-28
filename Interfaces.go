package main

import "fmt"

// Interfaces:
//it is type which contains set of method signatures
//syntax
//type interface_name interface{
//    method signature
// }
// all the methods inside the interface should be declared
// first we will create a interface and in main function we create a variable of type interface then we will assign values to it
//Empty interface :
// when interface contains zero methods then it is called empty interface
// empty interface are used by code that handles values of unknown type for example println function
type Inputs struct {
	length int
	width  int
	height int
}
type Myinterface interface {
	area() int
	volume() int
}

func (a Inputs) area() int {
	return a.length * a.width
}
func (a Inputs) volume() int {
	return a.length * a.width * a.height
}
func describe(i interface{}) {
	fmt.Printf("%v - %T\n", i, i)
}
func main() {
	//fmt.Println("Interfaces in Go Language")
	//var a Myinterface
	//a = Inputs{10, 20, 30}
	//fmt.Println(a.area())
	//fmt.Println(a.volume())
	var a interface{}
	describe(a)
	a = 42
	describe(a)
	a = "hello"
	describe(a)
}
