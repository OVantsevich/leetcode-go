package main

func longestCommonPrefix(strs []string) string {
	prfx := strs[0]

	for i := 0; i < len(prfx); i++ {
		for j := 1; j < len(strs); j++ {
			if len(strs[j]) == i || strs[j][i] != prfx[i] {
				return string(prfx[:i])
			}
		}
	}
	return string(prfx)
}
