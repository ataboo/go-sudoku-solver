package main

import (
	"fmt"
	"os"

	"github.com/ataboo/go-sudoku-solver/pkg/server"
)

func main() {
	serveAddr := os.Getenv("SU_SERVER_PORT")
	if serveAddr == "" {
		serveAddr = "localhost:3000"
	}

	server := server.NewServer()

	if err := server.Run(serveAddr); err != nil {
		fmt.Printf("failed to start server: %s\n", err)
	}
}
