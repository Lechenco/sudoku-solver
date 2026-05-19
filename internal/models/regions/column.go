package regions

import "Lechenco/sudoku-solver/internal/models/cells"

type ColumnRegion struct {
	linearRegion
}

func NewColumnRegion(cells [9]*cells.Cell) ColumnRegion {
	return ColumnRegion{
		linearRegion{
			Cells: cells,
		},
	}
}
