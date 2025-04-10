package day8_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day8"
	"github.com/brainboxweb/advent-2021/helpers"
	"github.com/stretchr/testify/assert"
)

// Should be internal
func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day8_test.txt",
			26,
		},
		{
			"../data/day8.txt",
			512,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day8.Part1(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			"../data/day8_test.txt",
			61229,
		},
		{
			"../data/day8.txt",
			1091165,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			result := day8.Part2(data)
			assert.Equal(t, tt.expected, result)
		})
	}
}
