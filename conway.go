package main

import "github.com/jqkard/fn"

func newConway(numRows, numCols, numCells int) Grid[bool] {
	grid := newGrid[bool](numRows, numCols)
	for _, idx := range randomNumbers(numRows*numCols, numCells) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = true
	}
	return grid
}

func displayConway(grid Grid[bool]) {
	displayGrid(grid, func(cell bool) string {
		return fn.Ternary(cell, "\u25A0", " ")
	})
}
