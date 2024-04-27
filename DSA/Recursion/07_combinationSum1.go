package main

// Input: candidates = [2,3,6,7], target = 7
// Output: [[2,2,3],[7]]
// Explanation:
// 2 and 3 are candidates, and 2 + 2 + 3 = 7. Note that 2 can be used multiple times.
// 7 is a candidate, and 7 = 7.
// These are the only two combinations.

// one element can be used multiple times

// unfortunately following code failing for few cases

func combinationSum1Helper(candidates []int, target int, index int, holder []int, result [][]int) [][]int {
	if target < 0 {
		return make([][]int, 0)
	}

	if target == 0 {
		return append(result, holder)
	}

	if index == len(candidates) {
		return make([][]int, 0)
	}

	leftValues := make([][]int, 0)

	if candidates[index] <= target {
		leftValues = combinationSum1Helper(candidates, target-candidates[index], index, append(holder, candidates[index]), result)
	}
	rightValues := combinationSum1Helper(candidates, target, index+1, holder, result)

	result = append(leftValues, rightValues...)
	return result
}

func combinationSum(candidates []int, target int) [][]int {
	return combinationSum1Helper(candidates, target, 0, make([]int, 0), make([][]int, 0))
}
