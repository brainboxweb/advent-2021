package day2_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day2"
	"github.com/brainboxweb/advent-2021/helpers"
)

func TestDay2(t *testing.T) {
	dataSet := helpers.GetDataString("../data/day2.txt")
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			150,
		},
		{
			dataSet,
			2272262,
		},
	}
	for _, tt := range tests {
		testname := "Day2_1"
		t.Run(testname, func(t *testing.T) {
			ans := day2.Part1(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestDay2Part2(t *testing.T) {
	dataSet := helpers.GetDataString("../data/day2.txt")
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{
				"forward 5",
				"down 5",
				"forward 8",
				"up 3",
				"down 8",
				"forward 2",
			},
			900,
		},
		{
			dataSet,
			2134882034,
		},
	}
	for _, tt := range tests {
		testname := "Day2_1"
		t.Run(testname, func(t *testing.T) {
			ans := day2.Part2(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
