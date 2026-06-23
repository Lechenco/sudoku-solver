package gamestate

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
)

type SetValueStep struct {
	cells.Position
	cells.Value
	StrategyName string
}

func (s *SetValueStep) GetPosition() cells.Position {
	return s.Position
}

func (s *SetValueStep) GetValue() cells.Value {
	return s.Value
}

func (s *SetValueStep) GetStrategyName() string {
	return s.StrategyName
}

func (s *SetValueStep) TakeStep(models.Board) error {
	return nil
}
