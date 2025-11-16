package main

import "time"

const (
	MINES  = "mines"
	GREED  = "greed"
	CONWAY = "conway"
	T2048  = "2048"
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
		runLoop(grid, displayConway, nextConway, 100)
	case T2048:
		grid := new2048()
		display2048(grid)
	}
}

func runLoop[T any](grid Grid[T], displayFn func(Grid[T]), nextFn func(Grid[T]) Grid[T], delayMs int) {
	for {
		clearScreen()
		displayFn(grid)
		if delayMs > 0 {
			time.Sleep(time.Duration(delayMs) * time.Millisecond)
		}
		grid = nextFn(grid)
	}
}
