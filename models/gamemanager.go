package models

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/step"
	"Lechenco/sudoku-solver/internal/strategy"
	"Lechenco/sudoku-solver/utils/format"
)

type GameConfig struct {
	InitialBoard models.Board
	Strategies   []strategy.Strategy
}

type GameManager interface {
	Init(GameConfig)
	Step() (step.Step, error)
	StepAll()
	ValidState() error
}

func BoardOfGameState(board string) models.Board {
	return format.BoardFromString(board)
}
