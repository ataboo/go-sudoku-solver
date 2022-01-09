package solver

import (
	"testing"
	"time"
)

var multiSolutionGrid = []int{
	0, 0, 0, 0, 0, 0, 0, 0, 0,
	1, 0, 0, 8, 0, 2, 0, 9, 0,
	0, 0, 9, 0, 0, 0, 4, 0, 0,
	8, 0, 0, 9, 0, 1, 0, 4, 0,
	0, 0, 0, 0, 7, 0, 0, 0, 0,
	0, 6, 0, 0, 0, 0, 0, 0, 3,
	7, 0, 0, 0, 4, 0, 0, 0, 0,
	0, 8, 0, 2, 0, 7, 6, 0, 0,
	0, 0, 0, 0, 5, 0, 0, 2, 0,
}

func BenchmarkParallelGrid4(b *testing.B) {
	grid, err := SudokuGridFromInts(TestGrid4)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 1,
		Timeout:       time.Second * 10,
		RunParallel:   true,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) == 0 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid", grid.String(), err)
	}
}

func BenchmarkParallelGrid5(b *testing.B) {
	grid, err := SudokuGridFromInts(TestGrid5)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 1,
		Timeout:       time.Second * 10,
		RunParallel:   true,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) == 0 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid", grid.String(), err)
	}
}

func BenchmarkParallelMultiSolutionGrid(b *testing.B) {
	grid, err := SudokuGridFromInts(multiSolutionGrid)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 10,
		Timeout:       10 * time.Second,
		RunParallel:   true,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) != 10 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid")
	}
}

func BenchmarkNonParallelGrid4(b *testing.B) {
	grid, err := SudokuGridFromInts(TestGrid4)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 1,
		Timeout:       time.Second * 10,
		RunParallel:   false,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) == 0 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid", grid.String(), err)
	}
}

func BenchmarkNonParallelGrid5(b *testing.B) {
	grid, err := SudokuGridFromInts(TestGrid5)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 1,
		Timeout:       time.Second * 10,
		RunParallel:   false,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) == 0 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid", grid.String(), err)
	}
}

func BenchmarkNonParallelMultiSolutionGrid(b *testing.B) {
	grid, err := SudokuGridFromInts(multiSolutionGrid)
	if err != nil {
		b.Error(err)
	}

	config := SolverConfig{
		SolutionLimit: 10,
		Timeout:       10 * time.Second,
		RunParallel:   false,
	}

	result, err := SolveGridWithConfig(grid, config)
	if err != nil {
		b.Error(err)
		return
	}

	if len(result.SolvedGrids) != 10 || result.SolvedGrids[0].GridState() != Solved {
		b.Error("failed to solve grid")
	}
}
