package main

import "fmt"

func main() {

	fmt.Println("====================================")

	var username string = "Atharva" //If we not put anything, go will autoassign this with ""
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n", username)

	fmt.Println("====================================")

	var isLoggedIn bool = true //false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn)

	fmt.Println("====================================")

	var smallInt uint8 = 255 //256 - gives error read docs to learn the bounds for each type
	fmt.Println(smallInt)
	fmt.Printf("Variable is of type : %T \n", smallInt)

	fmt.Println("====================================")

	var smallFloat float32 = 25.36237864366435 // precision is lost to 5 values after point.
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n", smallFloat)

	fmt.Println("====================================")

	var anotherVariable int = 2555555
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type : %T \n", anotherVariable)

	fmt.Println("====================================")

	// implicit type
	var webSite = "www.github.com/atharvamahamuni/"
	fmt.Println(webSite)
	fmt.Printf("Variable is of Type: %T \n", webSite)

	fmt.Println("====================================")

	// no var style
	numberOfUsers := 837782
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of Type: %T \n", numberOfUsers)

	fmt.Println("====================================")
}
