package services

import (
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"Lechenco/sudoku-solver/models"
	"errors"
)

type SudokuManager struct {
	GameConfig models.GameConfig
	GameState  gamestate.GameState
}

func (s *SudokuManager) Init(config models.GameConfig) {
	s.GameConfig = config
	s.GameState = gamestate.GameState{
		InitialBoard: config.InitialBoard,
		Board:        config.InitialBoard,
	}
	s.GameState.Board.Init()
}

func (s *SudokuManager) Step() (gamestate.Step, error) {
	for _, strateg := range s.GameConfig.Strategies {
		step, err := strateg.Step(s.GameState)

		if err != nil {
			continue
		}

		s.GameState.Steps = append(s.GameState.Steps, step)
		return step, nil
	}
	return nil, errors.New("Não foi possível determinar o próximo passo")
}

func (s *SudokuManager) StepAll() {

}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}
