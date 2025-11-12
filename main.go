package main

const (
	MINES  = "mines"
	GREED  = "greed"
	CONWAY = "conway"
)

func main() {
	option := CONWAY
	switch option {
	case MINES:
		grid := newMines(9, 9, 10)
		displayMines(grid)
	case GREED:
		grid := newGreed()
		displayGreed(grid)
	case CONWAY:
		grid := newConway(50, 250, 500)
		displayConway(grid)
	}
}
