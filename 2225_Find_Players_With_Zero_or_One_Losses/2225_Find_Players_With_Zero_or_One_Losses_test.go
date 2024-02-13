package _225_Find_Players_With_Zero_or_One_Losses

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_findWinners(t *testing.T) {
	require.Equal(t, [][]int{{1, 2, 10}, {4, 5, 7, 8}}, findWinners([][]int{{1, 3}, {2, 3}, {3, 6}, {5, 6}, {5, 7}, {4, 5}, {4, 8}, {4, 9}, {10, 4}, {10, 9}}))
}
