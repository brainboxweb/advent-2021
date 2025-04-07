package day1

func Part1(data []int) int {
	increase := 0
	for k := 0; k < len(data)-1; k++ {
		if data[k+1] > data[k] {
			increase++
		}
	}
	return increase
}

func Part2(data []int) int {
	increase := 0
	for k := 0; k < len(data)-3; k++ {
		range1 := data[k] + data[k+1] + data[k+2]
		range2 := data[k+1] + data[k+2] + data[k+3]
		if range2 > range1 {
			increase++
		}
	}
	return increase
}
