package day4

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2021/day4/bingo"
)

func Part1(data []string) int {
	calls, boardsData := parse(data)
	boards := []bingo.Board{}
	for _, boardData := range boardsData {
		board := bingo.NewBoard()
		board.AddRows(boardData)
		boards = append(boards, board)
	}

	for _, call := range calls {
		for _, board := range boards {
			calledValue, score := board.Call(call)
			if calledValue > 0 {
				return calledValue * score
			}
		}
	}

	return 0 // not expected
}

//revive:disable:cognitive-complexity

func Part2(data []string) int {
	calls, boardsData := parse(data)
	boards := []bingo.Board{}
	for _, boardData := range boardsData {
		board := bingo.NewBoard()
		board.AddRows(boardData)
		boards = append(boards, board)
	}
	winners := make(map[int]bool)
	for i := 1; i < len(boards)+1; i++ {
		winners[i] = false
	}
	for _, call := range calls {
		for boardNumber, board := range boards {
			calledValue, score := board.Call(call)
			if score > 0 {
				winners[boardNumber] = true
				winnerCount := 0
				for _, win := range winners {
					if win {
						winnerCount++
					}
				}
				if winnerCount == len(boards) {
					return calledValue * score
				}
			}
		}
	}

	return 0 // not expected
}

//revive:enable:cognitive-complexity

func parse(data []string) ([]int, map[int][][]int) {
	callsData := getCallData(data[0])
	boardData := getBoardData(data[2:])

	return callsData, boardData
}

//revive:disable:cognitive-complexity

func getBoardData(boardData []string) map[int][][]int {
	chunks := [][]string{}
	chunk := []string{}
	for _, str := range boardData {
		if str == "" {
			chunks = append(chunks, chunk) // store
			chunk = []string{}             // start again
			continue
		}
		chunk = append(chunk, str)
	}
	chunks = append(chunks, chunk) // finsih off
	ret := make(map[int][][]int)
	counter := 0
	for _, chunk := range chunks {
		rows := [][]int{}
		for _, rowData := range chunk {
			row := []int{}
			parts := strings.Split(rowData, " ")
			for _, part := range parts {
				if part == "" {
					continue
				}
				valInt, err := strconv.Atoi(part)
				if err != nil {
					panic("not expected")
				}
				row = append(row, valInt)
			}
			rows = append(rows, row)
		}
		ret[counter] = rows
		counter++
	}

	return ret
}

//revive:enable:cognitive-complexity

func getCallData(callData string) []int {
	callDataString := strings.Split(callData, ",")
	calls := []int{}
	for _, val := range callDataString {
		call, err := strconv.Atoi(val)
		if err != nil {
			panic("not expected")
		}
		calls = append(calls, call)
	}

	return calls
}
