package sudokusolver

import (
	"Lechenco/sudoku-solver/models"
	"Lechenco/sudoku-solver/services"
	"Lechenco/sudoku-solver/utils/format"
)

type SudokuSolver struct {
	models.GameManager
}

var newGameManager = func() models.GameManager {
	return &services.SudokuManager{}
}

func NewSudokuSolver(boardString string) (SudokuSolver, error) {
	solver := SudokuSolver{
		newGameManager(),
	}
	solver.Init(models.GameConfig{
		InitialBoard: format.BoardFromString(boardString),
	})
	err := solver.ValidState()

	return solver, err
}
