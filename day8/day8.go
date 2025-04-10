package day8

import (
	"strings"

	"github.com/brainboxweb/advent-2021/day8/display"
)

func Part1(data []string) int {
	chunkData := parse(data)
	ret := 0
	for _, cData := range chunkData {
		ret += countDigitsByChuckLength(cData.out)
	}
	return ret
}

func Part2(data []string) int {
	chunkData := parse(data)
	ret := 0
	for _, cData := range chunkData {
		disp := display.New()
		disp.Calibrate(cData.in)
		result := disp.Decode(cData.out)
		ret += result
	}
	return ret
}

func countDigitsByChuckLength(chunks []string) int {
	count := 0
	for _, chunk := range chunks {
		switch len(chunk) {
		case 2:
			count++
		case 3:
			count++
		case 4:
			count++
		case 7:
			count++
		default:
			// ignore
		}
	}
	return count
}

type chunkData struct {
	in  []string
	out []string
}

func parse(data []string) []chunkData {
	ret := []chunkData{}
	for _, line := range data {
		inout := strings.Split(line, " | ")
		inString := inout[0]
		inChunks := strings.Split(inString, " ")
		outString := inout[1]
		outChunks := strings.Split(outString, " ")
		cd := chunkData{inChunks, outChunks}
		ret = append(ret, cd)
	}
	return ret
}
