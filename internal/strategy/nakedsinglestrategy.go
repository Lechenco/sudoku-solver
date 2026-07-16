package strategy

import (
	"Lechenco/sudoku-solver/internal/logging"
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"errors"
	"fmt"
	"log/slog"
	"sync"
	"time"
)

type nakedSingleStrategy struct {
	name string
	logger *slog.Logger
}

// Step iterates for all the cells looking for a cell with only one valid 
// candidate. Creating a step to set this candidate as the cell value.
func (n *nakedSingleStrategy) Step(gameState gamestate.GameState) (
	step gamestate.Step, err error) {
	n.logger.Info("Iniciando busca por passo com estratégia naked_single")
	data := gamestate.StepData{}
	start := time.Now()
	for c := range models.EmptyCellsIterator(gameState.Board.GetCells()) {
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
			n.logger.Info(fmt.Sprintf("Passo [%v] encontrado com estratégia naked_single", step))
			break
		}
	}

	if step == nil {
		n.logger.Error("Estratégia não conseguiu encontrar um step válido.")
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
		instance = &nakedSingleStrategy{
			name: "naked_single",
			logger: logging.LoggerFactory("strategy/nakedSingleStrategy"),
		}
	})
	return instance
}
