package models_test

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/utils/format"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var newConst = strings.ReplaceAll(`
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

func TestInitRegions(t *testing.T) {
	assert := assert.New(t)

	b := format.BoardFromString(newConst)
	b.InitRegions()

	t.Run("Check all cells are correct", func(t *testing.T) {
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
			}
		}
	})
	t.Run("Check all rows are correct", func(t *testing.T) {
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
				assert.Equal(cells.Value(expectedSquares[(k/3)*3 + i][(k%3)*3 + j]), cell.Value)
			}
			}
		}
	})
}
