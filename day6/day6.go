package day6

import (
	"strconv"
	"strings"

	"github.com/brainboxweb/advent-2021/day6/ocean"
)

func Part1(data string, days int) int {
	myfishes := ocean.Fishes{}
	ages := parse(data)
	for _, age := range ages {
		myfishes.AddFish(age)
	}
	return myfishes.Move(days)
}

func Part2(data string, days int) int {
	myfishes := ocean.NewAdvancedFishes()
	ages := parse(data)
	for _, age := range ages {
		myfishes.AddFish(age)
	}
	return myfishes.Move(days)
}

func parse(data string) []int {
	ages := []int{}
	fishesString := strings.Split(data, ",")
	for _, dat := range fishesString {
		age, err := strconv.Atoi(dat)
		if err != nil {
			panic("not expected")
		}
		ages = append(ages, age)
	}
	return ages
}
