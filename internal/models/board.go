package models

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
	"fmt"
)

type Board struct {
	Cells      [9][9]*cells.Cell
	Columns    [9]regions.ColumnRegion
	Rows       [9]regions.RowRegion
	Squares    [9]regions.SquareRegion
	AllRegions []regions.Region
}

func (b *Board) GetCells() [9][9]*cells.Cell {
	return b.Cells
}

func (b *Board) Valid() error {

	for _, col := range b.AllRegions {
		if err := col.Valid(); err != nil {
			fmt.Printf("%v", col)
			return err
		}
	}

	return nil
}

func (b *Board) Init() {
	b.initPositions()
	b.AllRegions = []regions.Region{}
	b.initColumns()
	b.initRows()
	b.initSquares()
}

func (b *Board) initPositions() {
	for i, row := range b.Cells {
		for j, cell := range row {
			cell.Position = cells.Position{RowNumber: uint8(i), ColumnNumber: uint8(j)}
		}
	}
}

func (b *Board) initRows() {
	for i := range 9 {
		b.Rows[i] = regions.NewRowsRegion(b.Cells[i])
		b.AllRegions = append(b.AllRegions, b.Rows[i])
	}
}

func (b *Board) initColumns() {
	for j := range 9 {
		cells := [9]*cells.Cell{}
		for i := range 9 {
			cells[i] = b.Cells[i][j]
		}
		b.Columns[j] = regions.NewColumnRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Columns[j])
	}

}

func (b *Board) initSquares() {
	for i := range 9 {
		var cells [3][3]*cells.Cell
		for j := range 9 {
			row, col := (i/3)*3+(j/3), (i%3)*3+(j%3)
			cells[j/3][j%3] = b.Cells[row][col]
		}
		b.Squares[i] = regions.NewSquareRegion(cells)
		b.AllRegions = append(b.AllRegions, b.Squares[i])
	}
}
