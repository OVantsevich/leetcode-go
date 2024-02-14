package _727

import (
	"sort"
)

func largestSubmatrix(matrix [][]int) int {
	for i := 0; i < len(matrix[0]); i++ {
		for j := 1; j < len(matrix); j++ {
			matrix[j][i] += matrix[j-1][i] * matrix[j][i]
		}
	}

	for i := 0; i < len(matrix); i++ {
		sort.Ints(matrix[i])
	}

	var result int
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			size := matrix[j][i] * (len(matrix[0]) - j)
			if size > result {
				result = size
			}
		}
	}
	return result
}
