package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs demo!")

	// no inheritance in golang; no super or parent
	atharva := User{"Atharva", "atharvamahamuni@outlook.com", true, 23}
	fmt.Println("User is : ", atharva)
	fmt.Printf("User is : %+v\n", atharva)
	fmt.Printf("User name is %v and the email is %v\n", atharva.Name, atharva.Email)

	test := Test{atharva, "this is data"}
	fmt.Println(test.Details, test.Data)
}

type User struct {
	Name string
	Email string
	Status bool
	Age int
}

type Test struct {
	Details User
	Data string
}