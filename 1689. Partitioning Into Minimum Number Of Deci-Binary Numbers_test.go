package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_minPartitions(t *testing.T) {
	require.Equal(t, 3, minPartitions("32"))
}
