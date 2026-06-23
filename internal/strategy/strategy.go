package strategy

import (
	"Lechenco/sudoku-solver/internal/models/gamestate"
)

type Strategy interface {
	Step(gamestate.GameState) (gamestate.Step, error)
}
