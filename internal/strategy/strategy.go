package strategy

import "Lechenco/sudoku-solver/internal/models"


type Strategy interface {
	Step(models.GameState) (*models.Step, error)
}
