package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion in golang!")

	// Reader
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please give rating for our pizza out of 10: ")
	input, _ := reader.ReadString('\n')

	// to remove '/n'
	rating := strings.TrimSpace(input)

	fmt.Printf("Type of rating is %T\n", rating)

	floatRating, err := strconv.ParseFloat(rating, 64)
	intRating := int(math.Round(floatRating))

	if err == nil {
		if intRating > 0 && intRating < 11 {
			fmt.Println("Thanks for rating, ", intRating)
			fmt.Println("Rating + 1 is ", intRating+1)
		} else {
			fmt.Println("Enter between 0 and 10")
		}
	} else {
		panic(err)
	}

}
