package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Go Modules!")
	fmt.Println("Hello , I am pranav")

	var name string = "Pranav"
	fmt.Println(name)

	var money = 100000
	fmt.Println(money)

	var version = "10.0.01"
	fmt.Println(version)

	person := "Pranavgarg"
	fmt.Println(person)

	// if we initialize any variable with lowercase , we cant export it in any othe file

	var Public = "Public"     // can be exported and used outside
	var private = "Private"        // cant be used outside
	fmt.Println(Public)
	fmt.Println(private)
}
