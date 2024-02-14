package _149

func rearrangeArray(nums []int) []int {
	p := 0
	n := 1
	result := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if nums[i] < 0 {
			result[n] = nums[i]
			n += 2
		} else {
			result[p] = nums[i]
			p += 2
		}
	}
	return result
}
