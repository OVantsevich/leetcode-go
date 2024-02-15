package _971

func sm(arr []int) int64 {
	var s int64

	for _, n := range arr {
		s += int64(n)
	}

	return s
}

func largestPerimeter(nums []int) int64 {
	sum := sm(nums)

	var max, index int

	for {
		for i, v := range nums {
			if max < v {
				max = v
				index = i
			}
		}

		if int64(max) >= sum-int64(max) && len(nums) > 2 {
			sum -= int64(max)
			nums[index] = nums[len(nums)-1]
			nums = nums[:len(nums)-1]
			max = 0
		} else {
			break
		}
	}
	if len(nums) <= 2 {
		return -1
	}
	return sum
}
