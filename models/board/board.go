package board

import (
	"github.com/Lechenco/sudoku-solver/models/cells"
)

// CellGrid is a abstraction from a 9x9 matrix of cells
type CellGrid [9][9]*cells.Cell

// SetValue sets value to the cell on position
func (g *CellGrid) SetValue(position cells.Position, value cells.Value) error {
	return g[position.RowNumber][position.ColumnNumber].SetValue(value)
}

type Board struct {
	Cells CellGrid
}

