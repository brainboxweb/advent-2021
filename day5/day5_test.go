package day5_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day5"
	"github.com/brainboxweb/advent-2021/helpers"
)

const testDataPath = "../testdata"

func TestDay5(t *testing.T) {
	dataSet := helpers.GetDataString(testDataPath + "/day5.txt")
	tests := []struct {
		in       []string
		size     int
		expected int
	}{
		{
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			10,
			5,
		},
		{
			dataSet,
			999,
			6687,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			ans := day5.Part1(tt.in, tt.size)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestDay5Part2(t *testing.T) {
	dataSet := helpers.GetDataString(testDataPath + "/day5.txt")
	tests := []struct {
		in       []string
		size     int
		expected int
	}{
		{
			[]string{
				"0,9 -> 5,9",
				"8,0 -> 0,8",
				"9,4 -> 3,4",
				"2,2 -> 2,1",
				"7,0 -> 7,4",
				"6,4 -> 2,0",
				"0,9 -> 2,9",
				"3,4 -> 1,4",
				"0,0 -> 8,8",
				"5,5 -> 8,2",
			},
			10,
			12,
		},
		{
			dataSet,
			999,
			19851,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			ans := day5.Part2(tt.in, tt.size)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
