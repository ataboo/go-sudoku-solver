package server

import (
	"net/http"

	"github.com/ataboo/go-sudoku-solver/pkg/solver"
	"github.com/gin-gonic/gin"
)

type SudokuServer struct {
	server *http.Server
}

func NewServer() *gin.Engine {
	server := gin.Default()

	server.GET("/", handleIndex)
	server.POST("/solve", handleSolve)

	return server
}

func handleIndex(c *gin.Context) {
	c.JSON(200, gin.H{"message": "pong"})
}

func handleSolve(c *gin.Context) {
	solveVM := vmSolvePuzzle{}
	if err := c.Bind(&solveVM); err != nil {
		c.JSON(http.StatusBadRequest, newVMSolvePuzzleResponse(Unparsable, nil))
		return
	}

	grid, err := solver.SudokuGridFromInts(solveVM.Numbers)
	if err != nil {
		c.JSON(http.StatusBadRequest, newVMSolvePuzzleResponse(Unparsable, nil))
		return
	}

	result, err := solver.SolveGrid(grid)
	if err != nil || len(result.SolvedGrids) == 0 {
		c.JSON(http.StatusConflict, newVMSolvePuzzleResponse(Unsolvable, nil))
		return
	}

	c.JSON(http.StatusOK, newVMSolvePuzzleResponse(NoError, result.SolvedGrids[0].IntValues()))
}
