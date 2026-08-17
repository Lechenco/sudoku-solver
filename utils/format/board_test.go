package format

import (
	"testing"

	"github.com/Lechenco/sudoku-solver/models/cells"

	"github.com/stretchr/testify/assert"
)

func TestBoardFromEmptyString(t *testing.T) {
	assert := assert.New(t)

	board := BoardFromString("")

	assert.NotNil(board)

	for _, row := range board.Cells {
		for _, cell := range row {
			assert.Equal(cells.Value(0), cell.Value)
		}
	}
}

func TestBoardFromStringToBig(t *testing.T) {
	assert := assert.New(t)

	board := BoardFromString("1234567891234567891234567891234567891234567891234567891234567891234567891234567891")

	assert.NotNil(board)

	for _, row := range board.Cells {
		for _, cell := range row {
			assert.NotEqual(cells.Value(0), cell.Value)
		}
	}
}
