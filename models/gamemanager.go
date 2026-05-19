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
}

func BoardOfGameState(board string) models.Board {
	return format.BoardFromString(board)
}
