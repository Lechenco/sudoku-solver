package gamestate

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
)

type SetValueStep struct {
	cells.Position
	StepData
	cells.Value
	StrategyName string
}

func (s *SetValueStep) GetPosition() cells.Position {
	return s.Position
}

func (s *SetValueStep) GetStrategyName() string {
	return s.StrategyName
}

func (s *SetValueStep) TakeStep(board models.Board) error {
	return board.Cells.SetValue(s.Position, s.Value)
}

func (s *SetValueStep) GetData() StepData {
	return s.StepData
}

func (s *SetValueStep) String() string {
	return fmt.Sprintf("{SetValueStep[%v]: %v}", s.Position, s.Value)
}
