package day7

import (
	"strconv"
	"strings"
	"testing"

	"github.com/brainboxweb/advent-2021/helpers"
	"github.com/stretchr/testify/assert"
)

func TestCalls(t *testing.T) {
	tests := []struct {
		positions []int
		expected  int
	}{
		{
			positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected:  37,
		},
		{
			positions: testData(),
			expected:  341534,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := Part1(tt.positions)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestCalls2(t *testing.T) {
	tests := []struct {
		positions []int
		expected  int
	}{
		{
			positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			expected:  168,
		},
		{
			positions: testData(),
			expected:  93397632,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := Part2(tt.positions)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func testData() []int {
	dataString := helpers.GetDataString("../data/day7.txt")
	data := strings.Split(dataString[0], ",")
	testData := []int{}
	for _, val := range data {
		item, err := strconv.Atoi(val)
		if err != nil {
			panic("not expected")
		}
		testData = append(testData, item)
	}
	return testData
}
