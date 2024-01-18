package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch case in Golang!")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1

	fmt.Println("Value of dice : ", diceNumber)


	switch diceNumber {
	case 6:
		fmt.Println("You can move 6 spots and get one more chance")
		fallthrough // runs the below case as well
	case 1:
		fmt.Println("Dice value is 1 and you can open")
	case 2:
		fmt.Println("You can move 2 spot")
	case 3:
		fmt.Println("You can move 3 spot")
	case 4:
		fmt.Println("You can move 4 spot")
	case 5:
		fmt.Println("You can move 5 spot")
	}
}