package main

import "fmt"

func simpleFunction() {
	fmt.Println("simple function")
}

func multiply(a,b int)(result int){
	result = a*b
	return
}

func add(a, b int) int {
	return a + b
}

func main() {
	fmt.Println("Learning functions in golang")
	simpleFunction()
	ans := add(3, 4)
	mult := multiply(3,4)
	fmt.Println("add of 2 numbers is :", ans)
	fmt.Println("multiply of 2 numbers is :",mult)
}
