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

type baseRegion struct {
	Region
}

func (c *baseRegion) GetCandidates() cells.ValuesSet {
	candidate := cells.ValuesSet(0x0)

	for _, cell := range c.GetCells() {
		candidate |= cell.Candidates
	}

	return candidate
}

func (c *baseRegion) GetCell(index uint8) *cells.Cell {
	return c.GetCells()[index]
}

// Valid iterate over the region cells searching for a broken puzzle restrition:
//   - Duplicated value in a same region
//   - Empty cell with no possible candidate
// In one of these cases, return a new error
func (c *baseRegion) Valid() error {
	onRegion := cells.ValuesSet(0)

	for _, cell := range c.GetCells() {
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

// RemoveCandidate remove value from all region cells
func (c *baseRegion) RemoveCandidate(value cells.Value) {
	for _, cell := range c.GetCells() {
		cell.Candidates.Remove(value)
	}
}

func (c *baseRegion) GetCellsCandidates() []cells.ValuesSet {
	cs := c.GetCells()
	candidates := make([]cells.ValuesSet, len(cs))

	for i, cell := range cs{
		candidates[i] = cell.Candidates
	}
	return candidates
}

type linearRegion struct {
	*baseRegion
	Cells [9]*cells.Cell
}

func (c *linearRegion) GetCells() []*cells.Cell {
	return c.Cells[:]
}

