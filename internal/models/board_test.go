package models_test

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/utils/format"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var initBoardString = strings.ReplaceAll(`
1..4..7..
.2..5..8.
..3..6..9
..3..6..9
.2..5..8.
1..4..7..
1..4..7..
.2..5..8.
..3..6..9
`, "\n", "")

func TestInit(t *testing.T) {
	b := format.BoardFromString(initBoardString)
	b.Init()

	t.Run("Check all cells are correct", func(t *testing.T) {
		assert := assert.New(t)
		expectedCells := [][]uint8{
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
		}
		for i, row := range b.Cells {
			for j, cell := range row {
				assert.Equal(cells.Value(expectedCells[i][j]), cell.Value)
				assert.Equal(uint8(i), cell.Position.RowNumber)
				assert.Equal(uint8(j), cell.Position.ColumnNumber)
			}
		}
	})
	t.Run("Check candidates are correct", func(t *testing.T) {
		assert := assert.New(t)
		
		assert.Equal(cells.ValuesSet(0), b.Cells[0][0].Candidates)
		assert.Equal(cells.ValuesSet(0x1b0), b.Cells[0][1].Candidates)
		assert.Equal(cells.ValuesSet(0x36), b.Cells[6][8].Candidates)
	})
	t.Run("Check all rows are correct", func(t *testing.T) {
		assert := assert.New(t)
		expectedRows := [][]uint8{
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
		}
		for i, row := range b.Rows {
			for j, cell := range row.Cells {
				assert.Equal(cells.Value(expectedRows[i][j]), cell.Value)
			}
		}
	})
	t.Run("Check all columns are correct", func(t *testing.T) {
		assert := assert.New(t)
		expectedColumns := [][]uint8{
			{1, 0, 0, 0, 0, 1, 1, 0, 0},
			{0, 2, 0, 0, 2, 0, 0, 2, 0},
			{0, 0, 3, 3, 0, 0, 0, 0, 3},
			{4, 0, 0, 0, 0, 4, 4, 0, 0},
			{0, 5, 0, 0, 5, 0, 0, 5, 0},
			{0, 0, 6, 6, 0, 0, 0, 0, 6},
			{7, 0, 0, 0, 0, 7, 7, 0, 0},
			{0, 8, 0, 0, 8, 0, 0, 8, 0},
			{0, 0, 9, 9, 0, 0, 0, 0, 9},
		}
		for i, column := range b.Columns {
			for j, cell := range column.Cells {
				assert.Equal(cells.Value(expectedColumns[i][j]), cell.Value)
			}
		}
	})
	t.Run("Check all squares are correct", func(t *testing.T) {
		assert := assert.New(t)
		expectedSquares := [][]uint8{
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{1, 0, 0, 4, 0, 0, 7, 0, 0},
			{0, 2, 0, 0, 5, 0, 0, 8, 0},
			{0, 0, 3, 0, 0, 6, 0, 0, 9},
		}
		for k, square := range b.Squares {
			for i, row := range square.Cells {
				for j, cell := range row {
					expectedValue := expectedSquares[(k/3)*3+i][(k%3)*3+j]
					assert.Equal(cells.Value(expectedValue), cell.Value)
				}
			}
		}
	})
	t.Run("Check cleanedRegions was instanciated", func(t *testing.T) {
		assert := assert.New(t)
		assert.NotNil(b.CleanedRegions)
		assert.Equal(27, len(b.CleanedRegions))
		for _, b := range b.CleanedRegions {
			assert.False(b)
		}
	})
}

var setValueBoardString = strings.ReplaceAll(`
1..4..2..
.4..5..9.
..3..6..8
..1..2..4
.2..1..3.
9..8..6..
6..2..7..
.1..4..8.
..7..8..9
`, "\n", "")

func TestSetValue(t *testing.T) {

	b := format.BoardFromString(setValueBoardString)
	b.Init()

	t.Run("SetValue should change the cell position value", func(t *testing.T) {
		assert := assert.New(t)
		err := b.SetValue(cells.Position{
			RowNumber:    1,
			ColumnNumber: 0,
		}, cells.Value(7))

		assert.Nil(err)
		assert.Equal(cells.Value(7), b.Cells[1][0].Value)
	})

	t.Run("SetValue should remove candidate from regions", func(t *testing.T) {
		assert := assert.New(t)

		assert.True(b.Columns[0].GetCandidates().Has(8))
		assert.True(b.Rows[1].GetCandidates().Has(8))
		assert.True(b.Squares[0].GetCandidates().Has(8))

		err := b.SetValue(cells.Position{
			RowNumber:    0,
			ColumnNumber: 1,
		}, cells.Value(8))

		assert.Nil(err)
		assert.False(b.Columns[1].GetCandidates().Has(8))
		assert.False(b.Rows[0].GetCandidates().Has(8))
		assert.False(b.Squares[0].GetCandidates().Has(8))
	})

	t.Run("SetValue should remove candidate from correct regions (mirror)", func(t *testing.T) {
		assert := assert.New(t)

		assert.True(b.Columns[7].GetCandidates().Has(4))
		assert.True(b.Rows[2].GetCandidates().Has(4))
		assert.True(b.Squares[2].GetCandidates().Has(4))

		err := b.SetValue(cells.Position{
			RowNumber:    2,
			ColumnNumber: 7,
		}, cells.Value(4))

		assert.Nil(err)
		assert.False(b.Columns[7].GetCandidates().Has(4))
		assert.False(b.Rows[2].GetCandidates().Has(4))
		assert.False(b.Squares[2].GetCandidates().Has(4))
	})

	t.Run("SetValue should check if a region became completed", func(t *testing.T) {
		assert := assert.New(t)

		assert.False(b.CleanedRegions[18])

		err := b.SetValue(cells.Position{
			RowNumber:    0,
			ColumnNumber: 2,
		}, cells.Value(5))

		assert.Nil(err)
		assert.False(b.CleanedRegions[18])

		err = b.SetValue(cells.Position{
			RowNumber:    1,
			ColumnNumber: 2,
		}, cells.Value(6))

		assert.Nil(err)
		assert.False(b.CleanedRegions[18])

		err = b.SetValue(cells.Position{
			RowNumber:    2,
			ColumnNumber: 0,
		}, cells.Value(2))

		assert.Nil(err)
		assert.False(b.CleanedRegions[18])

		err = b.SetValue(cells.Position{
			RowNumber:    2,
			ColumnNumber: 1,
		}, cells.Value(9))

		assert.Nil(err)
		assert.True(b.CleanedRegions[18])
	})

	t.Run("SetValue should not allowed to override a existed value", func(t *testing.T) {
		assert := assert.New(t)
		err := b.SetValue(cells.Position{
			RowNumber:    0,
			ColumnNumber: 0,
		}, cells.Value(5))

		assert.EqualError(err, "{Cell[{0 0}]: 1} already filled, overrides are not allowed.")
	})

	t.Run("SetValue should not allowed to set as zero", func(t *testing.T) {

		assert := assert.New(t)
		err := b.SetValue(cells.Position{
			RowNumber:    0,
			ColumnNumber: 5,
		}, cells.Value(0))

		assert.EqualError(err, "{Cell[{0 5}]: _} cannot set a cell value to zero")
	})

	t.Run("SetValue should valid the Board after set the cell value", func(t *testing.T) {

		assert := assert.New(t)
		err := b.SetValue(cells.Position{
			RowNumber:    8,
			ColumnNumber: 7,
		}, cells.Value(7))

		assert.EqualError(err, "Valor duplicado na mesma região: [{RowRegion[_,_,7,_,_,8,_,7,9]}]")
	})
}
