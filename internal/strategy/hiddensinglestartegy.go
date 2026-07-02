package strategy

import (
	"Lechenco/sudoku-solver/internal/iterators"
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"errors"
	"slices"
	"sync"
	"time"
)

type hiddenSingleStrategy struct {
	name string
}

func (h *hiddenSingleStrategy) Step(gameState gamestate.GameState) (
	step gamestate.Step, err error,
) {
	data := gamestate.StepData{}
	start := time.Now()
	for region := range iterators.UnclearRegionsIterator(gameState.Board) {
		candidates := region.GetCellsCandidates()
		uniqueCandidate := cells.Uniques(candidates...)
		uniqueValues := uniqueCandidate.GetValues()

		data.Comparations += 1
		if len(uniqueValues) != 0 {
			cellIndex := slices.IndexFunc(region.GetCells(), func(c *cells.Cell) bool {
				return c.Candidates.Has(uniqueValues[0])
			})

			data.ExecutionTime = time.Since(start)
			step = &gamestate.SetValueStep{
				Position:     region.GetCells()[cellIndex].Position,
				Value:        uniqueValues[0],
				StrategyName: h.name,
				StepData:     data,
			}
			break
		}

	}

	if step == nil {
		err = errors.New("Estratégia não conseguiu encontrar um step válido")
	}
	return
}

var (
	hsinstance *hiddenSingleStrategy
	hsonce     sync.Once
)

func HiddenSingleStrategyInstance() *hiddenSingleStrategy {
	hsonce.Do(func() {
		hsinstance = &hiddenSingleStrategy{name: "hidden_single"}
	})
	return hsinstance
}
