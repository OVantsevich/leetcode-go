package _877_Minimize_Maximum_Pair_Sum_in_Array

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPairSum(t *testing.T) {
	require.Equal(t, 8, minPairSum([]int{3, 5, 4, 2, 4, 6}))
}
