package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkValid(t *testing.T) {
	require.Equal(t, true, checkValid([][]int{{1, 2, 3}, {3, 1, 2}, {2, 3, 1}}))
}
