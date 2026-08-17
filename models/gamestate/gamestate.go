package gamestate

import (
	"encoding/json"

	"github.com/Lechenco/sudoku-solver/models/board"
)

// GameState Store all the data for a particular puzzle.
//
// The initial and current Board, together with all the steps take it so far.
type GameState struct {
	Board        board.Board
	InitialBoard board.Board
	Steps        []Step
}

func (g *GameState) UnmarshalJSON(b []byte) error {
	aux := &struct {
		Steps []SetValueStep
		Board board.Board
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	g.Steps = []Step{}
	g.Board = aux.Board
	for _, step := range aux.Steps {
		g.Steps = append(g.Steps, &step)
	}
	return nil
}
