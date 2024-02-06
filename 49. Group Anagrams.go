package main

import "sort"

func groupAnagrams(strs []string) [][]string {
	result := make([][]string, 0)
	orderedStrings := make(map[string]int)

	for _, str := range strs {
		chars := []byte(str)
		sort.Slice(chars, func(i, j int) bool {
			return chars[i] < chars[j]
		})
		if idx, ok := orderedStrings[string(chars)]; ok {
			result[idx] = append(result[idx], str)
		} else {
			result = append(result, []string{str})
			orderedStrings[string(chars)] = len(result) - 1
		}
	}
	return result
}
