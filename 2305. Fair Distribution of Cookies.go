package main

import "math"

var children []int
var cookiesG []int
var itr int

func distributeCookies(cookies []int, k int) int {
	children = make([]int, k)
	cookiesG = cookies

	return Unfairness()
}

func Unfairness() int {
	unfairness := math.MaxInt

	if itr == len(cookiesG) {
		return Max(children)
	}

	var newUnfairness int
	for i := 0; i < len(children); i++ {
		children[i] += cookiesG[itr]
		itr++
		newUnfairness = Unfairness()
		itr--
		children[i] -= cookiesG[itr]
		if newUnfairness < unfairness {
			unfairness = newUnfairness
		}
		if children[i] == 0 {
			break
		}
	}
	return unfairness
}

func Max(slice []int) (max int) {
	for _, e := range slice {
		if e > max {
			max = e
		}
	}
	return
}
