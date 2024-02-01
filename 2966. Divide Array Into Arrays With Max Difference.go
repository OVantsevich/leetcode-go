package main

import (
	"sort"
)

func divideArray(nums []int, k int) [][]int {
	sort.Ints(nums)
	result := make([][]int, 0, len(nums)/3)
	for i := 0; i < len(nums); i += 3 {
		if nums[i+2]-nums[i] > k {
			return nil
		}
		result = append(result, nums[i:i+3])
	}
	return result
}
