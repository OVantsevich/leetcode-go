package _877_Minimize_Maximum_Pair_Sum_in_Array

import (
	"sort"
)

func minPairSum(nums []int) int {
	sort.Ints(nums)
	var m int
	for i := range nums {
		m = max(m, nums[i]+nums[len(nums)-1-i])
	}
	return m
}
