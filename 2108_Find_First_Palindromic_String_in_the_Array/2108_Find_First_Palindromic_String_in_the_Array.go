package _108_Find_First_Palindromic_String_in_the_Array

func firstPalindrome(words []string) (result string) {
	for _, word := range words {
		flag := true
		for i := 0; i < len(word)/2; i++ {
			if word[i] != word[len(word)-1-i] {
				flag = false
			}
		}
		if flag {
			return word
		}
	}
	return
}
