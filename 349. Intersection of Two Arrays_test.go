package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_intersection(t *testing.T) {
	require.Equal(t, []int{2}, intersection([]int{1, 2, 2, 1}, []int{2, 2}))
}
