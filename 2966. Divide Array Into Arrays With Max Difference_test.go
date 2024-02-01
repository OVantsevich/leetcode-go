package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_divideArray(t *testing.T) {
	require.Equal(t, [][]int{{1, 1, 3}, {3, 4, 5}, {7, 8, 9}}, divideArray([]int{1, 3, 4, 8, 7, 9, 3, 5, 1}, 2))
}
