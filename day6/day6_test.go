package day6_test

import (
	"fmt"
	"testing"

	"github.com/brainboxweb/advent-2021/day6"
	"github.com/brainboxweb/advent-2021/helpers"
	"github.com/stretchr/testify/assert"
)

const testDataPath = "../testdata"

func TestPart1(t *testing.T) {
	theString := "3,4,3,1,2"
	data := helpers.GetDataString(testDataPath + "/day6.txt")
	theBigString := data[0]
	tests := []struct {
		data     string
		days     int
		expected int
	}{
		{
			theString,
			1,
			5,
		},
		{
			theString,
			2,
			6,
		},
		{
			theString,
			3,
			7,
		},
		{
			theString,
			4,
			9,
		},
		{
			theString,
			18,
			26,
		},
		{
			theString,
			80,
			5934,
		},
		{
			theBigString,
			80,
			350605,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			result := day6.Part1(tt.data, tt.days)
			assert.Equal(t, tt.expected, result)
		})
	}
}

func TestPart2(t *testing.T) {
	theString := "3,4,3,1,2"
	data := helpers.GetDataString(testDataPath + "/day6.txt")
	theBigString := data[0]
	tests := []struct {
		data     string
		days     int
		expected int
	}{
		{
			theString,
			1,
			5,
		},
		{
			theString,
			2,
			6,
		},
		{
			theString,
			3,
			7,
		},
		{
			theString,
			4,
			9,
		},
		{
			theString,
			10,
			12,
		},
		{
			theString,
			18,
			26,
		},
		{
			theString,
			80,
			5934,
		},
		{
			theBigString,
			256,
			1592778185024,
		},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("Day %d ", tt.days), func(t *testing.T) {
			result := day6.Part2(tt.data, tt.days)
			assert.Equal(t, tt.expected, result)
		})
	}
}
