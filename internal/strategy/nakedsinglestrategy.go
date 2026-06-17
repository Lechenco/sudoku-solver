package strategy

import (
	"Lechenco/sudoku-solver/internal/iterators"
	"Lechenco/sudoku-solver/internal/models"
	"errors"
	"sync"
)

type nakedSingleStrategy struct{
	name string
}

func (n *nakedSingleStrategy) Step(gameState models.GameState) (step *models.Step, err error) {
	for c := range iterators.CellsIterator(gameState.Board.GetCells()){
			values := c.Candidates.GetValues()
			if len(values) == 1 {
				step = &models.Step{
					Position: c.Position,
					Value: values[0],
					StrategyName: n.name,
				}
			}

		if step != nil {
			break
		}
	}

	if step == nil {
		err = errors.New("Estratégia não conseguiu encontrar um step válido")
	}

	return
}

var (
	instance *nakedSingleStrategy
	once     sync.Once
)

func NakedSingleStrategyInstance() *nakedSingleStrategy {
	once.Do(func() {
		instance = &nakedSingleStrategy{name: "naked_single"}
	})
	return instance
}
