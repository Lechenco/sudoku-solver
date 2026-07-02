package sudokusolver_test

import (
	sudokusolver "Lechenco/sudoku-solver"
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/gamestate"
	"Lechenco/sudoku-solver/internal/strategy"
	"Lechenco/sudoku-solver/services"
	"context"
	"errors"
	"flag"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/cucumber/godog"
)

type sudokuCtxKey struct{}

type sudokuCtx struct {
	solver     sudokusolver.SudokuSolver
	strategies []strategy.Strategy
	lastStep   gamestate.Step
	err        error
}

func aoUtilizarAEstratgia(ctx context.Context, estrategia string) (context.Context, error) {
	strategs := []strategy.Strategy{}
	if ctx != nil && ctx.Value(sudokuCtxKey{}) != nil {
		sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)
		strategs = sudokuctx.strategies
	}

	var strat strategy.Strategy

	switch estrategia {
	case "naked_single":
		strat = strategy.NakedSingleStrategyInstance()
	case "hidden_single":
		strat = strategy.HiddenSingleStrategyInstance()
	default:
		return ctx, errors.New("Estratégia não implementada.")
	}

	strategs = append(strategs, strat)
	return context.WithValue(ctx, sudokuCtxKey{}, sudokuCtx{
		strategies: strategs,
	}), nil
}

func oPrximoPassoNaPosio(ctx context.Context, value, row, column int) (context.Context, error) {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	step, err := sudokuctx.solver.Step()

	if err != nil {
		return ctx, err
	}
	if s, ok := step.(*gamestate.SetValueStep); ok {
		if s.Value != cells.Value(value) {
			return ctx, fmt.Errorf("Esperava o valor %v mas encontrou o valor %v na posicao (%v)", value, s.Value, step.GetPosition())
		}
	}
	if step.GetPosition().ColumnNumber != uint8(column) ||
		step.GetPosition().RowNumber != uint8(row) {
		return ctx, fmt.Errorf("Esperava a posição (%d,%d) mas encontrou a posição (%v)",
			row, column, step.GetPosition())
	}

	return context.WithValue(ctx, sudokuCtxKey{}, sudokuCtx{
		strategies: sudokuctx.strategies,
		lastStep: step,
		solver: sudokuctx.solver,
		err: sudokuctx.err,
	}), nil
}

func noFoiPossvelDeterminarOPrximoPasso(ctx context.Context) (context.Context, error) {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	step, err := sudokuctx.solver.Step()

	if err != nil {
		return ctx, nil
	}

	return ctx, fmt.Errorf("Não esperava um próximo passo, mas encontrou o passo %v", step)
}

func termineOTabuleiro(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)
    
	err := sudokuctx.solver.StepAll()

	if err != nil {
		return fmt.Errorf("Não foi possível completar o tabuleiro, erro: %v", err)
	}
	
	if !sudokuctx.solver.Finished() {
		return errors.New("Tabuleiro não foi completado.")
	}

	return nil
}

func tomeEssePasso(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.lastStep == nil {
		return errors.New("Nenhum passo foi encontrado")
	}

	manager, ok := sudokuctx.solver.GameManager.(*services.SudokuManager)

	if !ok {
		return errors.New("Não foi possível recuperar o GameManager")
	}

	board := manager.GameState.Board
	position := sudokuctx.lastStep.GetPosition()
	cell := board.Cells[position.RowNumber][position.ColumnNumber]
	switch s := sudokuctx.lastStep.(type) {
	case *gamestate.SetValueStep:
		if cell.Value != s.Value {
			return fmt.Errorf("Passo %v não foi tomado na célula %v", s, cell)
		}

	case *gamestate.RemoveCandidatesStep:
		break
	default:
		return errors.New("Tipo de passo não reconhecido.")
	}

	return nil
}

func oJogoDeveSerInvlido(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.err == nil {
		return errors.New("O jogo foi considerado válido")
	}

	return nil
}

func oJogoDeveSerVlido(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.err != nil {
		return fmt.Errorf("O jogo foi considerado inválido: %s", sudokuctx.err)
	}

	return nil
}

func tabuleiroVlido(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.err != nil {
		return fmt.Errorf("O jogo foi considerado inválido, %v", sudokuctx.err)
	}

	return nil
}

func oTabuleiroAbaixo(ctx context.Context, arg1 *godog.DocString) (context.Context, error) {
	strateg := []strategy.Strategy{}

	if ctx != nil && ctx.Value(sudokuCtxKey{}) != nil {
		sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)
		strateg = sudokuctx.strategies
	}
	re := regexp.MustCompile("[ ┌─┬┐│├─┼┤└─┴┘]")
	s := re.ReplaceAllString(arg1.Content, "")
	solver, err := sudokusolver.NewSudokuSolver(strings.ReplaceAll(s, "\n", ""),
		strateg,
	)

	return context.WithValue(ctx, sudokuCtxKey{}, sudokuCtx{
		solver: solver, err: err, strategies: strateg,
	}), nil
}

var tags = flag.String("godog.tags", "", "tags to execute")
func TestFeature(t *testing.T) {

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"tests/features"},
			TestingT: t, // Testing instance that will run subtests.
			Dialect:  "pt",
			Tags: *tags,
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^o jogo deve ser inválido$`, oJogoDeveSerInvlido)
	sc.Step(`^o jogo deve ser válido$`, oJogoDeveSerVlido)
	sc.Step(`^tabuleiro válido$`, tabuleiroVlido)
	sc.Step(`^o tabuleiro abaixo:$`, oTabuleiroAbaixo)
	sc.Step(`^a estratégia "([^"]*)"$`, aoUtilizarAEstratgia)
	sc.Step(`^o próximo passo é (\d+) na posição \((\d+),(\d+)\)$`, oPrximoPassoNaPosio)
	sc.Step(`^não foi possível determinar o próximo passo$`, noFoiPossvelDeterminarOPrximoPasso)
	sc.Step(`^tome esse passo$`, tomeEssePasso)
	sc.Step(`^termine o tabuleiro$`, termineOTabuleiro)
}
