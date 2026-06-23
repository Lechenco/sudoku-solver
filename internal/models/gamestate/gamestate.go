package gamestate

import "Lechenco/sudoku-solver/internal/models"

type GameState struct {
	Board        models.Board
	InitialBoard models.Board
	Steps        []Step
}

func (g *GameState) Valid() error {
	return g.Board.Valid()
}
