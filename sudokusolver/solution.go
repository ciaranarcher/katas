package sudokusolver

import (
	"fmt"
)

// ValidateSolution will return true if a solution is valid.
func ValidateSolution(grid [][]int) bool {
	if !valid(grid) {
		return false
	}

	// We know we have a consistent grid now
	for i := 0; i < len(grid); i = i + 3 {
		for j := 0; j < len(grid); j = j + 3 {
			if !validSquare(grid, i, j) {
				return false
			}
		}
	}

	return true
}

func validSquare(grid [][]int, x, y int) bool {
	sum := 0
	for i := x; i < x+3; i = i + 1 {
		for j := y; j < y+3; j = j + 1 {
			sum = sum + grid[i][j]
		}
	}

	if sum == 45 {
		return true
	}

	return false
}

func valid(grid [][]int) bool {
	if len(grid) == 0 || len(grid)%3 != 0 {
		fmt.Println("incorrect grid size")
		return false
	}

	// A valid grid should have all rows the same length as the grid.
	// This will catch inconsistent length rows of all types.
	for _, row := range grid {
		if len(row) != len(grid) {
			fmt.Println("incorrect row size")
			return false
		}
	}

	return true
}
