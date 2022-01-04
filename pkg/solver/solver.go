package solver

import (
	"fmt"
	"time"
)

var solvedChan chan *suGrid = make(chan *suGrid)

type SudokuGrid interface {
	///String representation of the Sudoku grid pipe-separated.
	String() string
	///GridState - Unsolved, Solved, or Invalid
	GridState() SolveState
	///Integer values in the grid.  In order of left to right, top to bottom.
	IntValues() []int
}

///Solve the SudokuGrid represented by the provided integers.
///Expects 81 values ordered left-to-right rows, top-to-bottom.
func SolveGrid(sudokuGrid SudokuGrid) error {
	grid, ok := sudokuGrid.(*suGrid)
	if !ok {
		return fmt.Errorf("invalid grid")
	}

	if err := grid.initSetMaps(); err != nil {
		return fmt.Errorf("duplicates found in the grid")
	}

	go gridSolvingRoutine(grid)

	select {
	case solvedGrid := <-solvedChan:
		copy(grid.numbers, solvedGrid.numbers)
		grid.gridState = solvedGrid.gridState
		return nil
	case <-time.After(time.Second * 10):
		return fmt.Errorf("the solver timed out")
	}
}

func gridSolvingRoutine(grid *suGrid) {
	for {
		if grid.solveByElimination() {
			continue
		}

		if grid.gridState == Solved {
			solvedChan <- grid
			return
		}

		if grid.gridState == Unsolved {
			forks := grid.forkAtLowestOptions()

			for _, fork := range forks {
				go gridSolvingRoutine(fork)
			}
		}

		return
	}
}
