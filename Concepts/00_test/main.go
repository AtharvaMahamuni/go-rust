package main

func twoSum(nums []int, target int) []int {
	var myMap = make(map[int]int)
	for i := 0; i < len(nums); i++ {
		myMap[nums[i]] = i
	}

	for i := 0; i < len(nums); i++ {
		complement := target - nums[i]
		_, ok := myMap[complement]
		if ok && myMap[complement] != i {
			result := []int{i, myMap[complement]}
			return result
		}
	}

	return nil
}

func main() {
	twoSum([]int{1, 2, 3, 4, 5}, 5)
}
