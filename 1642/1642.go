package _642

import (
	"sort"
)

func insertSorted(s []int, e int) []int {
	s = append(s, 0)
	i := sort.Search(len(s), func(i int) bool { return s[i] < e })
	copy(s[i+1:], s[i:])
	s[i] = e
	return s
}

func furthestBuilding(heights []int, bricks int, ladders int) int {
	sortedList := make([]int, 0, ladders+1)

	for i := 0; i < len(heights)-1; i++ {
		diff := heights[i+1] - heights[i]
		if diff > 0 {
			sortedList = insertSorted(sortedList, diff)
			if len(sortedList) > ladders {
				if bricks >= sortedList[len(sortedList)-1] {
					bricks -= sortedList[len(sortedList)-1]
					sortedList = sortedList[:len(sortedList)-1]
				} else {
					return i
				}
			}
		}
	}

	return len(heights) - 1
}
