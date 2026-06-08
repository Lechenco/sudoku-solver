package models

import "Lechenco/sudoku-solver/internal/models/cells"

type Step struct {
	cells.Position
	cells.Value
}
