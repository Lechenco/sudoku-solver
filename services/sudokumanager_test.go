package services_test

import (
	"Lechenco/sudoku-solver/models"
	"Lechenco/sudoku-solver/services"
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
	assert := assert.New(t)

	board := format.BoardFromString(initBoardString)

	manager := services.SudokuManager{}
	expectedConfig := models.GameConfig{
		InitialBoard: board,
	}
	manager.Init(expectedConfig)

	assert.Equal(expectedConfig, manager.GameConfig)
	assert.Equal(expectedConfig.InitialBoard, manager.GameState.InitialBoard)
	assert.NotNil(manager.GameState.Board)
}
