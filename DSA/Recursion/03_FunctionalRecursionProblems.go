package main

// reverse the array
func reverseArrayHelper(nums []int, l int, r int) []int {

	if l > r {
		return nums
	}

	// swap
	// temp := nums[l]
	// nums[l] = nums[r]
	// nums[r] = temp

	nums[l], nums[r] = nums[r], nums[l]

	return reverseArrayHelper(nums, l+1, r-1)
}

func reverseArray(nums []int) []int {
	if nums == nil {
		return make([]int, 0)
	}

	if len(nums) <= 1 {
		return nums
	}

	return reverseArrayHelper(nums, 0, len(nums)-1)
}
