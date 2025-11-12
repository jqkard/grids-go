package main

const (
	MINES = "mines"
	GREED = "greed"
)

func main() {
	option := GREED
	switch option {
	case MINES:
		grid := newMines(9, 9, 10)
		displayMines(grid)
	case GREED:
		grid := newGreed()
		displayGreed(grid)
	}

}
