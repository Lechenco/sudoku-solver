package sudokusolver

import (
	"Lechenco/sudoku-solver/models"
	"Lechenco/sudoku-solver/services"
	"Lechenco/sudoku-solver/utils/format"
)

type SudokuSolver struct {
	gameManager models.GameManager
}

func NewSudokuSolver(boardString string) (SudokuSolver, error) {
	solver := SudokuSolver{
		gameManager: &services.SudokuManager{},
	}
	solver.gameManager.Init(models.GameConfig{
		InitialBoard: format.BoardFromString(boardString),
	})
	err := solver.gameManager.ValidState()

	return solver, err
}
