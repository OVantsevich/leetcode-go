package _69_Majority_Element

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_majorityElement(t *testing.T) {
	require.Equal(t, 3, majorityElement([]int{3, 2, 3}))
	require.Equal(t, 2, majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
}
