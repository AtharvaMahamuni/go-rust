package main

import (
	"fmt"
	"bufio"
	"os"
)

func main() {
	fmt.Println("Welcome, to user input program!");

	// creating reader
	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Please enter the rating for our Pizza: ");


	// {comma ok || comma err} syntax
	input, _ := reader.ReadString('\n');

	fmt.Print("Thanks!, for rating ", input);
	fmt.Printf("Type of rating given by you: %T \n", input);
}
