package models

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
	"Lechenco/sudoku-solver/utils"
	"slices"
)

type CellGrid [9][9]*cells.Cell

func (g *CellGrid) SetValue(position cells.Position, value cells.Value) error {
	return g[position.RowNumber][position.ColumnNumber].SetValue(value)
}

type Board struct {
	Cells          CellGrid
	Columns        [9]*regions.ColumnRegion
	Rows           [9]*regions.RowRegion
	Squares        [9]*regions.SquareRegion
	AllRegions     []regions.Region
	CleanedRegions []bool
}

func (b *Board) Finished() (res bool) {
	res = true
	for _, v := range b.CleanedRegions {
		res = res && v
	}

	return
}

func (b *Board) GetCells() CellGrid {
	return b.Cells
}

func (b *Board) SetValue(position cells.Position, value cells.Value) (err error) {
	err = b.Cells.SetValue(position, value)

	if err != nil {
		return
	}

	err = b.Valid()

	if err != nil {
		return
	}

	b.cleanFromRegions(position, value)

	return
}

func (b *Board) cleanFromRegions(position cells.Position, value cells.Value) {
	for _, region := range b.getRegions(position) {
		region.RemoveCandidate(value)
		if region.GetCandidates().IsEmpty() {
			index := slices.IndexFunc(b.AllRegions, func(r regions.Region) bool {
				return region == r
			})
			b.CleanedRegions[index] = true
		}
	}
}

func (b *Board) getRegions(position cells.Position) (arr []regions.Region) {
	arr = append(arr, b.Columns[position.ColumnNumber])
	arr = append(arr, b.Rows[position.RowNumber])

	row, col := utils.IndexSquareRegionPos(position)
	arr = append(arr, b.Squares[3*row+col])

	return
}

func (b *Board) Valid() error {

	for _, region := range b.AllRegions {
		if err := region.Valid(); err != nil {
			return err
		}
	}

	return nil
}

func (b *Board) Init() {
	b.initPositions()
	b.initRegions()
	b.initCandidates()
}

func (b *Board) initCandidates() {
	for _, rows := range b.Cells {
		for _, cell := range rows {
			if !cell.IsEmpty() {
				b.cleanFromRegions(cell.Position, cell.Value)
			}
		}
	}
}

func (b *Board) initRegions() {
	b.AllRegions = []regions.Region{}
	b.initColumns()
	b.initRows()
	b.initSquares()
	b.CleanedRegions = make([]bool, len(b.AllRegions))
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
