package sudokusolver

import (
	"Lechenco/sudoku-solver/internal/strategy"
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

func NewSudokuSolver(boardString string, strategies []strategy.Strategy) (SudokuSolver, error) {
	solver := SudokuSolver{
		newGameManager(),
	}
	solver.Init(models.GameConfig{
		InitialBoard: format.BoardFromString(boardString),
		Strategies:   strategies,
	})
	err := solver.ValidState()

	return solver, err
}
