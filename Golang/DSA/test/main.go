package main

func topKFrequent(nums []int, k int) []int {
	diary := make(map[int]int)
	bucketArray := make([][]int, len(nums)+1)

	for _, num := range nums {
		diary[num] += 1
	}

	for key := range diary {
		bucketArray[diary[key]] = append(bucketArray[diary[key]], key)
	}

	result := make([]int, 0)

	for i := len(bucketArray) - 1; k > 0; i-- {
		if len(bucketArray[i]) != 0 {
			for _, x := range bucketArray[i] {
				result = append(result, x)
				k--
			}
		}

	}

	return result
}

func main() {
	topKFrequent([]int{1, 1, 2, 1, 3, 4, 2}, 2)
}
