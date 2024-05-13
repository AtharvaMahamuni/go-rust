package main

// FIXME: go run --race main.go
// use `--race` to find out the race condition in code.

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Let's see race condition")
	var scores = []int{0}

	wg := &sync.WaitGroup{}
	mt := &sync.Mutex{} // using mutex will solve the race condition

	wg.Add(3)
	go func(wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("One R")

		mt.Lock()
		scores = append(scores, 1)
		mt.Unlock()

		wg.Done()
	}(wg, mt)

	go func(wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("Two R")

		mt.Lock()
		scores = append(scores, 2)
		mt.Unlock()

		wg.Done()
	}(wg, mt)

	go func(wg *sync.WaitGroup, mt *sync.Mutex) {
		fmt.Println("Three R")

		mt.Lock()
		scores = append(scores, 3)
		mt.Unlock()

		wg.Done()
	}(wg, mt)

	wg.Wait()

	fmt.Println(scores)
}
