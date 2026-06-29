package strategy

import (
	"Lechenco/sudoku-solver/internal/iterators"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"errors"
	"sync"
	"time"
)

type nakedSingleStrategy struct {
	name string
}

func (n *nakedSingleStrategy) Step(gameState gamestate.GameState) (
	step gamestate.Step, err error) {
	data := gamestate.StepData{}
	start := time.Now()
	for c := range iterators.EmptyCellsIterator(gameState.Board.GetCells()) {
		data.Comparations += 1
		if c.Candidates.GetNumberOfValues() == 1 {
			values := c.Candidates.GetValues()
			data.ExecutionTime = time.Since(start)

			step = &gamestate.SetValueStep{
				Position:     c.Position,
				Value:        values[0],
				StrategyName: n.name,
				StepData: data,
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
