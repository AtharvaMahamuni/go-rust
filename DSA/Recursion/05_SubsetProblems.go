package main

import (
	"fmt"
)

func subsetStringHelper(s string, index int, holder string) {

	if index == len(s) {
		fmt.Print("'", holder, "' ")
		return
	}

	subsetStringHelper(s, index+1, holder+string(s[index]))
	subsetStringHelper(s, index+1, holder)
}

func subsetString(s string) {
	subsetStringHelper(s, 0, "")
	fmt.Println()
}

func subsetArrayHelper(nums []int, index int, holder []int) {
	if index == len(nums) {
		fmt.Print(holder, " ")
		return
	}

	subsetArrayHelper(nums, index+1, holder)
	subsetArrayHelper(nums, index+1, append(holder, nums[index]))
}

func subsetArray(nums []int) {
	subsetArrayHelper(nums, 0, make([]int, 0))
	fmt.Println()
}
