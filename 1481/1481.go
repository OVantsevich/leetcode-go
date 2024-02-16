package _481

import "sort"

func findLeastNumOfUniqueInts(arr []int, k int) int {
	if len(arr) <= k+1 {
		return len(arr) - k
	}
	nums := make(map[int]int, len(arr))

	for _, i := range arr {
		nums[i] = nums[i] + 1
	}

	count := make([]int, 0, len(nums))

	for _, n := range nums {
		count = append(count, n)
	}

	sort.Ints(count)

	i := 0
	for ; i < len(count); i++ {
		if k >= count[i] {
			k -= count[i]
		} else {
			break
		}
	}

	return len(count) - i
}
