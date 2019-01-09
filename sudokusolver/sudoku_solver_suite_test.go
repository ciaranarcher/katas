package sudokusolver_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSudokuSolver(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SudokuSolver Suite")
}
