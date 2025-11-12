package main

import "fmt"

// 2048: 4x4, adds 2 randomly, with small chance of 4

func new2048() Grid[int] {
	// Make blank int grid
	numRows, numCols := 4, 4
	grid := newGrid[int](numRows, numCols)
	// Randomly place 2x 2 tiles
	for _, idx := range randomNumbers(numRows*numCols, 2) {
		row, col := indexToCoords(idx, numCols).Tuple()
		grid[row][col] = 2
	}
	return grid
}

func display2048(grid Grid[int]) {
	displayGrid(grid, func(cell int) string {
		if cell == 0 {
			return fmt.Sprintf("%5s", ".")
		}
		return fmt.Sprintf("%5d", cell)
	})
}
