package models

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
)

type Board struct {
	Cells      [9][9]*cells.Cell
	columns    [9]regions.ColumnRegion
	rows       [9]regions.RowRegion
	squares    [9]regions.SquareRegion
	allRegions []regions.Region
}

func (b *Board) GetCells() [9][9]*cells.Cell {
	return b.Cells
}

func (b *Board) Valid() error {

	for _, col := range b.allRegions {
		if err := col.Valid(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) InitRegions() {
	b.allRegions = []regions.Region{}
	b.initColumns()
	b.initRows()
	b.initSquares()
}

func (b *Board) initColumns() {
	for i := range 9 {
		b.columns[i] = regions.NewColumnRegion(b.Cells[i])
		b.allRegions = append(b.allRegions, b.columns[i])
	}
}

func (b *Board) initRows() {
	for j := range 9 {
		cells := [9]*cells.Cell{}
		for i := range 9 {
			cells[i] = b.Cells[i][j]
		}
		b.rows[j] = regions.NewRowsRegion(cells)
		b.allRegions = append(b.allRegions, b.rows[j])
	}

}

func (b *Board) initSquares() {
	for i := range 9 {
		var cells [3][3]*cells.Cell
		for j := range 9 {
			cells[j/3][j%3] = b.Cells[(i%3)*3+(j%3)][(i/3)+(j%3)]
		}
		b.squares[i] = regions.NewSquareRegion(cells)
		b.allRegions = append(b.allRegions, b.squares[i])
	}
}
