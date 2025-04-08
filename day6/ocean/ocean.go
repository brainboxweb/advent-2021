package ocean

type Ocean interface {
	AddFish(age int)
	Breed(dayCount int) int
}
