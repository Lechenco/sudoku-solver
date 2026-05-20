package models

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/utils/format"
)

type GameConfig struct {
	InitialBoard models.Board
}

type GameManager interface {
	Init(GameConfig)
	Step(GameState) GameState
	StepAll()
	ValidState() error
}

func BoardOfGameState(board string) models.Board {
	return format.BoardFromString(board)
}
