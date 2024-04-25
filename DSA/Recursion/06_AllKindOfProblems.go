package main

import "fmt"

// subsequences sum equal to the target
// {1, 2, 1}, Target=2 => {[1,1], [2]}

func subsequenceSumHelper(nums []int, index int, holder []int, target int, result [][]int) [][]int {

	if index == len(nums) {
		sum := 0
		for _, num := range holder {
			sum += num
		}
		if sum == target {
			result = append(result, holder)
		}
		return result
	}

	leftResults := subsequenceSumHelper(nums, index+1, append(holder, nums[index]), target, result)
	rightResults := subsequenceSumHelper(nums, index+1, holder, target, result)

	return append(leftResults, rightResults...)
}

func subSequenceSum(nums []int, target int) [][]int {
	return subsequenceSumHelper(nums, 0, make([]int, 0), target, make([][]int, 0))
}

// print only 1 sum of subsequence

func onlyOneSubsequenceSumHelper(nums []int, index int, holder []int, target int) bool {

	if index == len(nums) {
		sum := 0

		for _, num := range holder {
			sum += num
		}

		if sum == target {
			fmt.Println(holder)
			return true
		}

		return false
	}

	if onlyOneSubsequenceSumHelper(nums, index+1, holder, target) {
		return true
	}

	if onlyOneSubsequenceSumHelper(nums, index+1, append(holder, nums[index]), target) {
		return true
	}

	return false
}

func onlyOneSubsequenceSum(nums []int, target int) bool {
	return onlyOneSubsequenceSumHelper(nums, 0, make([]int, 0), target)
}

// count of subsequence sums

func countSubsequenceSumHelper(nums []int, index int, sum int, target int) int {

	if index == len(nums) {
		if sum == target {
			return 1
		}

		return 0
	}

	leftCounts := countSubsequenceSumHelper(nums, index+1, sum+nums[index], target)
	rightCounts := countSubsequenceSumHelper(nums, index+1, sum, target)

	return leftCounts + rightCounts
}

func countSubsequenceSum(nums []int, target int) int {
	return countSubsequenceSumHelper(nums, 0, 0, target)
}
