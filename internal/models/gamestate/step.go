package gamestate

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
)

type Step interface {
	GetPosition() cells.Position
	GetValue() cells.Value
	GetStrategyName() string
	TakeStep(models.Board) error
}
