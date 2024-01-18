package main

import "fmt"

func main() {
	fmt.Println("Welcome to if-else demo!")

	loginCount := 10
	var result string

	if loginCount < 10 { // not allow to move { to next line
		result = "Regular user count"
	} else if loginCount > 10{
		result = "So many users"
	} else {
		result = "Exactly 10"
	}

	fmt.Println(result)

	if 	theNumber := 12; theNumber%2==0 {
		fmt.Println("The number is even")
	} else {
		fmt.Println("The number is odd")
	}

	// if err != nil {

	// }
}