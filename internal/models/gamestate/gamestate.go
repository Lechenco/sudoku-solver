package gamestate

import "Lechenco/sudoku-solver/internal/models"

// GameState Store all the data for a particular puzzle.
//
// The initial and current Board, together with all the steps take it so far.
type GameState struct {
	Board        models.Board
	InitialBoard models.Board
	Steps        []Step
}

func (g *GameState) Valid() error {
	return g.Board.Valid()
}
