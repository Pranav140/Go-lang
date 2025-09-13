package main

import "fmt"

func main() {
	// numbers := []int{1,2,3,4,5}
	// fmt.Println("Numhrrs:",numbers)
	// fmt.Printf("Number has data type : %T\n",numbers)
	// fmt.Println("Length:", len(numbers))

	// name:= []string{}
	// fmt.Println(name)

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 10) //capacity doubles itself when filled

	fmt.Println("Slice:", numbers)
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))
}
