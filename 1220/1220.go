package _220

func countVowelPermutation(n int) int {
	mod := int(1e9 + 7)

	var (
		a = 1
		e = 1
		i = 1
		o = 1
		u = 1
	)

	for j := 0; j < n-1; j++ {
		a, e, i, o, u =
			(e+i+u)%mod, (a+i)%mod, (e+o)%mod, i, (i+o)%mod
	}

	return (a + e + i + o + u) % mod
}
