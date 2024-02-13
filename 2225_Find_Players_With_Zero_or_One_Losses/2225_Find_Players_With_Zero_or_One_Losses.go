package _225_Find_Players_With_Zero_or_One_Losses

import "sort"

func findWinners(matches [][]int) [][]int {
	zero := make(map[int]struct{})
	one := make(map[int]struct{})
	more := make(map[int]struct{})

	for _, match := range matches {
		var ok bool
		_, ok = more[match[0]]
		if !ok {
			_, ok = one[match[0]]
			if !ok {
				zero[match[0]] = struct{}{}
			}
		}

		_, ok = more[match[1]]
		if !ok {
			_, ok = one[match[1]]
			if ok {
				delete(one, match[1])
				more[match[1]] = struct{}{}
				continue
			}
			delete(zero, match[1])
			one[match[1]] = struct{}{}
		}
	}

	resultZero := make([]int, 0, len(zero))
	for i := range zero {
		resultZero = append(resultZero, i)
	}
	resultOne := make([]int, 0, len(one))
	for i := range one {
		resultOne = append(resultOne, i)
	}
	sort.Ints(resultZero)
	sort.Ints(resultOne)
	return [][]int{resultZero, resultOne}
}
