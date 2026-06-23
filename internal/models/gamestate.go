package models

import "Lechenco/sudoku-solver/internal/models/step"

type GameState struct {
	Board        Board
	InitialBoard Board
	Steps        []step.Step
}

func (g *GameState) Valid() error {
	return g.Board.Valid()
}
