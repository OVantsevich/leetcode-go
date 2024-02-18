package _402

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestName(t *testing.T) {
	require.Equal(t, 0, mostBooked(2, [][]int{{1, 5}, {0, 10}, {2, 7}, {3, 4}}))
	require.Equal(t, 0, mostBooked(4, [][]int{{18, 19}, {3, 12}, {17, 19}, {2, 13}, {7, 10}}))
	require.Equal(t, 0, mostBooked(4, [][]int{{68331, 77431}, {7172, 49753}, {66587, 99610}, {70825, 81081}, {75611, 90380}, {2796, 98961}}))
}
