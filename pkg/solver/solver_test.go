package solver

import "testing"

func BenchmarkGrid2WithGroups(b *testing.B) {
	grid, err := SudokuGridFromInts(TestGrid5)
	if err != nil {
		b.Error(err)
	}

	b.Logf("Grid Before: \n%s\n", grid.String())

	err = SolveGrid(grid)
	if err != nil {
		b.Error(err)
		return
	}

	b.Logf("Grid After: \n%s\n", grid.String())

	if grid.GridState() != Solved {
		b.Error("failed to solve grid", grid.String(), err)
	}
}
