package main

func checkValid(matrix [][]int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		c := 0
		r := 0

		for j := 0; j < n; j++ {
			if (c&(1<<(matrix[i][j]-1))) > 0 ||
				(r&(1<<(matrix[j][i]-1))) > 0 {
				return false
			}

			c |= 1 << (matrix[i][j] - 1)
			r |= 1 << (matrix[j][i] - 1)
		}
	}

	return true
}
