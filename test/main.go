package main

import (
	"Lechenco/sudoku-solver/internal/models"
	"fmt"
)

func main() {
	boardStr := "53..7....6..195....98....6.8...6...34..8..6...3...17...2...6....28....419..5...."
	board := models.NewBoard(boardStr)
	fmt.Println("Initial Board:")
	fmt.Println(board)
}
