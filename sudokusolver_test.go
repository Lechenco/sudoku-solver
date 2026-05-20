package sudokusolver_test

import (
	sudokusolver "Lechenco/sudoku-solver"
	"context"
	"errors"
	"strings"
	"testing"

	"github.com/cucumber/godog"
)

type sudokuCtxKey struct{}

type sudokuCtx struct {
	solver sudokusolver.SudokuSolver
	err    error
}

func oJogoDeveSerInvlido(ctx context.Context) error {
	sudokuctx := ctx.Value(sudokuCtxKey{}).(sudokuCtx)

	if sudokuctx.err == nil {
		return errors.New("O jogo foi considerado válido")
	}

	return nil
}

func oTabuleiroAbaixo(ctx context.Context, arg1 *godog.DocString) (context.Context, error) {
	s := strings.Trim(arg1.Content, " ")
	solver, err := sudokusolver.NewSudokuSolver(strings.ReplaceAll(s, "\n", ""))

	return context.WithValue(ctx, sudokuCtxKey{}, sudokuCtx{
		solver: solver, err: err,
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
	sc.Step(`^o tabuleiro abaixo:$`, oTabuleiroAbaixo)
}
