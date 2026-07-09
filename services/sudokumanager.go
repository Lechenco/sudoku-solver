package services

import (
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"Lechenco/sudoku-solver/models"
	"errors"
	"fmt"
)

// SudokuManager takes a game of sudoku, step by step, from multiple strategies.
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

// Step search for the next step based on a array of strategies:
//   - take the first strategy
//   - search for a step
//   - check for a error or a step was not found. If so continue for the next strategy.
//   - take the step
func (s *SudokuManager) Step() (gamestate.Step, error) {
	for _, strateg := range s.GameConfig.Strategies {
		step, err := strateg.Step(s.GameState)

		if err != nil || step == nil {
			continue
		}

		s.GameState.Steps = append(s.GameState.Steps, step)

		return step, step.TakeStep(s.GameState.Board)
	}
	return nil, errors.New("Não foi possível determinar o próximo passo")
}

// StepAll continues to take steps until find a error or finish the puzzle.
func (s *SudokuManager) StepAll() error {
	for {
		step, err := s.Step()

		fmt.Println(step, err)

		if err != nil {
			return err
		}
		if valid := s.ValidState(); valid != nil {
			return valid
		}
		if s.Finished() {
			break
		}
	}

	return nil
}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}

func (s *SudokuManager) Finished() bool {
	return s.GameState.Board.Finished()
}
