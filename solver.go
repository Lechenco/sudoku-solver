package sudokusolver

import (
	"Lechenco/sudoku-solver/models"
	"Lechenco/sudoku-solver/services"
	"Lechenco/sudoku-solver/utils/format"
)

type SudokuSolver struct {
	gameManager models.GameManager
}

var newGameManager = func() models.GameManager {
	return &services.SudokuManager{}
}

func NewSudokuSolver(boardString string) (SudokuSolver, error) {
	solver := SudokuSolver{
		gameManager: newGameManager(),
	}
	solver.gameManager.Init(models.GameConfig{
		InitialBoard: format.BoardFromString(boardString),
	})
	err := solver.gameManager.ValidState()

	return solver, err
}
