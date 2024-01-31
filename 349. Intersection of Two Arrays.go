package main

func intersection(nums1 []int, nums2 []int) []int {
	mask := make([]bool, 1000)
	result := make([]int, 0)

	for i := range nums1 {
		if !mask[nums1[i]] {
			for j := range nums2 {
				if nums1[i] == nums2[j] {
					mask[nums1[i]] = true
					result = append(result, nums1[i])
					break
				}
			}
		}
	}
	return result
}
