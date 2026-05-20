package services

import "Lechenco/sudoku-solver/models"

type SudokuManager struct {
	GameConfig models.GameConfig
	GameState  models.GameState
}

func (s *SudokuManager) Init(config models.GameConfig) {
	s.GameConfig = config
	s.GameState = models.GameState{
		InitialBoard: config.InitialBoard,
		Board:        config.InitialBoard,
	}
	s.GameState.Board.InitRegions()
}

func (s *SudokuManager) Step(state models.GameState) models.GameState {
	return state
}

func (s *SudokuManager) StepAll() {

}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}
