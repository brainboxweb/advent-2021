package day2

import (
	"strconv"
	"strings"
)

func Part1(data []string) int {
	x := 0
	y := 0
	for _, str := range data {
		parts := strings.Split(str, " ")
		direction := parts[0]
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("not expected")
		}
		switch direction {
		case "forward":
			x += size
		case "down": // NB
			y += size
		case "up":
			y -= size
		}
	}
	return x * y
}

func Part2(data []string) int {
	aim := 0
	x := 0
	y := 0
	for _, str := range data {
		parts := strings.Split(str, " ")
		direction := parts[0]
		size, err := strconv.Atoi(parts[1])
		if err != nil {
			panic("not expected")
		}
		switch direction {
		case "forward":
			x += size
			// It increases your depth by your aim multiplied by X.
			y += aim * size
		case "down": // NB
			aim += size
		case "up":
			aim -= size
		}
	}
	return x * y
}
