package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string
	
	// Assigning the values
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	fmt.Println("FruitList", fruitList)

	var numbers [7]int
	fmt.Println(numbers)

	var  vegList = [3]string {"potato", "beans", "cabagge"}
	fmt.Println(len(vegList))
	fmt.Printf("Type of vegList %T\n", vegList)
}