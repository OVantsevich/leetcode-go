package _79

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_kthGrammar(t *testing.T) {
	require.Equal(t, 1, kthGrammar(2, 2))
	require.Equal(t, 0, kthGrammar(1, 1))
	require.Equal(t, 0, kthGrammar(2, 1))
}
