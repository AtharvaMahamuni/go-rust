package main

import "fmt"

func nTimesNameHelper(name string, i int, n int) {
	if i > n {
		return
	}

	fmt.Println(name)
	nTimesNameHelper(name, i+1, n)
}

func printNameNTimes(name string, n int) {
	if n == 0 {
		return
	}

	nTimesNameHelper(name, 1, n)
}

func printNnumbersHelper(i int, n int) {
	if i > n {
		return
	}
	print(i, " ")

	printNnumbersHelper(i+1, n)
}

func printNnumbers(n int) {
	if n < 0 {
		return
	}

	printNnumbersHelper(1, n)
}

func printNto1BackTrackHelper(i int, n int) {

	if i > n {
		return
	}

	printNto1BackTrackHelper(i+1, n)
	print(i, " ")
}

func printNto1BackTrack(n int) {
	if n < 0 {
		return
	}
	printNto1BackTrackHelper(1, n)
}

func printNto1Helper(n int) {
	if n < 1 {
		return
	}

	print(n, " ")
	printNto1Helper(n - 1)
}

func printNto1(n int) {
	if n < 0 {
		return
	}
	printNto1Helper(n)
}
