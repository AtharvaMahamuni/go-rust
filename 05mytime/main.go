package main

import (
	"fmt"
	"time"
)


func main() {
	fmt.Println("Welcome, to explore the time!");

	// using current time
	currentTime := time.Now();

	fmt.Println("Current time is: ", currentTime);
	fmt.Println("In format: ", currentTime.Format("02-01-2006 Monday 15:04:05"));

	
	fmt.Println("----");

	
	// creating time
	createdTime := time.Date(2000, time.April, 05, 04, 30, 00, 00, time.Local);

	fmt.Println("Created time is: ", createdTime);
	fmt.Println("In format: ", createdTime.Format("02-01-2006 Monday 15:04:05"));

}
