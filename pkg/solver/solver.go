package solver

import (
	"fmt"
	"time"
)

type SolverConfig struct {
	SolutionLimit int
	Timeout       time.Duration
	RunParallel   bool
}

type SudokuGrid interface {
	///String representation of the Sudoku grid pipe-separated.
	String() string
	///GridState - Unsolved, Solved, or Invalid
	GridState() SolveState
	///Integer values in the grid.  In order of left to right, top to bottom.
	IntValues() []int
}

type SolverResult struct {
	SolvedGrids  []SudokuGrid
	ForkCount    int
	DeadEndCount int
	TimedOut     int
}

type sudokuSolver struct {
	solvedChan      chan *suGrid
	forkChan        chan bool
	deadChan        chan bool
	activeGridCount int
	result          *SolverResult
	config          SolverConfig
}

///Solve the SudokuGrid represented by the provided integers.
///Expects 81 values ordered left-to-right rows, top-to-bottom.
func SolveGrid(sudokuGrid SudokuGrid) (*SolverResult, error) {
	config := SolverConfig{
		SolutionLimit: 1,
		Timeout:       1 * time.Second,
		RunParallel:   true,
	}

	return SolveGridWithConfig(sudokuGrid, config)
}

func SolveGridWithConfig(sudokuGrid SudokuGrid, config SolverConfig) (*SolverResult, error) {
	suSolver := &sudokuSolver{
		solvedChan:      make(chan *suGrid),
		forkChan:        make(chan bool),
		deadChan:        make(chan bool),
		activeGridCount: 1,
		result: &SolverResult{
			SolvedGrids:  make([]SudokuGrid, 0),
			ForkCount:    0,
			DeadEndCount: 0,
			TimedOut:     0,
		},
		config: config,
	}

	grid, ok := sudokuGrid.(*suGrid)
	if !ok {
		return nil, fmt.Errorf("invalid grid")
	}

	if err := grid.initSetMaps(); err != nil {
		return nil, fmt.Errorf("duplicates found in the grid")
	}

	go gridSolvingRoutine(grid, suSolver)

	suSolver.run(grid)

	return suSolver.result, nil
}

func (s *sudokuSolver) run(grid *suGrid) {
	timeout := make(<-chan time.Time)
	if s.config.Timeout > 0 {
		timeout = time.After(s.config.Timeout)
	}

	for {
		select {
		case solvedGrid := <-s.solvedChan:
			s.result.SolvedGrids = append(s.result.SolvedGrids, solvedGrid)
			s.activeGridCount--
		case <-s.forkChan:
			s.result.ForkCount++
			s.activeGridCount++
		case <-s.deadChan:
			s.activeGridCount--
			s.result.DeadEndCount++
		case <-timeout:
			s.result.TimedOut = s.activeGridCount
			return
		}

		if s.activeGridCount == 0 {
			return
		}

		if s.config.SolutionLimit > 0 && len(s.result.SolvedGrids) >= s.config.SolutionLimit {
			return
		}
	}
}

func gridSolvingRoutine(grid *suGrid, solver *sudokuSolver) {
	for {
		if grid.solveByElimination() {
			continue
		}

		if grid.gridState == Solved {
			solver.solvedChan <- grid
			return
		}

		if grid.gridState == Unsolved {
			forks := grid.forkAtLowestOptions()

			for _, fork := range forks {
				solver.forkChan <- true
				if solver.config.RunParallel {
					go gridSolvingRoutine(fork, solver)
				} else {
					gridSolvingRoutine(fork, solver)
				}
			}
		}

		solver.deadChan <- true

		return
	}
}
