# Go-Sudoku-Solver

Small Golang utility for solving Sudoku puzzles. This uses fairly typical back-stepping with a few different approaches (parallel, caching lowest options).

## Usage

See `cmd/consolesolver` for an example calling the solver directly.  See `cmd/suserver` for a simple server implementation.  Use an array of 81 integers to create the grid ordered left-to-right, top-to-bottom.  0 represents an empty square.  `pkg/solver/puzzles.go` has a few puzzles to test and there is a benchmark test.

The benchmarks are all over the place but in practice, it seems the typical puzzle solved in around 5ms.