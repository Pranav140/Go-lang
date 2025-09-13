package main

import "fmt"

func main() {
	fmt.Println("we are learning array in golang")

	var name [5]string
	name[0] = "pranav"
	name[2] = "garg"

	fmt.Println("names of person is : ", name)

	var numbers = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Numbers are:", numbers)


	fmt.Println("Length of numbers array is :",len(numbers))

	fmt.Println("value of name at index 2 is :",name[2])

	var price[10]string
	fmt.Println("Price is :",price)
}
