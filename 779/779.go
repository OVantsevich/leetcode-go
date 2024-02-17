package _79

var d = [2][2]int{
	{0, 1},
	{1, 0},
}

func kthGrammar(n int, k int) int {
	return kF(n, k-1)
}

func kF(n int, k int) int {
	if n == 1 {
		return 0
	}
	return d[kF(n-1, k/2)][k%2]
}
