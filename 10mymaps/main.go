package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to map code!")

	languages := make(map[string]string)
	languages["js"] = "JavaScript"
	languages["py"] = "Python"
	languages["go"] = "Golang"

	fmt.Println("Languages map is: ", languages)
	fmt.Println("go is nothing but: ", languages["go"])

	delete(languages, "py")
	fmt.Println(languages)

	// Loop introduction

	for key, value := range languages {
		fmt.Printf("The key %v, is for the %v\n", key, value)
	}
}