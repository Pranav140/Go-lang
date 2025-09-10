package main

import "fmt"

func main() {
	age := 25
	name := "Alice"
	height := 5.8234567

	fmt.Println("age:" ,age, "name:",name,"height:",height)

	fmt.Printf("age is %d\n",age)
	fmt.Printf("height is %f\n",height)
	fmt.Printf("name is %s\n",name)
	fmt.Printf("Type of name is %T\n",name)

	//multiple specifier in single line
	fmt.Printf("name:%s , age:%d , height:%f",name,age,height)
}
