package main

import (
	"github.com/jqkard/fn"
	"github.com/jqkard/fn/ds"
)

func newConway(numRows, numCols, numCells int) Grid[bool] {
	grid := newGrid[bool](numRows, numCols)
	for _, idx := range randomNumbers(numRows*numCols, numCells) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = true
	}
	return grid
}

func nextConway(grid Grid[bool]) Grid[bool] {
	numRows, numCols := shape(grid)
	grid2 := make(Grid[bool], numRows)
	isAlive := func(cell bool) bool {
		return cell
	}
	for row, line := range grid {
		line2 := make([]bool, numCols)
		for col, alive := range line {
			count := countNeighbors(grid, ds.Coords{row, col}, isAlive)
			alive2 := false
			if alive && count < 2 {
				alive2 = false
			} else if alive && (count == 2 || count == 3) {
				alive2 = true
			} else if alive && count > 3 {
				alive2 = false
			} else if !alive && count == 3 {
				alive2 = true
			}
			line2[col] = alive2
		}
		grid2[row] = line2
	}
	return grid2
}

func displayConway(grid Grid[bool]) {
	displayGrid(grid, func(cell bool) string {
		return fn.Ternary(cell, "\u25A0", " ")
	})
}
