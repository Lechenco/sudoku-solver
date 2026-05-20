package format

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardFromEmptyString(t *testing.T) {
	assert := assert.New(t)

	board := BoardFromString("")

	assert.NotNil(board)

	for _, row := range board.GetCells() {
		for _, cell := range row {
			assert.Equal(cells.Value(0), cell.Value)
		}
	}
}

func TestBoardFromStringToBig(t *testing.T) {
	assert := assert.New(t)

	board := BoardFromString("1234567891234567891234567891234567891234567891234567891234567891234567891234567891")

	assert.NotNil(board)

	for _, row := range board.GetCells() {
		for _, cell := range row {
			assert.NotEqual(cells.Value(0), cell.Value)
		}
	}
}
