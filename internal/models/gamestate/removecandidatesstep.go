package gamestate

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
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

func (s *RemoveCandidatesStep) TakeStep(board models.Board) error {
	return nil
}

func (s *RemoveCandidatesStep) GetData() StepData {
	return s.StepData
}
