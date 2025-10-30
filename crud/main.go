package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userID"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response :", res.Status)
		return
	}

	// data, err := io.ReadAll(res.Body)
	// if err != nil {
	//     fmt.Println("Error reading:", err)
	//     return
	// }
	// fmt.Println("Data:", string(data))

	var todo Todo
	json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding :", err)
		return
	}
	fmt.Println("Todo:", todo)
}

func performPostRequest() {
	todo := Todo{
		UserID:    10,
		Title:     "Pranav",
		Completed: true,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error marshalling :", err)
		return
	}

	jsonString := string(jsonData)

	//convert string data to reader
	jsonReader := strings.NewReader(jsonString)
	myUrl := "https://jsonplaceholder.typicode.com/todos"

	//send post request

	res, err := http.Post(myUrl,"application/json", jsonReader)
	if err!= nil{
		fmt.Println("Error sending request:",err)
		return
	}

	defer res.Body.Close()

	data , _ := io.ReadAll(res.Body)
	fmt.Println("Response :",string(data))

}

func main() {
	fmt.Println("Learning crud..")
	//performGetRequest()
	performPostRequest()
}
