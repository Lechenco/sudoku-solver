package services

import (
	internalModels "Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/step"
	"Lechenco/sudoku-solver/models"
	"errors"
)

type SudokuManager struct {
	GameConfig models.GameConfig
	GameState  internalModels.GameState
}

func (s *SudokuManager) Init(config models.GameConfig) {
	s.GameConfig = config
	s.GameState = internalModels.GameState{
		InitialBoard: config.InitialBoard,
		Board:        config.InitialBoard,
	}
	s.GameState.Board.Init()
}

func (s *SudokuManager) Step() (step.Step, error) {
	for _, strateg := range s.GameConfig.Strategies {
		step, err := strateg.Step(s.GameState)

		if err != nil {
			continue
		}
		
		s.GameState.Steps = append(s.GameState.Steps, *step)
		return *step, nil
	}
	return step.Step{}, errors.New("Não foi possível determinar o próximo passo")
}

func (s *SudokuManager) StepAll() {

}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}
