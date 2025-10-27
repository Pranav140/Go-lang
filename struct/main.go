package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

type Contact struct{
	Email string
	Phone int 
}

type Address struct{
	House int
	Area string 
	State string 
}

type Employee struct{
	Person_details Person 
	Person_contact Contact 
	Person_adress Address
}

func main() {
	var pranav Person
	fmt.Println("Pranav person :", pranav)
	pranav.FirstName = "Pranav"
	pranav.LastName = "garg"
	pranav.Age = 24
	fmt.Println("Prince Person : ", pranav)


	//2nd method
	person1:= Person{
		FirstName:"OK",
		LastName:"OK",
		Age:20,
	}
	fmt.Println("person 1:",person1)

	//new keyword 
	var person2 = new(Person)
	person2.FirstName = "Ok"
	person2.LastName = "bye"
	person2.Age=21

	fmt.Println("Perosn2 : ",person2)

	var employee1 Employee
	employee1.Person_details = Person{
		FirstName : "pranav",
		LastName :"garg",
		Age:20,
	}
	employee1.Person_contact=Contact{
		Email:"abc@gmail.com",
		Phone:6281910010,
	}
	employee1.Person_adress= Address{
		House : 1,
		Area :"premnagar",
		State:"punjab",
	}

	fmt.Println("Employee1:",employee1.Person_adress.State)
}
