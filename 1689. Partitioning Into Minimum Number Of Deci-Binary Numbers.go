package main

func minPartitions(n string) int {
	var result byte
	for i := 0; i < len(n); i++ {
		currentResult := n[i] - 48
		if currentResult == 9 {
			return int(currentResult)
		} else if result < currentResult {
			result = currentResult
		}
	}

	return int(result)
}
