package strategy

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"Lechenco/sudoku-solver/internal/models/regions"
	"errors"
	"slices"
	"sync"
	"time"
)

type hiddenSingleStrategy struct {
	name string
}

// Step search for region with a certain value in a single cell, create a step 
// to apply the change. 
//
// Return error if none valid step was founded
func (h *hiddenSingleStrategy) Step(gameState gamestate.GameState) (
	step gamestate.Step, err error,
) {
	data := gamestate.StepData{}
	start := time.Now()
	for region := range models.UnclearRegionsIterator(gameState.Board) {
		uniqueValues := findUniqueValues(region)

		data.Comparations += 1
		if len(uniqueValues) != 0 {
			cell := findCell(region.GetCells(), uniqueValues[0])

			data.ExecutionTime = time.Since(start)
			step = &gamestate.SetValueStep{
				Position:     cell.Position,
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

// findUniqueValues iterates over all region cell candidates searching for 
// values that appears only once. Returns the founded values array.
func findUniqueValues(region regions.Region) []cells.Value {
	candidates := region.GetCellsCandidates()
	uniqueCandidate := cells.Uniques(candidates...)
	return uniqueCandidate.GetValues()
}

// findCell searchs in a cells array for the first cell with the candidate value.
// Returns the founded cell.
func findCell(cellsArray []*cells.Cell, value cells.Value) *cells.Cell {
	cellIndex := slices.IndexFunc(cellsArray, func(c *cells.Cell) bool {
		return c.Candidates.Has(value)
	})

	return cellsArray[cellIndex]
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
