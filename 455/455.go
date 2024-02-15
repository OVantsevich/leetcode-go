package _55

import "sort"

func findContentChildren(g []int, s []int) int {
	var result int

	sort.Ints(g)
	sort.Ints(s)

	for i := 0; i < len(s) && result < len(g); i++ {
		if s[i] >= g[result] {
			result++
			continue
		}
	}

	return result
}
