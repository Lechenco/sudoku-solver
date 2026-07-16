package models

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"Lechenco/sudoku-solver/internal/strategy"
	"Lechenco/sudoku-solver/utils/format"
	"log/slog"
)

type GameConfig struct {
	InitialBoard models.Board
	Strategies   []strategy.Strategy
	LoggerLevel  slog.Level
}

type GameManager interface {
	Init(GameConfig)
	Step() (gamestate.Step, error)
	StepAll() error
	ValidState() error
	Finished() bool
}

func BoardOfGameState(board string) models.Board {
	return format.BoardFromString(board)
}
