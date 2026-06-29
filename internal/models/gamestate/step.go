package gamestate

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
	"time"
)

type Step interface {
	GetPosition() cells.Position
	GetData() StepData
	GetStrategyName() string
	TakeStep(models.Board) error
}

type StepData struct {
	ExecutionTime time.Duration
	Comparations int
}
