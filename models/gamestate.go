package models

import "Lechenco/sudoku-solver/internal/models"

type GameState struct {
	Board        models.Board
	InitialBoard models.Board
	Steps        []models.Step
}
