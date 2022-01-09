package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ataboo/go-sudoku-solver/pkg/solver"
)

var singleSolutionGrid = []int{
	0, 0, 2, 0, 0, 0, 0, 6, 0,
	5, 6, 0, 3, 0, 0, 0, 0, 7,
	0, 0, 8, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 8,
	6, 3, 0, 0, 0, 9, 0, 1, 0,
	0, 2, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 7, 0, 0, 4, 0, 0,
	9, 1, 0, 0, 0, 3, 0, 8, 0,
	0, 0, 5, 0, 0, 0, 0, 0, 0,
}

var multiSolutionGrid = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 0, 0, 0, 0, 9, 0,
	0, 0, 9, 0, 0, 0, 4, 0, 0,
	8, 0, 0, 9, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 7, 0, 0, 0, 0,
	0, 6, 0, 0, 0, 0, 0, 0, 3,
	7, 0, 0, 0, 4, 0, 0, 0, 0,
	0, 8, 0, 2, 0, 7, 6, 0, 0,
	0, 0, 0, 0, 5, 0, 0, 2, 0,
}

var noSolutionGrid = []int{
	1, 0, 2, 0, 0, 0, 0, 6, 0,
	5, 6, 0, 3, 0, 0, 0, 0, 7,
	0, 0, 8, 0, 0, 5, 0, 0, 0,
	0, 0, 0, 0, 1, 0, 0, 0, 8,
	6, 3, 0, 0, 0, 9, 0, 1, 0,
	0, 2, 0, 0, 0, 0, 0, 0, 0,
	0, 0, 0, 7, 0, 0, 4, 0, 0,
	9, 1, 0, 0, 0, 3, 0, 8, 0,
	0, 0, 5, 0, 0, 0, 0, 0, 0,
}

func main() {
	grid, err := solver.SudokuGridFromInts(singleSolutionGrid)
	if err != nil {
		log.Fatal("failed to parse grid:", err)
	}

	config := solver.SolverConfig{
		SolutionLimit: 1,
		Timeout:       time.Second,
	}

	fmt.Println("Puzzle with Single Solution:")
	solveGridAndPrintResult(grid, config)

	grid, err = solver.SudokuGridFromInts(multiSolutionGrid)
	if err != nil {
		log.Fatal("failed to parse grid:", err)
	}

	config.SolutionLimit = 50
	fmt.Println("Puzzle with Multiple Solutions:")
	solveGridAndPrintResult(grid, config)

	grid, err = solver.SudokuGridFromInts(noSolutionGrid)
	if err != nil {
		log.Fatal("failed to parse grid:", err)
	}
	fmt.Println("Puzzle with No Solutions:")
	solveGridAndPrintResult(grid, config)

}

func solveGridAndPrintResult(grid solver.SudokuGrid, config solver.SolverConfig) {
	fmt.Printf("Puzzle before:\n%s\n", grid.String())

	startTime := time.Now()
	result, err := solver.SolveGridWithConfig(grid, config)
	if err != nil {
		log.Fatal("error solving grid:", err)
	}
	solveTime := time.Since(startTime)

	fmt.Printf("Solve time: %dμs, Forks: %d, Deadends: %d, Timedout: %d, Solutions: %d\n\n", solveTime.Microseconds(), result.ForkCount, result.DeadEndCount, result.TimedOut, len(result.SolvedGrids))

	if len(result.SolvedGrids) > 0 {
		gridShowCount := 3

		for i, g := range result.SolvedGrids {
			if i >= gridShowCount {
				break
			}
			fmt.Printf("Solution%d\n%s\n", i+1, g.String())
		}
	}

	fmt.Print("\n")
}
