package gamestate

import (
	"time"

	"github.com/Lechenco/sudoku-solver/internal/models"
	"github.com/Lechenco/sudoku-solver/internal/models/cells"
)

type Step interface {
	GetPosition() cells.Position
	GetData() StepData
	GetStrategyName() string
	TakeStep(models.Board) error
}

type StepData struct {
	ExecutionTime time.Duration
	Comparations  int
}
