package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
)

type RowRegion struct {
	linearRegion
}

func NewRowsRegion(cells [9]*cells.Cell) *RowRegion {
	res := &RowRegion{
		linearRegion{
			Cells: cells,
			baseRegion: &baseRegion{},
		},
	}
	res.baseRegion.Region = res

	return res
}

