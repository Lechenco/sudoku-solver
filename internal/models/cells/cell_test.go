package cells_test

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetValue(t *testing.T) {
	assert := assert.New(t)

	t.Run("Set Cell value should change value", func(t *testing.T) {
		cell := cells.Cell{Value: cells.Value(0)}

		cell.SetValue(2)

		assert.Equal(cells.Value(2), cell.Value)
	})

	t.Run("Set Cell value should clean Candidates", func(t *testing.T) {
		candidate := cells.ValuesSet(0b01110)
		cell := cells.Cell{Value: cells.Value(0),
			Candidates: candidate,
		}

		cell.SetValue(2)

		expected := cells.ValuesSet(0x0)
		assert.Equal(expected, cell.Candidates)
	})
}
