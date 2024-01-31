package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_bulbSwitch(t *testing.T) {
	require.Equal(t, 1, bulbSwitch(3))
	require.Equal(t, 0, bulbSwitch(0))
	require.Equal(t, 1, bulbSwitch(1))
	require.Equal(t, 2, bulbSwitch(4))
}
