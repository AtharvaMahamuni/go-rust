package main

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
