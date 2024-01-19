package main 

import "fmt"

func sayHi() {
	fmt.Println("Hey, Welcome!")
}

func adder(valOne int, valTwo int) int{
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum, "Pro function"
}

func main() {
	fmt.Println("Welcome to main function!")
	sayHi()

	result := adder(3, 8)
	fmt.Println("Result is ", result)

	proResult, message := proAdder(4, 6, 8)
	fmt.Println("Result is ", proResult)
	fmt.Println("Message: ", message)

}