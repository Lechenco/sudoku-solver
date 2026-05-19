package models

import "Lechenco/sudoku-solver/internal/models/cells"

type Step struct {
	Position
	cells.Value
}
