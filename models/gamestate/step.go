package gamestate

import (
	"time"

	"github.com/Lechenco/sudoku-solver/models/cells"
	"github.com/Lechenco/sudoku-solver/internal/services"
)

type Step interface {
	GetPosition() cells.Position
	GetData() StepData
	GetStrategyName() string
	TakeStep(services.BoardService) error
}

type StepData struct {
	ExecutionTime time.Duration
	Comparations  int
}
