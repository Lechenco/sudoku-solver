package regions

import "Lechenco/sudoku-solver/internal/models/cells"

type SquareRegion struct {
	Cells [3][3]*cells.Cell
}

func (s SquareRegion) GetCandidates() cells.Candidates {
	candidate := cells.Candidates(0x0)

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
