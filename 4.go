package main

import "log"

func longestIncreasingSubsequence(nums []int) []int {
	longest, longestIndex := -1, -1
	length := make([]int, len(nums))
	last := make([]int, len(nums))

	for i := range nums {
		length[i] = 1
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] && length[i] <= length[j] {
				length[i] = length[j] + 1
				last[i] = j
			}
		}
		if longest < length[i] {
			longest = length[i]
			longestIndex = i
		}
	}

	lis := make([]int, longest)
	for i, j := longestIndex, longest - 1; i > 0; i, j = last[i], j - 1 {
		lis[j] = nums[i]
	}

	return lis
}

func main() {
	log.Print(longestIncreasingSubsequence([]int{6, 2, 5, 1, 7, 4, 8, 3}))
}