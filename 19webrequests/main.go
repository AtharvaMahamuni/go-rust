package main

import (
	"fmt"
	"net/http"
	"io"
)

const url string = "https://atharvamahamuni.github.io/"
// const url string = "https://api.chucknorris.io/jokes/random"

func main() {
	fmt.Println("Let's work on the webrequests!")

	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	// Response structure - https://pkg.go.dev/net/http#Response
	fmt.Printf("Response is of type:%T\n", response)
	fmt.Println("Response status is: ", response.Status)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}

	content := string(databytes)

	fmt.Println(content)
}