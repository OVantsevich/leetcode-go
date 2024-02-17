package _642

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	require.Equal(t, 7, furthestBuilding([]int{4, 12, 2, 7, 3, 18, 20, 3, 19}, 10, 2))
	require.Equal(t, 0, furthestBuilding([]int{1, 2}, 0, 0))
	require.Equal(t, 3, furthestBuilding([]int{2, 7, 9, 12}, 5, 1))
}
