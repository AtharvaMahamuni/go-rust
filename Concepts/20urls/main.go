package main 

import (
	"fmt"
	"net/url"
)

const myUrl = "https://www.github.com:3000/AtharvaMahamuni?test=1234&test2=5678"
// const myUrl = "https://www.github.com/AtharvaMahamuni"

func main() {
	fmt.Println("Welcome to the urls, module!")
	fmt.Println(myUrl)

	//parsing 
	result, _ := url.Parse(myUrl)

	// fmt.Println(result)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params are: %T\n", qparams)

	for _, val := range qparams{
		fmt.Println("Param is:", val)
	}

	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "github.com",
		Path: "/AtharvaMahamuni",
		RawPath: "test=1234",
	}

	anotherURL := partsOfUrl.String()
	fmt.Println(anotherURL)
}