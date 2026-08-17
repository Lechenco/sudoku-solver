package io

import (
	"os"

	"github.com/Lechenco/sudoku-solver/models/board"
	"github.com/Lechenco/sudoku-solver/utils/format"
)

func ReadBoardFromFile(filename string) (board.Board, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return board.Board{}, err
	}
	return format.BoardFromString(string(data)), nil
}

func WriteBoardToFile(filename string, board board.Board) error {
	return os.WriteFile(filename, []byte(format.BoardToString(board)), 0644)
}
