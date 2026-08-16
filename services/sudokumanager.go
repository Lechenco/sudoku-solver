package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"

	"github.com/Lechenco/sudoku-solver/internal/logging"
	"github.com/Lechenco/sudoku-solver/internal/models/gamestate"
	"github.com/Lechenco/sudoku-solver/models"
	"github.com/Lechenco/sudoku-solver/utils/io"
)

// SudokuManager takes a game of sudoku, step by step, from multiple strategies.
type SudokuManager struct {
	GameConfig models.GameConfig
	GameState  gamestate.GameState
	logger	*slog.Logger
}

func (s *SudokuManager) Init(config models.GameConfig) {
	s.GameConfig = config
	s.setLogger()

	s.logger.Debug("Iniciando SudokuManager...")
	s.GameState = gamestate.GameState{
		InitialBoard: config.InitialBoard,
		Board:        config.InitialBoard,
	}
	s.GameState.Board.Init()
	s.logger.Debug("SudokuManager Iniciado.")
}

func (s *SudokuManager) setLogger() {
	logging.SetLevel(s.GameConfig.LoggerLevel)
	s.logger = logging.LoggerFactory("services/SudokuManager")
}

// Step search for the next step based on a array of strategies:
//   - take the first strategy
//   - search for a step
//   - check for a error or a step was not found. If so continue for the next strategy.
//   - take the step
func (s *SudokuManager) Step() (gamestate.Step, error) {
	s.logger.Info("Preparando o próximo passo")
	for _, strateg := range s.GameConfig.Strategies {
		s.logger.Info(fmt.Sprintf("Procurando próximo passo com estratégia %v", strateg))
		step, err := strateg.Step(s.GameState)

		if err != nil || step == nil {
			continue
		}

		s.GameState.Steps = append(s.GameState.Steps, step)

		s.logger.Info(fmt.Sprintf("Passo encontrado, tomando passo: %v", step))
		return step, step.TakeStep(s.GameState.Board)
	}

	s.logger.Error("Não foi possível determinar o próximo passo")
	return nil, errors.New("Não foi possível determinar o próximo passo")
}

// StepAll continues to take steps until find a error or finish the puzzle.
func (s *SudokuManager) StepAll() error {
	s.logger.Info("Preparando para tomar todos os passos possíveis")
	for {
		_, err := s.Step()

		if err != nil {
			return err
		}
		if valid := s.ValidState(); valid != nil {
			return valid
		}
		if s.Finished() {
			break
		}
	}

	s.logger.Info("Todos os passos foram tomados, sudoku concluído")
	return nil
}

func (s *SudokuManager) ValidState() error {
	return s.GameState.Valid()
}

func (s *SudokuManager) Finished() bool {
	return s.GameState.Board.Finished()
}

func (s *SudokuManager) GetState() gamestate.GameState {
	return s.GameState
}
func (s *SudokuManager) ToFile(filename string) error {
	data, err := json.MarshalIndent(&s, "", "    ")

	if err != nil {
		return err
	}
	
	return io.SaveToFile(data, filename)
}

func (s *SudokuManager) InitFromFile(filename string) error  {
	bytes, err := io.ReadFile(filename)

	if err != nil {
		return err
	}
	
	if err := json.Unmarshal(bytes, s); err != nil {
		return err
	}
	
	s.GameConfig.InitialBoard = s.GameState.Board
	s.Init(s.GameConfig)

	return nil	
}

func (s *SudokuManager) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct{
		Config models.GameConfig
		GameState gamestate.GameState
	}{
		Config: s.GameConfig,
		GameState: s.GameState,
	})
}

func (s *SudokuManager) UnmarshalJSON(b []byte) error {
	aux := &struct{
		Config models.GameConfig
		GameState gamestate.GameState
	}{}

	if err := json.Unmarshal(b, aux); err != nil {
		return err
	}

	s.GameConfig = aux.Config
	s.GameState = aux.GameState

	return nil
}
