package day1_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day1"
	"github.com/brainboxweb/advent-2021/helpers"
)

const testDataPath = "../testdata"

func TestPart1(t *testing.T) {
	dataSet, err := helpers.GetDataInt(testDataPath + "/day1.txt")
	if err != nil {
		panic("not expected")
	}
	tests := []struct {
		in       []int
		expected int
	}{
		{
			[]int{1, 2, 3},
			2,
		},
		{
			[]int{1, 2, 1},
			1,
		},
		{
			[]int{3, 2, 1},
			0,
		},
		{
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			7,
		},
		{
			dataSet,
			1387,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			ans := day1.Part1(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}

func TestDay21(t *testing.T) {
	dataSet, err := helpers.GetDataInt(testDataPath + "/day1.txt")
	if err != nil {
		panic("not expected")
	}
	tests := []struct {
		in       []int
		expected int
	}{
		{
			[]int{
				607,
				618,
				618,
				617,
				647,
				716,
				769,
				792,
			},
			5,
		},
		{
			dataSet,
			1362,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			ans := day1.Part2(tt.in)
			if ans != tt.expected {
				t.Errorf("got %d, want %d", ans, tt.expected)
			}
		})
	}
}
