package io

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/utils/format"
	"os"
)

func ReadBoardFromFile(filename string) (models.Board, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return models.Board{}, err
	}
	return format.BoardFromString(string(data)), nil
}

func WriteBoardToFile(filename string, board models.Board) error {
	return os.WriteFile(filename, []byte(format.BoardToString(board)), 0644)
}
