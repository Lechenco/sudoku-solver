package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
)

type Region interface {
	GetCandidates() cells.Candidates
	GetCell(index uint8) *cells.Cell
}

type linearRegion struct {
	Cells [9]*cells.Cell
}

func (c linearRegion) GetCandidates() cells.Candidates {
	candidate := cells.Candidates(0x0)

	for _, cell := range c.Cells {
		candidate |= cell.Candidates
	}

	return candidate
}

func (c linearRegion) GetCell(index uint8) *cells.Cell {
	return c.Cells[index]
}
