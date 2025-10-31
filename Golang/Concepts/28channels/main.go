package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("This is channel")

	myCh := make(chan int, 2)
	wt := &sync.WaitGroup{}

	// // fatal error: all goroutines are asleep - deadlock!
	// myCh <- 5
	// fmt.Println(<-myCh)

	wt.Add(2)

	// <-chan receive only
	go func(ch <-chan int, wt *sync.WaitGroup) {
		val, isChannelOpen := <-ch

		if isChannelOpen {
			fmt.Println(val)
		} else {
			fmt.Println("Channel is closed -", isChannelOpen)
		}
		// fmt.Println(<-ch)
		wt.Done()
	}(myCh, wt)

	// chan<- send only
	go func(ch chan<- int, wt *sync.WaitGroup) {

		// ch <- 5
		ch <- 7
		close(ch)

		wt.Done()
	}(myCh, wt)

	wt.Wait()
}
