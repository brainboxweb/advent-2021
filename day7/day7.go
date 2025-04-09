package day7

import (
	"slices"
)

func Part1(positions []int) int {
	var costs []int
	for i := 1; i < len(positions)+1; i++ { // starts at one
		totalCost := 0
		for _, p := range positions {
			totalCost += diff(p, i)
		}
		costs = append(costs, totalCost)
	}
	slices.Sort(costs)
	return costs[0]
}

func Part2(positions []int) int {
	var costs []int
	for i := 1; i < len(positions)+1; i++ { // starts at one
		totalCost := 0
		for _, p := range positions {
			diff := diff(p, i)
			totalCost += gaussSum(diff)
		}
		costs = append(costs, totalCost)
	}
	slices.Sort(costs)
	return costs[0]
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func gaussSum(n int) int {
	return n * (n + 1) / 2
}
