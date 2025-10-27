package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println("Numbers is :", i)
	}

	counter := 0

	for {
		fmt.Println("Infinite loop")
		counter++
		if counter == 3 {
			break
		}
	}

	//range keyword
	numbers := []int{1, 2, 3, 4, 5}
	for index, value := range numbers {
		fmt.Println("Index of numbers is %d, and value is %d\n", index, value)
	}

	data := "hello world"
	for index, char := range data {
		fmt.Printf("Index of data is %d, and value is %c\n", index, char)
	}
}
