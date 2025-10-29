package main

import (
	"encoding/json"
	"fmt"
)

type Person struct{
	Name string `json:"name"`
	Age int `json:"age"`
	IsAdult bool `json:"is_adult"`
}

func main() {
	fmt.Println("Learning json")
	person := Person{Name:"john",Age:20,IsAdult:true}
	fmt.Println("person is:",person)


	//convert person in json(marshalling)

	jsonData,err := json.Marshal(person)
	if err!= nil {
		fmt.Println("Error marshalling",err)
		return 
	}
	fmt.Println("person data is :",string(jsonData))


	//decoding (unmarshalling)

	var personData Person
	json.Unmarshal(jsonData,&personData)
	if err!=nil{
		fmt.Println("error unmarshalling:",err)
		return
	}
	fmt.Println("person data is :",personData)



}