package main

import (
	"fmt"
)

func main() {
	fmt.Println("Let's use the go to consume the APIs")
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	responsehttp.Get(myurl)
}