package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("learning web req")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error while fetching data", err)
		return
	}
	defer res.Body.Close()
	fmt.Printf("Type of response is : %T\n", res)
	fmt.Println("response:", res)

	//read the response body
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Errror while reading:", err)
		return
	}
	fmt.Println("respone:", string(data))
}
