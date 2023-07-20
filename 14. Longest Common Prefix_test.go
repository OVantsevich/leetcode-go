package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_longestCommonPrefix(t *testing.T) {
	r := longestCommonPrefix([]string{"flower", "flow", "flight"})
	require.Equal(t, "fl", r)

	r = longestCommonPrefix([]string{"dog", "racecar", "car"})
	require.Equal(t, "", r)

	r = longestCommonPrefix([]string{""})
	require.Equal(t, "", r)
}
