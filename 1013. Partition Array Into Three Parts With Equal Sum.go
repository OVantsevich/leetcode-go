package main

//func sum(arr []int) int {
//	var s int
//
//	for _, n := range arr {
//		s += n
//	}
//
//	return s
//}

func canThreePartsEqualSum(arr []int) bool {
	s := sum(arr)
	if s%3 != 0 {
		return false
	}

	var p1 int
	var i int
	for {
		p1 += arr[i]
		i++
		if !(p1 != s/3 && i < len(arr)) {
			break
		}
	}
	if p1 != s/3 || i == len(arr) {
		return false
	}

	var p2 int
	j := i
	for {
		p2 += arr[j]
		j++
		if p2 != s/3 && j < len(arr) {
			continue
		}
		break
	}
	if p2 != s/3 || j == len(arr) {
		return false
	}

	return true
}
