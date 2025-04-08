package ocean_test

import (
	"testing"

	"github.com/brainboxweb/advent-2021/day6/ocean"
	"github.com/stretchr/testify/assert"
)

func TestFishes(t *testing.T) {
	tests := []struct {
		data     []int
		days     int
		expected int
	}{
		{
			[]int{3, 4, 3, 1, 2},
			1,
			5,
		},
		{
			[]int{3, 4, 3, 1, 2},
			2,
			6,
		},
		{
			[]int{3, 4, 3, 1, 2},
			3,
			7,
		},
		{
			[]int{3, 4, 3, 1, 2},
			4,
			9,
		},
		{
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			myFishes := ocean.NewFishes()
			for _, age := range tt.data {
				myFishes.AddFish(age)
			}
			total := myFishes.Move(tt.days)
			assert.Equal(t, tt.expected, total)
		})
	}
}

func TestAdvancedFishes(t *testing.T) {
	tests := []struct {
		data     []int
		days     int
		expected int
	}{
		{
			[]int{3, 4, 3, 1, 2},
			1,
			5,
		},
		{
			[]int{3, 4, 3, 1, 2},
			2,
			6,
		},
		{
			[]int{3, 4, 3, 1, 2},
			3,
			7,
		},
		{
			[]int{3, 4, 3, 1, 2},
			4,
			9,
		},
		{
			[]int{3, 4, 3, 1, 2},
			18,
			26,
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			myFishes := ocean.NewAdvancedFishes()
			for _, age := range tt.data {
				myFishes.AddFish(age)
			}
			total := myFishes.Move(tt.days)
			assert.Equal(t, tt.expected, total)
		})
	}
}
