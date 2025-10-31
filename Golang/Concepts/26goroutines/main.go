package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Do not communicate by sharing memory; instead, share memory by communicating.

var wg sync.WaitGroup // pointer
var mut sync.Mutex    // pointer

// func greeter(s string) {
// 	for i := 0; i < 6; i++ {
// 		time.Sleep(1 * time.Microsecond)
// 		fmt.Println(i+1, s)
// 	}
// }

var signals []string

func getStatusCode(endpoint string) {
	defer wg.Done()

	res, err := http.Get(endpoint)

	if err != nil {
		// panic(err)
		fmt.Println("OOps error")
	} else {
		mut.Lock()
		signals = append(signals, endpoint)
		mut.Unlock()

		fmt.Printf("%d status code for %s\n", res.StatusCode, endpoint)
	}

}

func main() {
	// go greeter("Hello")
	// greeter("World")

	endPointList := []string{
		"https://reddit.com",
		"https://go.dev",
		"https://fb.com",
		"https://google.com",
		"https://github.com",
	}

	for _, endPoint := range endPointList {
		go getStatusCode(endPoint)
		wg.Add(1)
	}

	wg.Wait()

	fmt.Println(signals)

}
