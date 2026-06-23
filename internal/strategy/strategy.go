package strategy

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/step"
)


type Strategy interface {
	Step(models.GameState) (*step.Step, error)
}
