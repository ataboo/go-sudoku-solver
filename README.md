# Go-Sudoku-Solver

Small Golang utility for solving Sudoku puzzles. This uses fairly typical back-stepping with a few different approaches (parallel, caching lowest options).

## Usage

See `cmd/consolesolver` for an example calling the solver directly.  See `cmd/suserver` for a simple server implementation.  Use an array of 81 integers to create the grid ordered left-to-right, top-to-bottom.  0 represents an empty square.  `pkg/solver/puzzles.go` has a few puzzles to test and there is a benchmark test.

My PC is averaging ~120ms consistently regardless of puzzle difficulty.