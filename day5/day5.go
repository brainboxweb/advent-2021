package day5

import (
	"strconv"
	"strings"
)

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func Part1(data []string, size int) int {
	board := make([][]int, size)
	for q := range board {
		board[q] = make([]int, size)
	}
	for _, dat := range data {
		parts := strings.Split(dat, " ")
		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[2], ",")
		x1, err := strconv.Atoi(start[0])
		if err != nil {
			panic("not expected")
		}
		x2, err := strconv.Atoi(end[0])
		if err != nil {
			panic("not expected")
		}
		y1, err := strconv.Atoi(start[1])
		if err != nil {
			panic("not expected")
		}
		y2, err := strconv.Atoi(end[1])
		if err != nil {
			panic("not expected")
		}
		if isHoriz(x1, x2, y1, y2) {
			x11, x22 := swap(x1, x2)
			y11, y22 := swap(y1, y2)
			for i := x11; i < x22+1; i++ {
				for j := y11; j < y22+1; j++ {
					board[i][j]++
				}
			}
		}
	}
	counter := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] > 1 {
				counter++
			}
		}
	}

	return counter
}

func Part2(data []string, size int) int {
	board := make([][]int, size)
	for q := range board {
		board[q] = make([]int, size)
	}
	for _, dat := range data {
		parts := strings.Split(dat, " ")
		start := strings.Split(parts[0], ",")
		end := strings.Split(parts[2], ",")
		x1, err := strconv.Atoi(start[0])
		if err != nil {
			panic("not expected")
		}
		x2, err := strconv.Atoi(end[0])
		if err != nil {
			panic("not expected")
		}
		y1, err := strconv.Atoi(start[1])
		if err != nil {
			panic("not expected")
		}
		y2, err := strconv.Atoi(end[1])
		if err != nil {
			panic("not expected")
		}
		if isHoriz(x1, x2, y1, y2) {
			x11, x22 := swap(x1, x2)
			y11, y22 := swap(y1, y2)
			for i := x11; i < x22+1; i++ {
				for j := y11; j < y22+1; j++ {
					board[i][j]++
				}
			}
		} else if isDiag(x1, x2, y1, y2) {
			switch {
			case x1 < x2 && y1 < y2:
				j := y1
				for i := x1; i < x2+1; i++ {
					board[i][j]++
					j++
				}
			case x1 > x2 && y1 > y2:
				j := y1
				for i := x1; i > x2-1; i-- {
					board[i][j]++
					j--
				}
			case x1 < x2 && y1 > y2:
				j := y1
				for i := x1; i < x2+1; i++ {
					board[i][j]++
					j--
				}
			case x1 > x2 && y1 < y2:
				j := y1
				for i := x1; i > x2-1; i-- {
					board[i][j]++
					j++
				}
			}
		}
	}
	counter := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if board[i][j] > 1 {
				counter++
			}
		}
	}

	return counter
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

func swap(a, b int) (x, y int) {
	if a > b {
		return b, a
	}

	return a, b
}

func isHoriz(x1, x2, y1, y2 int) bool {
	if x1 == x2 || y1 == y2 {
		return true
	}
	return false
}

func isDiag(x1, x2, y1, y2 int) bool {
	if (x2-x1) == (y2-y1) || (x2-x1) == (y1-y2) {
		return true
	}
	return false
}
