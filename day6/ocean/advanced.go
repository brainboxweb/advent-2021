package ocean

func NewAdvancedFishes() *AdvancedFishes {
	fishMap := make(map[int]int)
	return &AdvancedFishes{fishMap: fishMap}
}

type AdvancedFishes struct {
	fishMap map[int]int // age, count
}

func (f *AdvancedFishes) AddFish(age int) {
	f.fishMap[age]++
}

func (f *AdvancedFishes) Move(days int) int {
	for i := 0; i < days; i++ {
		f.moveOnce()
	}
	return f.Count()
}

// this time, we're changing ALL the fish of each age.
func (f *AdvancedFishes) moveOnce() {
	// Start with the zeros
	zeroCount := f.fishMap[0]
	f.fishMap[0] = 0               // they all go away
	for age := 1; age < 9; age++ { // check max value
		// increment the counter down across the board
		f.fishMap[age-1] += f.fishMap[age]
		f.fishMap[age] = 0
	}
	// reset the 0's to 6
	f.fishMap[6] += zeroCount
	// spawn new 8's
	f.fishMap[8] += zeroCount
}

func (f *AdvancedFishes) Count() int {
	ret := 0
	for _, cnt := range f.fishMap {
		ret += cnt
	}
	return ret
}
