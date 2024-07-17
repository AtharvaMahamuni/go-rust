package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Let's use the go to consume the APIs")
	// PerformGetRequest()
	// PerformPostJSONRequest()
	PerformPostFormRequest()
}

func PerformGetRequest() { // In real world handle it by getting url as a parameter
	const myurl = "http://localhost:8000/get"
	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)
	fmt.Println("Content length:", response.ContentLength)

	content, _ := io.ReadAll(response.Body)

	// Directly converting bytes to string
	// fmt.Println(content)
	// fmt.Println(string(content))

	// With String builder
	var responseString strings.Builder
	responseString.Write(content) // Write function returns the bytecount & error
	fmt.Println(responseString.String())
	// fmt.Println("Byte count is")
}

func PerformPostJSONRequest() {
	myUrl := "http://localhost:8000/post"

	// fake json payload
	requestBody := strings.NewReader(`
		{
			"name":"Atharva",
			"country":"India",
			"age":23
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)

	// responseBody
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println("Body:", string(responseBody))

}

func PerformPostFormRequest() {

	myUrl := "http://localhost:8000/postform"

	// form data
	data := url.Values{}
	data.Add("firstname", "Atharva")
	data.Add("lastname", "Mahamuni")
	data.Add("email", "atharva@go.dev")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code:", response.StatusCode)

	// response
	responseBody, _ := io.ReadAll(response.Body)
	fmt.Println(string(responseBody))
}
