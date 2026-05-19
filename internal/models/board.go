package models

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
)

type Board struct {
	Cells   [9][9]*cells.Cell
	columns [9]regions.ColumnRegion
	rows    [9]regions.RowRegion
	squares [9]regions.SquareRegion
}

func (b *Board) InitRegions() {
	b.initColumns()
	b.initRows()
	b.initSquares()
}

func (b *Board) initColumns() {
	for i := range 9 {
		b.columns[i] = regions.NewColumnRegion(b.Cells[i])
	}
}

func (b *Board) initRows() {
	for j := range 9 {
		cells := [9]*cells.Cell{}
		for i := range 9 {
			cells[j] = b.Cells[i][j]
		}
		b.rows[j] = regions.NewRowsRegion(cells)
	}

}

func (b *Board) initSquares() {

}
