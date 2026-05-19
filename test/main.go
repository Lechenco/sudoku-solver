package main

import (
	"Lechenco/sudoku-solver/utils/format"
	"fmt"
)

func main() {
	boardStr := "53..7....6..195....98....6.8...6...34..8..6...3...17...2...6....28....419..5...."
	board := format.BoardFromString(boardStr)
	fmt.Println("Initial Board:")
	fmt.Println(format.BoardToString(board))
}
