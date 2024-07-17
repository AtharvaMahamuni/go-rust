package main

import "fmt"

func main() {
	fmt.Println("Welcome to test the loops!")

	days := []string{"Sunday", "Tuesday", "Wednesday", "Friday", "Saturday"}
	fmt.Println(days)


	// Regular for loop
	// for i:=0; i<len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// Range loop
	// for i:= range days {
	// 	fmt.Println(days[i])
	// }

	// for index, day := range days {
	// 	fmt.Printf("The day at index %v, is %v\n", index, day)
	// }

	rougueValue := 1 

	// like a while loop
	for rougueValue < 10 {

		if rougueValue == 2 {
			goto lco
		}

		// if rougueValue==5 {
		// // 	continue
		// 	break
		// }

		fmt.Println("Value is : ", rougueValue)
		
		rougueValue++;
	}

	lco:
		fmt.Println("Hey, you are 2.")
}