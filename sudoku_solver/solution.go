package sudoku_solver

import "fmt"

func ValidateSolution(grid [][]int) bool {
	if !valid(grid) {
		return false
	}

	return true
}

func valid(grid [][]int) bool {
	if len(grid) == 0 || len(grid)%3 != 0 {
		fmt.Println("incorrect grid size")
		return false
	}

	for _, row := range grid {
		if len(row) == 0 || len(row)%3 != 0 {
			fmt.Println("incorrect row size")
			return false
		}
	}

	return true
}
