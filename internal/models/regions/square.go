package regions

import "Lechenco/sudoku-solver/internal/models"

type Square struct {
	Values [3][3]*models.Value
}
