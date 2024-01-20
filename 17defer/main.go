package main

import "fmt"

func main() {
	defer fmt.Println("Three")
	defer fmt.Println("Two")
	defer fmt.Println("One")
	// defer fmt.Println("World!")
	fmt.Println("Hello")
	deferCounter()
}

func deferCounter() {
	defer fmt.Println()
	for i:=0; i<5; i++ {
		defer fmt.Print(i)
	}
}