package sudokusolver_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/ciaranarcher/katas/sudokusolver"
)

var testTrue = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 5, 3, 4, 8},
	{1, 9, 8, 3, 4, 2, 5, 6, 7},
	{8, 5, 9, 7, 6, 1, 4, 2, 3},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 6, 1, 5, 3, 7, 2, 8, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 4, 5, 2, 8, 6, 1, 7, 9},
}

var testFalse = [][]int{
	{5, 3, 4, 6, 7, 8, 9, 1, 2},
	{6, 7, 2, 1, 9, 0, 3, 4, 8},
	{1, 0, 0, 3, 4, 2, 5, 6, 0},
	{8, 5, 9, 7, 6, 1, 0, 2, 0},
	{4, 2, 6, 8, 5, 3, 7, 9, 1},
	{7, 1, 3, 9, 2, 4, 8, 5, 6},
	{9, 0, 1, 5, 3, 7, 2, 1, 4},
	{2, 8, 7, 4, 1, 9, 6, 3, 5},
	{3, 0, 0, 4, 8, 1, 1, 7, 9},
}

var _ = Describe("Test Solution", func() {
	It("returns the correct value", func() {
		Expect(ValidateSolution(testTrue)).To(Equal(true))
		Expect(ValidateSolution(testFalse)).To(Equal(false))
	})

})

var _ = Describe("a valid grid", func() {
	It("rejects empty arrays", func() {
		Expect(ValidateSolution([][]int{})).To(Equal(false))
	})

	It("rejects arrays where length is not a factor of 3", func() {
		gridWrongRows := [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 0, 3, 4, 8},
			{1, 0, 0, 3, 4, 2, 5, 6, 0},
			{8, 5, 9, 7, 6, 1, 0, 2, 0},
		}

		gridWrongCols := [][]int{
			{5, 3, 4, 6, 7, 8, 9},
			{6, 7, 2, 1, 9, 0, 3},
			{1, 0, 0, 3, 4, 2, 5},
		}

		Expect(ValidateSolution(gridWrongRows)).To(Equal(false))
		Expect(ValidateSolution(gridWrongCols)).To(Equal(false))
	})

	It("rejects arrays where rows are of differing length", func() {
		gridDifferingLengthRows := [][]int{
			{5, 3, 4, 6, 7, 8, 9, 1, 2},
			{6, 7, 2, 1, 9, 0},
			{1, 0, 0, 3, 4, 2, 5, 6, 0},
		}

		Expect(ValidateSolution(gridDifferingLengthRows)).To(Equal(false))
	})
})
