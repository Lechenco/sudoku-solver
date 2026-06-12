package strategy

import (
	"Lechenco/sudoku-solver/internal/models"
	"sync"
)

type nakedSingleStrategy struct{}

func (n *nakedSingleStrategy) Step(gameState models.GameState) (step *models.Step, err error) {
	return nil, nil
}

var (
	instance *nakedSingleStrategy
	once     sync.Once
)

func NakedSingleStrategyInstance() *nakedSingleStrategy {
	once.Do(func() {
		instance = &nakedSingleStrategy{}
	})
	return instance
}
