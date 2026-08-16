package strategy

import (
	"github.com/Lechenco/sudoku-solver/models/gamestate"
)

type Strategy interface {
	Step(gamestate.GameState) (gamestate.Step, error)
}
