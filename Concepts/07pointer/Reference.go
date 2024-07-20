package main

import "fmt"

func Reference() {
	var score int = 300
	fmt.Println(score)  // 300
	fmt.Println(&score) // 0xc000096090

	var myptr *int = &score
	fmt.Println(*myptr) // 300
	fmt.Println(myptr)  // 0xc000096090
	fmt.Println(&myptr) // 0xc0000a8020

	*myptr += 1
	fmt.Println(score) // 301
}
