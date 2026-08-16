package gamestate

import (
	"encoding/json"

	"github.com/Lechenco/sudoku-solver/internal/models"
)

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

func (g *GameState) UnmarshalJSON(b []byte) error {
	aux := &struct {
		Steps []SetValueStep
		Board models.Board
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
