package solver

import (
	"fmt"
	"strconv"
	"strings"
)

type SolveState int

const (
	Unsolved SolveState = 0
	Solved   SolveState = 1
	Invalid  SolveState = 2
)

type suGrid struct {
	numbers    []int
	rowSetMaps []map[int]bool
	colSetMaps []map[int]bool
	boxSetMaps []map[int]bool

	gridState SolveState
	rows      [][]*int
	cols      [][]*int
	boxes     [][]*int

	lowestOpenIndex   int
	lowestOpenNumbers []int
}

func (g *suGrid) String() string {
	b := strings.Builder{}
	b.WriteString(strings.Repeat("_", 19))
	b.WriteRune('\n')
	for y := 0; y < 9; y++ {
		row := g.rows[y]

		b.WriteRune('|')
		for x := 0; x < 9; x++ {
			b.WriteString(strconv.Itoa(*row[x]))
			b.WriteRune('|')
		}
		b.WriteRune('\n')
	}

	b.WriteString(strings.Repeat("_", 19))
	b.WriteRune('\n')

	return b.String()
}

func (g *suGrid) GridState() SolveState {
	return g.gridState
}

func (g *suGrid) IntValues() []int {
	numbersCpy := make([]int, 81)
	copy(numbersCpy, g.numbers)

	return numbersCpy
}

///suGridFromInts - Initialize a grid using a slice of integers.
///Ordered by left-to-right rows, top to bottom. 0 represents an empty square.
func SudokuGridFromInts(numbers []int) (SudokuGrid, error) {
	if len(numbers) != 81 {
		return nil, fmt.Errorf("int grid must be of length 81")
	}

	for _, v := range numbers {
		if v < 0 || v > 9 {
			return nil, fmt.Errorf("numbers must be between 0 and 9")
		}
	}

	grid := newSuGrid()
	copy(grid.numbers, numbers)

	return grid, nil
}

///newSuGrid - Constructor for suGrid.
func newSuGrid() *suGrid {
	grid := suGrid{
		numbers: make([]int, 81),

		//SetMaps keep track of which numbers have been set in their given domain.
		//i.e Each row has a map that is set true for every number that exists in that row.
		rowSetMaps: make([]map[int]bool, 9),
		colSetMaps: make([]map[int]bool, 9),
		boxSetMaps: make([]map[int]bool, 9),

		gridState: Unsolved,
		rows:      make([][]*int, 9),
		cols:      make([][]*int, 9),
		boxes:     make([][]*int, 9),
	}

	for i := 0; i < 9; i++ {
		grid.rows[i] = grid.sliceRow(i)
		grid.cols[i] = grid.sliceColumn(i)
		grid.boxes[i] = grid.sliceBox(i)
	}

	return &grid
}

///copy - Make a deep copy of the grid.
func (g *suGrid) copy() (*suGrid, error) {
	sudokuGrid, err := SudokuGridFromInts(g.numbers)
	if err != nil {
		return nil, err
	}

	grid, ok := sudokuGrid.(*suGrid)
	if !ok {
		return nil, fmt.Errorf("invalid grid type")
	}

	for i := 0; i < 9; i++ {
		grid.rowSetMaps[i] = make(map[int]bool)
		copyMap(g.rowSetMaps[i], grid.rowSetMaps[i])

		grid.colSetMaps[i] = make(map[int]bool)
		copyMap(g.colSetMaps[i], grid.colSetMaps[i])

		grid.boxSetMaps[i] = make(map[int]bool)
		copyMap(g.boxSetMaps[i], grid.boxSetMaps[i])
	}

	return grid, nil
}

///sliceRow - Create a slice of pointers referencing the 9 numbers in the row at the given index.
func (g *suGrid) sliceRow(idx int) []*int {
	if idx < 0 || idx > 8 {
		panic("index must be 0...8")
	}

	sliceStart := idx * 9

	slice := make([]*int, 9)

	for i := 0; i < 9; i++ {
		slice[i] = &g.numbers[sliceStart+i]
	}

	return slice
}

///sliceColumn - Create a slice of pointers referencing the 9 numbers in the column at the given index.
func (g *suGrid) sliceColumn(idx int) []*int {
	if idx < 0 || idx > 8 {
		panic("index must be 0...8")
	}

	col := make([]*int, 9)
	for i := 0; i < 9; i++ {
		col[i] = &g.numbers[9*i+idx]
	}

	return col
}

///sliceBox - Create a slice of pointers referencing the 9 numbers in the box at the given index.
///Boxes are 3x3 subgrids numbered left-to-right, top-to-bottom
func (g *suGrid) sliceBox(idx int) []*int {
	if idx < 0 || idx > 8 {
		panic("index must be 0...8")
	}

	colStartIdx := (idx % 3) * 3
	rowStartIdx := (idx / 3) * 3

	box := make([]*int, 9)

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			box[i+j*3] = &g.numbers[(rowStartIdx+j)*9+colStartIdx+i]
		}
	}

	return box
}

func copyMap(original map[int]bool, dest map[int]bool) {
	for k, v := range original {
		dest[k] = v
	}
}

///initSetMaps - Step through each square and update the row, col, and box SetMaps.
///Returns an error if duplicates are found.
func (g *suGrid) initSetMaps() error {
	for i := 0; i < 9; i++ {
		g.rowSetMaps[i] = make(map[int]bool, 9)
		g.colSetMaps[i] = make(map[int]bool, 9)
		g.boxSetMaps[i] = make(map[int]bool, 9)

		for j := 0; j < 9; j++ {
			rowVal := *g.rows[i][j]
			if rowVal > 0 {
				if _, ok := g.rowSetMaps[i][rowVal]; ok {
					return fmt.Errorf("duplicate row val %d found at (%d, %d)", rowVal, i, j)
				} else {
					g.rowSetMaps[i][rowVal] = true
				}
			}

			colVal := *g.cols[i][j]
			if colVal > 0 {
				if _, ok := g.colSetMaps[i][colVal]; ok {
					return fmt.Errorf("duplicate col val %d found at (%d, %d)", colVal, i, j)
				} else {
					g.colSetMaps[i][colVal] = true
				}
			}

			boxVal := *g.boxes[i][j]
			if boxVal > 0 {
				if _, ok := g.boxSetMaps[i][boxVal]; ok {
					return fmt.Errorf("duplicate box val %d found at (%d, %d)", boxVal, i, j)
				} else {
					g.boxSetMaps[i][boxVal] = true
				}
			}
		}
	}

	return nil
}

///solveByElimination - Step through each empty box checking which potential values which can fill said box.
///The grid state is updated to indicate if the grid is unsolved, invalid, or solved.
///Returns true if a box has been filled.
func (g *suGrid) solveByElimination() bool {
	lowestOpenCount := 9
	g.lowestOpenIndex = -1
	solved := true

	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {
			idx := x + y*9
			if g.numbers[idx] != 0 {
				continue
			}

			solved = false

			number := g.numbers[idx]

			if number != 0 {
				continue
			}

			boxIdx := ((y / 3) * 3) + x/3

			openNumbers := make([]int, 0, 9)
			for i := 1; i <= 9; i++ {
				if !g.rowSetMaps[y][i] && !g.colSetMaps[x][i] && !g.boxSetMaps[boxIdx][i] {
					openNumbers = append(openNumbers, i)
				}
			}

			if len(openNumbers) == 0 {
				//There are no numbers that can go in this empty box.  The grid has no valid solution.
				g.gridState = Invalid
				return false
			}

			if len(openNumbers) == 1 {
				//There is only 1 possible solution for this box.
				g.numbers[idx] = openNumbers[0]
				g.rowSetMaps[y][openNumbers[0]] = true
				g.colSetMaps[x][openNumbers[0]] = true
				g.boxSetMaps[boxIdx][openNumbers[0]] = true

				g.gridState = Unsolved
				return true
			}

			if len(openNumbers) < lowestOpenCount {
				//Keep track of the open box with the least number of options for later guessing if needed.
				lowestOpenCount = len(openNumbers)
				g.lowestOpenIndex = idx
				g.lowestOpenNumbers = make([]int, lowestOpenCount)
				copy(g.lowestOpenNumbers, openNumbers)
			}
		}
	}

	if solved {
		g.gridState = Solved
	} else {
		g.gridState = Unsolved
	}

	return false
}

///formAtLowestOptions - Fork new grids with guesses at the value of the empty square with the least number of options.
///This relies on solveByElimination() having set the lowestOpenNumbers and lowestOpenIndex.
func (g *suGrid) forkAtLowestOptions() []*suGrid {
	if g.lowestOpenIndex < 0 || len(g.lowestOpenNumbers) < 2 {
		//solveByElimination() must be run first to populate these values.
		panic("lowest open index and numbers unset")
	}

	forks := make([]*suGrid, len(g.lowestOpenNumbers))
	for i := 0; i < len(g.lowestOpenNumbers); i++ {
		fork, _ := g.copy()
		forks[i] = fork

		guessVal := g.lowestOpenNumbers[i]

		fork.numbers[g.lowestOpenIndex] = guessVal

		//Update the SetMaps to reflect the filled guess.
		rowIdx := g.rowOfIndex(g.lowestOpenIndex)
		fork.rowSetMaps[rowIdx][guessVal] = true
		colIdx := g.colOfIndex(g.lowestOpenIndex)
		fork.colSetMaps[colIdx][guessVal] = true
		boxIdx := g.boxOfIndex(g.lowestOpenIndex)
		fork.boxSetMaps[boxIdx][guessVal] = true
	}

	return forks
}

///rowOfIndex - Get the row corresponding to the given number index.
func (g *suGrid) rowOfIndex(idx int) int {
	return idx / 9
}

///colOfIndex - Get the column corresponding to the given number index.
func (g *suGrid) colOfIndex(idx int) int {
	return idx % 9
}

///boxOfIndex - Get the box corresponding to the given number index.
func (g *suGrid) boxOfIndex(idx int) int {
	row := g.rowOfIndex(idx)
	col := g.colOfIndex(idx)

	return col/3 + (row/3)*3
}
