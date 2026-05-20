package sudokusolver

import (
	"errors"
	"testing"

	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/models"

	"github.com/stretchr/testify/assert"
)

type spyGameManager struct {
	initCalled       bool
	validStateCalled bool
	initConfig       models.GameConfig
	validStateErr    error
}

func (s *spyGameManager) Init(config models.GameConfig) {
	s.initCalled = true
	s.initConfig = config
}

func (s *spyGameManager) Step(state models.GameState) models.GameState {
	return state
}

func (s *spyGameManager) StepAll() {}

func (s *spyGameManager) ValidState() error {
	s.validStateCalled = true
	return s.validStateErr
}

func TestNewSudokuSolver_CallsInitAndValidState(t *testing.T) {
	assert := assert.New(t)
	previousManager := newGameManager
	defer func() { newGameManager = previousManager }()

	spy := &spyGameManager{}
	newGameManager = func() models.GameManager { return spy }

	solver, err := NewSudokuSolver("5")

	assert.Nil(err, "expected no error, got %v", err)
	assert.True(spy.initCalled, "expected Init to be called")
	assert.True(spy.validStateCalled)
	assert.Equal(solver.gameManager, spy)
	assert.Equal(cells.Value(5), spy.initConfig.InitialBoard.GetCells()[0][0].Value)
}

func TestNewSudokuSolver_PropagatesValidStateError(t *testing.T) {
	assert := assert.New(t)
	previousManager := newGameManager
	defer func() { newGameManager = previousManager }()

	expectedErr := errors.New("invalid game state")
	spy := &spyGameManager{validStateErr: expectedErr}
	newGameManager = func() models.GameManager { return spy }

	solver, err := NewSudokuSolver("5")
	assert.Equal(expectedErr, err)
	assert.True(spy.initCalled)
	assert.True(spy.validStateCalled)
	assert.Equal(spy, solver.gameManager)
}
