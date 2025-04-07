package bingo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalls(t *testing.T) {
	tests := []struct {
		calls    []int
		expected [][]int
	}{
		{
			calls: []int{7, 4, 9, 5, 11},
			expected: [][]int{
				{0, 0, 0, 1, 0},
				{0, 0, 0, 1, 0},
				{0, 1, 0, 0, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
			},
		},
		{
			calls: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21},
			expected: [][]int{
				{0, 0, 1, 1, 1},
				{0, 1, 1, 1, 0},
				{1, 1, 1, 0, 1},
				{0, 0, 0, 0, 1},
				{0, 0, 0, 0, 0},
			},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			bb := NewBoard()
			bb.rows = [][]int{
				{22, 13, 17, 11, 0},
				{8, 2, 23, 4, 24},
				{21, 9, 14, 16, 7},
				{6, 10, 3, 18, 5},
				{1, 12, 20, 15, 19},
			}
			for _, val := range tt.calls {
				bb.Call(val)
			}
			assert.Equal(t, tt.expected, bb.called)
		})
	}
}
