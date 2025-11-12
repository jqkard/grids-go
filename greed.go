package main

import (
	"fmt"
	"slices"

	"github.com/jqkard/fn/list"
)

const current int = -1

func newGreed() Grid[int] {
	// Make blank int grid
	numRows, numCols := 22, 79
	grid := newGrid[int](numRows, numCols)
	// Randomize the grid digits
	perDigit := 193
	digits := make([]int, 0, numRows*numCols)
	digits = append(digits, current)
	for i := 1; i <= 9; i++ {
		digits = append(digits, slices.Repeat([]int{i}, perDigit)...)
	}
	list.Shuffle(digits)
	for idx, digit := range digits {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = digit
	}
	return grid
}

func displayGreed(grid Grid[int]) {
	displayGrid(grid, func(cell int) string {
		switch cell {
		case 0:
			return fmt.Sprintf("%2s", ".")
		case current:
			return fmt.Sprintf("%2s", "@")
		}
		return fmt.Sprintf("%2d", cell)
	})
}
