package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("example.txt")
	if err!= nil {
		fmt.Println("Error while creating file")
		return
	}
	defer file.Close()

	content:= "MY NAME IS PRANAV"
	_,errors := io.WriteString(file,content)
	if errors!= nil {
		fmt.Println("error while writing file :",errors)
		return
	}

	fmt.Println("successfully created ")


	//create a buffer to read a file
	buffer :=make([]byte,1024)

	//read the file content into buffer
	for{
		n,err := file.Read(buffer)
		if err!= io.EOF{
			break
		} 
		if err!= nil{
			fmt.Println("Error while reading the file")
			return
		}

		//process the read content
		fmt.Println(string(buffer[:n]))
	}
}
