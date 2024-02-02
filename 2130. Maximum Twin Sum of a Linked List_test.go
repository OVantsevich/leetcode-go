package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_pairSum(t *testing.T) {
	require.Equal(t, 6, pairSum(
		&ListNode{
			Val: 5,
			Next: &ListNode{
				Val: 4,
				Next: &ListNode{
					Val: 2,
					Next: &ListNode{
						Val:  1,
						Next: nil,
					},
				},
			},
		},
	))
}
