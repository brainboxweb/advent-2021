package bingo

type Board struct {
	rows   [][]int
	called [][]int
	done   bool
}

func NewBoard() Board { // Rename
	b := Board{}
	a := make([][]int, 5)
	for i := range a {
		a[i] = make([]int, 5)
	}
	b.called = a

	return b
}

func (b *Board) AddRows(row [][]int) {
	b.rows = row
}

//revive:disable:cognitive-complexity
//revive:disable:cyclomatic

func (b *Board) Call(val int) (match int, score int) {
	if b.done {
		return 0, 0
	}
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.rows[x][y] == val {
				b.called[x][y] = 1
			}
		}
	}
	// checkrows
	for y := 0; y < 5; y++ {
		xsum := 0
		for x := 0; x < 5; x++ {
			xsum += b.called[x][y]
		}
		if xsum > 4 {
			b.done = true
			return val, b.GetScore()
		}
	}
	// checkcolumns
	for x := 0; x < 5; x++ {
		ysum := 0
		for y := 0; y < 5; y++ {
			ysum += b.called[x][y]
		}
		if ysum > 4 {
			b.done = true
			return val, b.GetScore()
		}
	}
	return 0, 0
}

//revive:enable:cognitive-complexity
//revive:enable:cyclomatic

func (b *Board) GetScore() int {
	score := 0
	for y := 0; y < 5; y++ {
		for x := 0; x < 5; x++ {
			if b.called[x][y] == 0 {
				score += b.rows[x][y]
			}
		}
	}
	return score
}
