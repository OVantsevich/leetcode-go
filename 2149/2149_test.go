package _149

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_rearrangeArray(t *testing.T) {
	require.Equal(t, []int{3, -2, 1, -5, 2, -4}, rearrangeArray([]int{3, 1, -2, -5, 2, -4}))
}
