package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
)

type Region interface {
	GetCandidates() cells.ValuesSet
	GetCell(index uint8) *cells.Cell
	GetCells() []*cells.Cell
	GetCellsCandidates() []cells.ValuesSet
	Valid() error
	RemoveCandidate(value cells.Value)
}

type linearRegion struct {
	Cells [9]*cells.Cell
}

func (c *linearRegion) GetCandidates() cells.ValuesSet {
	candidate := cells.ValuesSet(0x0)

	for _, cell := range c.Cells {
		candidate |= cell.Candidates
	}

	return candidate
}

func (c *linearRegion) GetCell(index uint8) *cells.Cell {
	return c.Cells[index]
}

func (c *linearRegion) Valid() error {
	onRegion := cells.ValuesSet(0)

	for _, cell := range c.Cells {
		if cell.Value != 0 {
			if onRegion.Has(cell.Value) {
				return fmt.Errorf("Valor duplicado na mesma região: [%v]", cell)
			}
			onRegion.Add(cell.Value)
		} else if cell.Candidates.IsEmpty() {
			return fmt.Errorf("Célula vazia com nenhum candidato possível: [%v]", cell)
		}
	}
	return nil
}

func (c *linearRegion) RemoveCandidate(value cells.Value) {
	for _, cell := range c.Cells {
		cell.Candidates.Remove(value)
	}
}

func (c *linearRegion) GetCells() []*cells.Cell {
	return c.Cells[:]
}

func (c *linearRegion) GetCellsCandidates() []cells.ValuesSet {
	cs := c.GetCells()
	candidates := make([]cells.ValuesSet, len(cs))

	for i, cell := range cs{
		candidates[i] = cell.Candidates
	}
	return candidates
}

