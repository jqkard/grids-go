package main

import (
	"fmt"
	"strings"

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
	for _, row := range grid {
		line := make([]string, len(row))
		for i, cell := range row {
			if cell == bomb {
				line[i] = fmt.Sprintf("%3s", "X")
			} else {
				line[i] = fmt.Sprintf("%3d", cell)
			}
		}
		fmt.Println(strings.Join(line, ""))
	}
}
