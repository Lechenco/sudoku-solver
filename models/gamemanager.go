package models

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/Lechenco/sudoku-solver/internal/models"
	"github.com/Lechenco/sudoku-solver/internal/models/gamestate"
	"github.com/Lechenco/sudoku-solver/services/strategy"
	"github.com/Lechenco/sudoku-solver/utils/format"
)

type GameConfig struct {
	InitialBoard models.Board
	Strategies   []strategy.Strategy
	LoggerLevel  slog.Level
}

type GameManager interface {
	Init(GameConfig)
	Step() (gamestate.Step, error)
	StepAll() error
	ValidState() error
	Finished() bool
	ToFile(filename string) error
	InitFromFile(filename string) error
}

func BoardOfGameState(board string) models.Board {
	return format.BoardFromString(board)
}

func (c GameConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		Strategies []strategy.Strategy
	}{
		Strategies: c.Strategies,
	})
}

func (c *GameConfig) UnmarshalJSON(b []byte) error {
	aux := &struct{
		Strategies []struct{
			Name string
		}
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	c.Strategies = []strategy.Strategy{}
	for _, s := range aux.Strategies {
		switch s.Name {
		case "hidden_single":
			c.Strategies = append(c.Strategies, strategy.HiddenSingleStrategyInstance())
		case "naked_single":
			c.Strategies = append(c.Strategies, strategy.NakedSingleStrategyInstance())
		default:
			return fmt.Errorf("Estratégia %s não reconhecida", s)
		}
	}

	return nil
}
