package sudoku_solver

import "fmt"

func ValidateSolution(grid [][]int) bool {
	if !valid(grid) {
		return false
	}

	return true
}

func valid(grid [][]int) bool {
	prevRowLen := 0
	if len(grid) == 0 || len(grid)%3 != 0 {
		fmt.Println("incorrect grid size")
		return false
	}

	for _, row := range grid {
		if len(row) == 0 || len(row)%3 != 0 {
			fmt.Println("incorrect row size")
			return false
		}

		if prevRowLen != 0 && len(row) != prevRowLen {
			fmt.Println("inconsistent row length")
			return false
		}

		prevRowLen = len(row)
	}

	return true
}
