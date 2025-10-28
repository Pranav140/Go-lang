package main

import "fmt"

func add(a,b int) int {
	return a+b
}

//jiske aage differ lga hai vo last mei chla jata hai 
//agar multiple defer lge hue hai to stack mei chla jata hai

func main() {
	fmt.Println("Starting of the program")
	data := add(5,6)
	fmt.Println("add:",data)
	defer fmt.Println("middle of the program")
	fmt.Println("end of the program")
}

