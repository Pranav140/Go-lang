package main

import "fmt"

func modifyValueByReference(num *int) {
	*num = *num * 2
}

func main() {
	value := 10
	modifyValueByReference(&value)
	fmt.Println("Modified value:", value)
}
