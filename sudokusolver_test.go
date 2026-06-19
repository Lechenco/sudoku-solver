package sudokusolver_test

import (
	sudokusolver "Lechenco/sudoku-solver"
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/strategy"
	"context"
	"errors"
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
	if step.Value != cells.Value(value) {
		return ctx, fmt.Errorf("Esperava o valor %v mas encontrou o valor %v na posicao (%v)", value, step.Value, step.Position)
	}
	if step.Position.ColumnNumber != uint8(column) ||
		step.Position.RowNumber != uint8(row) {
		return ctx, fmt.Errorf("Esperava a posição (%d,%d) mas encontrou a posição (%v)",
			row, column, step.Position)
	}

	return ctx, nil
}

func oJogoDeveSerInvlido(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.err == nil {
		return errors.New("O jogo foi considerado válido")
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

func TestFeature(t *testing.T) {
	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{"tests/features"},
			TestingT: t, // Testing instance that will run subtests.
			Dialect:  "pt",
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	sc.Step(`^o jogo deve ser inválido$`, oJogoDeveSerInvlido)
	sc.Step(`^tabuleiro válido$`, tabuleiroVlido)
	sc.Step(`^o tabuleiro abaixo:$`, oTabuleiroAbaixo)
	sc.Step(`^a estratégia "([^"]*)"$`, aoUtilizarAEstratgia)
	sc.Step(`^o próximo passo é (\d+) na posição \((\d+),(\d+)\)$`, oPrximoPassoNaPosio)
}
