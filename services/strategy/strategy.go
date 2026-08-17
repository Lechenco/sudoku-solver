package strategy

import (
	"github.com/Lechenco/sudoku-solver/internal/services"
	"github.com/Lechenco/sudoku-solver/models/gamestate"
)

type Strategy interface {
	Step(gamestate.GameState, services.BoardService) (gamestate.Step, error)
}
