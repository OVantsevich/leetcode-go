package main

import "math"

func isSquare(n int) bool {
	sqrt := int(math.Sqrt(float64(n)))
	return sqrt*sqrt == n
}

func numSquares(n int) int {
	if isSquare(n) {
		return 1
	}

	for n&3 == 0 {
		n >>= 2
	}
	if n&7 == 7 {
		return 4
	}

	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if isSquare(n - i*i) {
			return 2
		}
	}

	return 3
}
