package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning url")
	myURL := "https://developer.mozilla.org/en-US/docs/Web/CSS/display"
	fmt.Printf("Type of url is : %T\n", myURL)

	parsedURL,err := url.Parse(myURL)
	if err != nil {
		fmt.Println("cant parse url :", err)
		return
	}
	fmt.Printf("Type of url is %T\n", parsedURL)
	fmt.Println("scheme:",parsedURL.Scheme)
	fmt.Println("Host:",parsedURL.Host)
	fmt.Println("Path:",parsedURL.Path)
}
