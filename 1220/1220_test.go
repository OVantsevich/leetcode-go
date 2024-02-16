package _220

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_countVowelPermutation(t *testing.T) {
	require.Equal(t, 68, countVowelPermutation(5))
	require.Equal(t, 18208803, countVowelPermutation(144))
}
