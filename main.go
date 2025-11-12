package main

const (
	MINES  = "mines"
	GREED  = "greed"
	CONWAY = "conway"
	T2048  = "2048"
)

func main() {
	option := T2048
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
	case T2048:
		grid := new2048()
		display2048(grid)
	}
}
