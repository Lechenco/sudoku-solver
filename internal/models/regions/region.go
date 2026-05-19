package regions

import "Lechenco/sudoku-solver/internal/models"

type Region interface {
	MissingValues() []*models.Value
}
