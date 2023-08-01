package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_distributeCookies(t *testing.T) {
	unfairness := distributeCookies([]int{8, 15, 10, 20, 8}, 2)
	require.Equal(t, 31, unfairness)

	unfairness = distributeCookies([]int{6, 1, 3, 2, 2, 4, 1, 2}, 3)
	require.Equal(t, 7, unfairness)
}
