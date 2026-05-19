package regions

import "Lechenco/sudoku-solver/internal/models/cells"

type RowRegion struct {
	linearRegion
}

func NewRowsRegion(cells [9]*cells.Cell) RowRegion {
	return RowRegion{
		linearRegion{
			Cells: cells,
		},
	}
}
