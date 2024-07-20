package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string // size 4 of this array can't be changed afterwards, it is fixed for length 4
	// neither be able to use methods like append nor be able to reduce the length of the array once declared.

	// Assigning the values
	fruitList[0] = "Apple"
	fruitList[1] = "Mango"
	fruitList[3] = "Peach"

	// fruitList = fruitList[:2] // can't be done with array

	var fruitListSlice []string
	fruitListSlice = fruitList[:2] // but possible with the slices
	fruitListSlice = fruitListSlice[:1]
	fmt.Println(fruitListSlice)

	fmt.Println("FruitList", fruitList)

	var numbers [7]int
	fmt.Println(numbers)

	var vegList = [3]string{"potato", "beans", "cabagge"}
	fmt.Println(len(vegList))
	fmt.Printf("Type of vegList %T\n", vegList)
}
