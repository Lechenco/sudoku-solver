package io

import (
	"os"

	"github.com/Lechenco/sudoku-solver/internal/models"
	"github.com/Lechenco/sudoku-solver/utils/format"
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
