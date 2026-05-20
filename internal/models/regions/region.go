package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
)

type Region interface {
	GetCandidates() cells.ValuesSet
	GetCell(index uint8) *cells.Cell
	Valid() error
}

type linearRegion struct {
	Cells [9]*cells.Cell
}

func (c linearRegion) GetCandidates() cells.ValuesSet {
	candidate := cells.ValuesSet(0x0)

	for _, cell := range c.Cells {
		candidate |= cell.Candidates
	}

	return candidate
}

func (c linearRegion) GetCell(index uint8) *cells.Cell {
	return c.Cells[index]
}

func (c linearRegion) Valid() error {
	onRegion := cells.ValuesSet(0)

	for _, cell := range c.Cells {
		if cell.Value != 0 {
			if onRegion.Has(cell.Value) {
				return fmt.Errorf("Valor duplicado na mesma região: [%d]", cell.Value)
			}
			onRegion.Add(cell.Value)
		}
	}
	return nil
}
