package main

func sum(arr []int) int {
	var s int

	for _, n := range arr {
		s += n
	}

	return s
}

func minOperations(nums []int, x int) int {
	sum := sum(nums)

	if sum == x {
		return len(nums)
	}
	if sum < x {
		return -1
	}

	result := -0
	currSum := 0
	j := 0
	medSum := sum - x
	for i := 0; i < len(nums); i++ {
		currSum += nums[i]

		for currSum > medSum && j <= i {
			currSum -= nums[j]
			j++
		}

		if currSum == medSum {
			result = max(result, i-j+1)
		}
	}
	if result == 0 {
		return -1
	}
	return len(nums) - result
}
