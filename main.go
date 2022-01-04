package main

import (
	"fmt"
	"log"

	"github.com/ataboo/go-sudoku-solver/pkg/solver"
)

func main() {
	grid, err := solver.SudokuGridFromInts(solver.TestGrid5)
	if err != nil {
		log.Fatal("failed to parse grid:", err)
	}

	fmt.Printf("grid before:\n%s\n", grid.String())

	if err := solver.SolveGrid(grid); err != nil {
		log.Fatal("error solving grid:", err)
	}

	fmt.Printf("grid after: \n%s\n", grid.String())

}
