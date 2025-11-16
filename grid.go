package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"os/exec"
	"strings"

	"github.com/jqkard/fn/ds"
)

type Grid[T any] = [][]T

func newGrid[T any](numRows, numCols int) Grid[T] {
	grid := make(Grid[T], numRows)
	for i := range numRows {
		grid[i] = make([]T, numCols)
	}
	return grid
}

func shape[T any](grid Grid[T]) (int, int) {
	return len(grid), len(grid[0])
}

func indexToCoords(idx, width int) ds.Coords {
	row, col := idx/width, idx%width
	return ds.Coords{row, col}
}

func countNeighbors[T any](grid Grid[T], coords ds.Coords, ok func(T) bool) int {
	numRows, numCols := shape(grid)
	y, x := coords.Tuple()
	count := 0
	deltas := []int{-1, 0, 1}
	for _, dy := range deltas {
		for _, dx := range deltas {
			if dy == 0 && dx == 0 {
				continue
			}
			ny, nx := y+dy, x+dx
			if notInsideBounds(ny, nx, numRows, numCols) {
				continue
			}
			if ok(grid[ny][nx]) {
				count += 1
			}
		}
	}
	return count
}

func notInsideBounds(y, x, numRows, numCols int) bool {
	return y < 0 || x < 0 || y >= numRows || x >= numCols
}

func randomNumbers(limit, count int) []int {
	numbers := ds.NewSet[int]()
	for numbers.Len() != count {
		numbers.Add(rand.IntN(limit))
	}
	return numbers.Items()
}

func displayGrid[T any](grid Grid[T], toString func(T) string) {
	for _, row := range grid {
		line := make([]string, len(row))
		for i, cell := range row {
			line[i] = toString(cell)
		}
		fmt.Println(strings.Join(line, ""))
	}
}

func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
