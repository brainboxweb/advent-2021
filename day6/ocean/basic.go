package ocean

func NewFishes() *Fishes {
	return &Fishes{}
}

type Fishes struct {
	fishes []fish
}

type fish struct {
	age int
}

func (f *Fishes) AddFish(age int) {
	newFish := fish{age: age}
	f.fishes = append(f.fishes, newFish)
}

func (f *Fishes) Move(days int) int {
	for i := 0; i < days; i++ {
		f.moveOnce()
	}
	return f.Count()
}

func (f *Fishes) moveOnce() {
	birthers := 0
	for index, fish := range f.fishes {
		currentAge := fish.age
		if currentAge == 0 {
			f.fishes[index].age = 6 // reset
			birthers++
		} else {
			f.fishes[index].age = fish.age - 1
		}
	}
	for i := 0; i < birthers; i++ {
		f.AddFish(8)
	}
}

func (f *Fishes) Count() int {
	return len(f.fishes)
}
