package day3

import "math"

func Part1(data []string) int {
	gamma := make(map[int]int)
	epsilon := make(map[int]int)
	for _, str := range data {
		for i, char := range str {
			if string(char) == "1" {
				gamma[i]++
				epsilon[i]--
			} else {
				gamma[i]--
				epsilon[i]++
			}
		}
	}
	return getDecimal(gamma) * getDecimal(epsilon)
}

func Part2(data []string) int {
	oxygen := getStuff(data, 1)
	scrubber := getStuff(data, 0)

	return oxygen * scrubber
}

func getDecimal(dat map[int]int) int {
	binarySlice := []int{}
	for i := 0; i < len(dat); i++ {
		switch {
		case dat[i] > 0:
			binarySlice = append(binarySlice, 1)
		case dat[i] < 0:
			binarySlice = append(binarySlice, 0)
		default:
			panic("not expected")
		}
	}
	return decFromSlice(binarySlice)
}

func decFromSlice(binarySlice []int) int {
	power := 0.0
	dec := 0
	for i := len(binarySlice) - 1; i > -1; i-- {
		dec += binarySlice[i] * int(math.Pow(2, power))
		power++
	}
	return dec
}

func getStuff(data []string, pick int) int {
	loopData := data
	for i := 0; i < len(data[0]); i++ { // column at a time
		gamma := getGamma(loopData, i)
		switch {
		case gamma > 0: // keep 1s
			loopData = collect(loopData, pick, i)
		case gamma < 0: // keep 0s
			loopData = collect(loopData, 1-pick, i)
		default:
			loopData = collect(loopData, pick, i)
		}
		if len(loopData) == 1 {
			break
		}
	}
	str := loopData[0] // should be just one
	theSlice := []int{}
	for _, char := range str {
		if string(char) == "1" {
			theSlice = append(theSlice, 1)
		} else {
			theSlice = append(theSlice, 0)
		}
	}

	return decFromSlice(theSlice)
}

func getGamma(data []string, index int) int {
	gamma := 0
	for _, str := range data {
		char := str[index]
		if string(char) == "1" {
			gamma++
		} else {
			gamma--
		}
	}
	retValue := 0
	switch {
	case gamma > 0:
		retValue = 1
	case gamma < 0:
		retValue = -1
	case gamma == 0:
		retValue = 0
	default:
		panic("not expected")
	}
	return retValue
}

func collect(data []string, match int, index int) []string {
	stringMatch := "1"
	if match == 0 {
		stringMatch = "0"
	}
	output := []string{}
	for _, str := range data {
		if string(str[index]) == stringMatch {
			output = append(output, str)
			continue
		}
	}
	return output
}
