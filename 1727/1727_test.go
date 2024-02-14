package _727

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_largestSubmatrix(t *testing.T) {
	require.Equal(t, 4, largestSubmatrix([][]int{{0, 0, 1}, {1, 1, 1}, {1, 0, 1}}))
}
