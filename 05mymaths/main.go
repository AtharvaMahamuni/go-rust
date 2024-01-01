package main

import (
	"fmt"
	// "math/rand"
	// "time"
	"crypto/rand"
	"math/big"
)

func main() {
	fmt.Println("Welcome, to maths code!");

	// var number1 int = 34;
	// var number2 float64 = 5.6;

	// we need to typecast one of the two to do addition, otherwise gives an error.
	// fmt.Println("Addition is : ", number1+int(number2));


	// random numbers (with time as a seed and math/rand pkg

	// rand.Seed(time.Now().UnixNano());
	// fmt.Println("Random number is : ", rand.Intn(5) + 1); // +1, because [0,5)

	// random number with crypto/rand pkg

	myRandNum, _ := rand.Int(rand.Reader, big.NewInt(5));

	one := big.NewInt(1);

	fmt.Println("Random numer is : ", myRandNum.Add(myRandNum, one));
	// fmt.Printf("%T \n", myRandNum);
}
