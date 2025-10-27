package main

import "fmt"

func main(){
	//name <-> grade
	studentGrades := make(map[string]int)

	studentGrades["Pranav"]= 98
	studentGrades["ok"]=78
	studentGrades["no"]=90

	fmt.Println("Marks of ok:",studentGrades["ok"])
	studentGrades["ok"]= 100
	fmt.Println("new marks of ok :", studentGrades["ok"])

	delete(studentGrades,"ok")
	fmt.Println("Marks of ok :",studentGrades["ok"])

	//checking whether a key exists

	grades, exists := studentGrades["garg"]
	fmt.Println("David's grades:",grades)
	fmt.Println("David exists:",exists)

	for index,value := range studentGrades{
		fmt.Printf("key is %s and marks is %d\n",index,value)
	}

	
}