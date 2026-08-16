package sudokusolver

import (
	"log/slog"

	"github.com/Lechenco/sudoku-solver/models"
	"github.com/Lechenco/sudoku-solver/services"
	"github.com/Lechenco/sudoku-solver/services/strategy"
	"github.com/Lechenco/sudoku-solver/utils/format"
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
		LoggerLevel: slog.LevelDebug,
	})
	err := solver.ValidState()

	return solver, err
}
