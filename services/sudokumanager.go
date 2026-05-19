package services

import "Lechenco/sudoku-solver/models"

type SudokuManager struct {
	models.GameConfig
	models.GameState
}

func (s *SudokuManager) Init(config models.GameConfig) {
	s.GameConfig = config
	s.GameState = models.GameState{}
}

func (s *SudokuManager) Step(state models.GameState) models.GameState {
	return state
}

func (s *SudokuManager) StepAll() {

}
