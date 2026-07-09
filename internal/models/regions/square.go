package regions

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"slices"
)

type SquareRegion struct {
	*baseRegion
	Cells [3][3]*cells.Cell
}

func NewSquareRegion(cells [3][3]*cells.Cell) *SquareRegion {
	res := &SquareRegion{
		Cells: cells,
		baseRegion: &baseRegion{},
	}

	res.baseRegion.Region = res

	return res
}

func (s *SquareRegion) GetCells() []*cells.Cell {
	return slices.Concat(
		s.Cells[0][:],
		s.Cells[1][:],
		s.Cells[2][:],
	)
}

