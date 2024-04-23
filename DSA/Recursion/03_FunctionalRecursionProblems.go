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

func checkPalindromeHelper(str string, l int, r int) bool {
	if l > r {
		return true
	}

	if str[l] != str[r] {
		return false
	}

	return checkPalindromeHelper(str, l+1, r-1)
}

func checkPalindrome(str string) bool {
	if str == "" {
		return false
	}

	if len(str) == 1 {
		return true
	}

	return checkPalindromeHelper(str, 0, len(str)-1)
}
