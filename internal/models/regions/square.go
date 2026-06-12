package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"fmt"
)

type SquareRegion struct {
	Cells [3][3]*cells.Cell
}

func (s SquareRegion) GetCandidates() cells.ValuesSet {
	candidate := cells.ValuesSet(0x0)

	for _, cellRow := range s.Cells {
		for _, cell := range cellRow {
			candidate |= cell.Candidates
		}
	}

	return candidate
}

func (s SquareRegion) GetCell(index uint8) *cells.Cell {
	return s.Cells[index/3][index%3]
}

func (s SquareRegion) Valid() error {
	onRegion := cells.ValuesSet(0)

	for _, rows := range s.Cells {
		for _, cell := range rows {
			if cell.Value != 0 {
				if onRegion.Has(cell.Value) {
					return fmt.Errorf("Valor duplicado na mesma região: [%v]", cell)
				}
				onRegion.Add(cell.Value)
			}
		}
	}
	return nil
}

func NewSquareRegion(cells [3][3]*cells.Cell) *SquareRegion {
	return &SquareRegion{
		Cells: cells,
	}
}

func (s *SquareRegion) RemoveCandidate(value cells.Value) {
	for _, rows := range s.Cells {
		for _, cell := range rows {
			cell.Candidates.Remove(value)
		}
	}
}
