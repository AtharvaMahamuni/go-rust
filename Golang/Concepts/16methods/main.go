package main

import "fmt"

// When functions go inside the structs, we call them Methods!

func main() {
	fmt.Println("Welcome to use methods!")

	atharva := User{"Atharva", "atharvamahamuni@outlook.com", true, 2}

	fmt.Println(atharva)
	atharva.GetStatus()
	atharva.NewMail() // Will not make any change in object.
	fmt.Println(atharva)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("User status is:", u.Status)
}

// User passed here is not the actual reference and just the copy
// so it won't change anything as such in actual object
func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("New user email is:", u.Email)
}
