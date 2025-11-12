package main

import (
	"fmt"

	"github.com/jqkard/fn/ds"
)

// Easy		9x9		10
// Medium	16x16	40
// Hard 	30x16	99

const bomb int = -1

func newMines(numRows, numCols, numMines int) Grid[int] {
	// Make blank int grid
	grid := newGrid[int](numRows, numCols)
	// Randomly place bombs
	for _, idx := range randomNumbers(numRows*numCols, numMines) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = bomb
	}
	isBomb := func(cell int) bool {
		return cell == bomb
	}
	// Count neighbor mines
	for row, line := range grid {
		for col, cell := range line {
			if cell == bomb {
				continue
			}
			grid[row][col] = countNeighbors(grid, ds.Coords{row, col}, isBomb)
		}
	}
	return grid
}

func displayMines(grid Grid[int]) {
	displayGrid(grid, func(cell int) string {
		if cell == bomb {
			return fmt.Sprintf("%3s", "X")
		}
		return fmt.Sprintf("%3d", cell)
	})
}
