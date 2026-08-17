package gamestate

import (
	"github.com/Lechenco/sudoku-solver/models/cells"
	"github.com/Lechenco/sudoku-solver/internal/services"
)

type RemoveCandidatesStep struct {
	cells.Position
	RemoveSet cells.ValuesSet
	StepData
	StrategyName string
}

func (s *RemoveCandidatesStep) GetPosition() cells.Position {
	return s.Position
}

func (s *RemoveCandidatesStep) GetStrategyName() string {
	return s.StrategyName
}

func (s *RemoveCandidatesStep) TakeStep(service services.BoardService) error {
	return nil
}

func (s *RemoveCandidatesStep) GetData() StepData {
	return s.StepData
}
