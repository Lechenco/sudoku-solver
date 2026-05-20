package services

import (
	internalModels "Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInit(t *testing.T) {
	assert := assert.New(t)

	manager := SudokuManager{}
	expectedConfig := models.GameConfig{
		InitialBoard: internalModels.Board{},
	}
	manager.Init(expectedConfig)

	assert.Equal(expectedConfig, manager.GameConfig)
	assert.Equal(expectedConfig.InitialBoard, manager.GameState.InitialBoard)
	assert.Equal(expectedConfig.InitialBoard, manager.GameState.Board)
}
