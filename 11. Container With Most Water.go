package main

func maxArea(height []int) int {
	s, e := 0, len(height)-1

	var result int
	for s != e {
		volume := (e - s) * min(height[e], height[s])

		if volume > result {
			result = volume
		}

		if height[e] > height[s] {
			s++
		} else {
			e--
		}
	}

	return result
}
