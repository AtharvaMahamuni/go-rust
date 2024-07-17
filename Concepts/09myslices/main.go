package main

import (
	"fmt"
	"sort"	
)

func main() {
	fmt.Println("Welcome to slices in golang!")

	var fruitList = []string{"Apple", "Tomato", "Mango"}
	fmt.Println(fruitList)

	// fmt.Printf("Type of fruitList : %T\n", fruitList)

	// fruitList = append(fruitList, "Banana", "Peach", "Orange")
	// fmt.Println(fruitList)

	// fmt.Println(fruitList[1:3])
	// fmt.Println(fruitList[1:])
	// fmt.Println(fruitList[:3])

	// It will create array and append method can't be used
	// highScores := [4]int{} 

	// It will create the slice and you can use the append method
	highScores := make([]int, 4) 
	
	highScores[0] = 123
	highScores[1] = 213
	highScores[2] = 312
	highScores[3] = 321

	fmt.Println(highScores)

	highScores = append(highScores, 111, 222, 333)
	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)

	// Removing value at some index in slice
	// [111 123 213 222 312 321 333] removing value at 2nd index
	var index int = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}