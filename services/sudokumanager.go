package services

import (
	internalModels "Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/models"
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

func (s *SudokuManager) Step(state internalModels.GameState) internalModels.GameState {
	return state
}

func (s *SudokuManager) StepAll() {

}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}
