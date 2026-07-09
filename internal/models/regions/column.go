package regions

import "Lechenco/sudoku-solver/internal/models/cells"

type ColumnRegion struct {
	linearRegion
}

func NewColumnRegion(cells [9]*cells.Cell) *ColumnRegion {
	res := &ColumnRegion{
		linearRegion{
			Cells: cells,
			baseRegion: &baseRegion{},
		},
	}

	res.baseRegion.Region = res
	return res
}
