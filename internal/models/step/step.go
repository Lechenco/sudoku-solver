package step

import "Lechenco/sudoku-solver/internal/models/cells"

type Step struct {
	cells.Position
	cells.Value
	StrategyName string
}
