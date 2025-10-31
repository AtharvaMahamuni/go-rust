package main

import "fmt"

// functions closure helps for data isolation
// example from go-tour

func ASum() func() {
	sum := 0

	return func() {
		sum += 1
		fmt.Println(sum)
	}
}

func main() {
	addr1 := ASum()
	for i := 0; i < 5; i++ {
		addr1()
	}

}
