package Project1

import (
	"fmt"
)

// Function :Function:A function is a block of code which will be executed only when we call it
// it enables code reusability
/*  syntax is
  func function_name(){
  // statements
}
functions can be different types basing on the parameters and return value
it can have parameters, and it can return value
The return type is declared beside the parenthesis
syntax is
func function_name(a,b int) int {
    return a
}
here int is the return type of the function
we can have multiple return statements and also named return values

*/
func Function(a int) int {
	return a
}

//with single return value
func sum(a, b int) int {
	return a + b
}

//multiple return  values
func swap(a, b int) (int, int) {
	temp := b
	b = a
	a = temp
	return a, b
}

//variadic functions
// are type of functions in which the parameters of the function are not fixed they can be changed of any length
// we can also pass array or slice to the function, but we need to  use ... for it
func add(num ...int) int {
	sum := 0
	for i := 0; i < len(num); i++ {
		sum = sum + num[i]
	}
	return sum

}

//passing array to the function
//array will be sent as pass by value not as pass by reference in the functions
// so any changes inside the function will not affect the original array
func changes(nums []int) {
	nums[0] = nums[0] + 100
	nums[1] = nums[1] + 200
	nums = append(nums, 5)
	nums[0] = 2
	fmt.Println(cap(nums))

}

// recursion
// function calling itself in its function definition is called as recursion
// it requires 2 conditions 1.Base case 2.Recursive case
// base case is used to terminate the function call and without base case recursion leads to stack overflow error
//Program to find fibonacci series
func fibo(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

//Closures:
//
func addition(a []int) int {
	summ := 0
	for i := 0; i < len(a); i++ {
		summ = summ + a[i]
	}
	return summ
}

// sending function as a parameter
func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

//call by value and call by reference
func Swap(a, b *int) {
	temp := *b
	*b = *a
	*a = temp
	fmt.Println(*a, *b)

}
func main() {
	//fmt.Println("hello This is about functions")
	//fmt.Println(sum(3, 4))                                 // function with parameter and return value
	//fmt.Println(swap(5, 10))                               // function with 2 return values
	//fmt.Println(add(1, 2, 3, 4))                           // variadic function call
	var nums = []int{1, 2, 3, 4}
	//fmt.Println(add(nums...))                              // variadic function call with array as parameter
	//fmt.Println(addition(nums))
	fmt.Println(nums)
	changes(nums)
	fmt.Println(nums)
	//fmt.Println(compute(math.Pow))
	// we can create anonymous function(functions with no name)
	//hypot := func(a float64, b float64) float64 {
	//	return math.Sqrt((a * a) + (b * b))
	//}
	//	fmt.Println(compute(hypot))                            // sending function as parameter
	// fmt.Println(fibo(2))                                    // recursion
	//a := 10
	//b := 20
	//swap(&a, &b)
	//fmt.Println(a, b)

}
