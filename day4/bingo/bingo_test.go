package bingo_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day4/bingo"
	"github.com/stretchr/testify/assert"
)

func TestCalls(t *testing.T) {
	rowsData := [][]int{
		{14, 21, 17, 24, 4},
		{10, 16, 15, 9, 19},
		{18, 8, 23, 26, 20},
		{22, 11, 13, 6, 5},
		{2, 0, 12, 3, 7},
	}
	tests := []struct {
		calls []int
		match int
		score int
	}{
		{
			calls: []int{7, 4, 9, 5, 11},
			match: 0,
			score: 0,
		},
		{
			calls: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21},
			match: 0,
			score: 0,
		},
		{
			calls: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24},
			match: 24,
			score: 188,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			bb := getBoard(rowsData)
			for _, val := range tt.calls {
				_, score := bb.Call(val)
				if score > 0 {
					assert.Equal(t, tt.score, score)
				}
			}
		})
	}
}

func getBoard(rowsData [][]int) bingo.Board {
	bb := bingo.NewBoard()
	bb.AddRows(rowsData)
	// for _, rowData := range rowsData {
	// 	bb.AddRow(rowData)
	// }

	return bb
}
