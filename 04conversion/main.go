package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome, to type conversion")

	// Reader
	reader := bufio.NewReader(os.Stdin);

	fmt.Print("Please enter rating for our Pizza: ");
	input, _ := reader.ReadString('\n');

	// removing '/n' from input
	rating := strings.TrimSpace(input);

	fmt.Printf("Thanks for rating, %s! \n", rating);

	fmt.Printf("Type of rating is : %T \n", rating);

	// type conversion 
	floatRating, err := strconv.ParseFloat(rating, 64);

	if err != nil {
		fmt.Println(err);
	} else {
		fmt.Println("Rating + 1 is : ", floatRating+1);
	}

}
