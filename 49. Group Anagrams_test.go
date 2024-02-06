package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_groupAnagrams(t *testing.T) {
	require.Equal(t, [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}, groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))

}
