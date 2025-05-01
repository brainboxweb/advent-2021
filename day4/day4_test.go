package day4_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day4"
	"github.com/brainboxweb/advent-2021/helpers"
)

const testDataPath = "../testdata"

func TestPart1(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			testDataPath + "/day4_test.txt",
			4512,
		},
		{
			testDataPath + "/day4.txt",
			89001,
		},
	}
	for _, tt := range tests {
		testname := "Run the calls "
		t.Run(testname, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			ans := day4.Part1(data)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		dataFile string
		expected int
	}{
		{
			testDataPath + "/day4_test.txt",
			1924,
		},
		{
			testDataPath + "/day4.txt",
			7296,
		},
	}
	for _, tt := range tests {
		testname := "Run the calls "
		t.Run(testname, func(t *testing.T) {
			data := helpers.GetDataString(tt.dataFile)
			ans := day4.Part2(data)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
