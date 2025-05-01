package day3_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day3"
	"github.com/brainboxweb/advent-2021/helpers"
)

const testDataPath = "../testdata"

func TestDay3(t *testing.T) {
	dataSet := helpers.GetDataString(testDataPath + "/day3.txt")
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			198,
		},
		{
			dataSet,
			1092896,
		},
	}
	for _, tt := range tests {
		testname := "Day3"
		t.Run(testname, func(t *testing.T) {
			ans := day3.Part1(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestDay3Part2(t *testing.T) {
	dataSet := helpers.GetDataString(testDataPath + "/day3.txt")
	tests := []struct {
		in       []string
		expected int
	}{
		{
			[]string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			},
			230,
		},
		{

			dataSet,
			4672151,
		},
	}
	for _, tt := range tests {
		testname := "Day3"
		t.Run(testname, func(t *testing.T) {
			ans := day3.Part2(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
