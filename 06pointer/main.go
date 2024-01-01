package main

import "fmt";

func main() {
	fmt.Println("Welcome, to pointers!");

	// var ptr *int; // Type is `<nil>` if not initialized
	// fmt.Println("Type of ptr is when not initialized: ", ptr);

	number := 22;
	ptr := &number;

	fmt.Println("number is : ", number);
	fmt.Println("&number is : ", &number);
	fmt.Println("ptr is : ", ptr);
	fmt.Println("*ptr is : ", *ptr);
	

	fmt.Println("---- Changing number with pointer ----");

	// operation on pointer
	*ptr = *ptr + 3;
	fmt.Println("*ptr = *ptr + 3 makes number : ", number);
}
