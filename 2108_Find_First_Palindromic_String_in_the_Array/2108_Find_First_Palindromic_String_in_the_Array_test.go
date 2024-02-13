package _108_Find_First_Palindromic_String_in_the_Array

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_firstPalindrome(t *testing.T) {
	require.Equal(t, "ada", firstPalindrome([]string{"abc", "car", "ada", "racecar", "cool"}))
}
