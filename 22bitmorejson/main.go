package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func encodingJSON() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "lco", "ase123", []string{"webdev", "react"}},
		{"Flutter Bootcamp", 299, "lco", "asd333", []string{"Mobile", "dart"}},
		{"HTML&CSS Bootcamp", 99, "lco", "sdfa324", []string{"webdev", "dev"}},
		{"Android Bootcamp", 999, "lco", "bcvvb2", nil},
	}

	// Pacakge this data as a JSON data
	// finalJson, err := json.Marshal(lcoCourses)

	// To show good formatted json object we can use.
	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("Final JSON: %s\n", finalJson)
	// fmt.Println("Final JSON:", finalJson)
}

func decodingJSON() {
	webJsonResponse := (`
	{
		"coursename": "ReactJS Bootcamp",
		"Price": 299,
		"website": "lco",
		"password": "abc@123",
		"tags": [
				"webdev",
				"react"
		]
	}
	`)

	var lcoCourse course

	checkValid := json.Valid([]byte(webJsonResponse))

	if checkValid {
		fmt.Println("JSON is valid")
		// don't want to pass copy but want to pass the reference
		json.Unmarshal([]byte(webJsonResponse), &lcoCourse)

		// normal println
		fmt.Println(lcoCourse)
		fmt.Printf("%v\n", lcoCourse)
		fmt.Printf("%#v\n", lcoCourse)

		// some cases where you want to add data as key, value
		var myOnlineData map[string]interface{}
		json.Unmarshal([]byte(webJsonResponse), &myOnlineData)
		fmt.Printf("JSON DATA: %v\n", myOnlineData)

		for k, v := range myOnlineData {
			fmt.Printf("Key: %s value: %v and Type is %T\n", k, v, v)
		}

	} else {
		fmt.Println("JSON is not valid")
	}
}

func main() {
	fmt.Println("Will create some JSON.")
	// encodingJSON()
	decodingJSON()
}
