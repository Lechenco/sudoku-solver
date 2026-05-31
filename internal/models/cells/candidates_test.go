package cells_test

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	assert := assert.New(t)

	t.Run("Add value should raise bit", func(t *testing.T) {
		set := cells.ValuesSet(0b0)

		set.Add(cells.Value(3))

		assert.Equal(uint16(0b100), uint16(set))
	})

	t.Run("Add value should raise bit and keep others", func(t *testing.T) {
		set := cells.ValuesSet(0b010)

		set.Add(cells.Value(1))

		assert.Equal(uint16(0b11), uint16(set))
	})

	t.Run("Add value already in set should not chang", func(t *testing.T) {
		set := cells.ValuesSet(0b010)

		set.Add(cells.Value(2))

		assert.Equal(uint16(0b10), uint16(set))
	})
}

func TestRemove(t *testing.T) {
	assert := assert.New(t)

	t.Run("Remove value should lower bit", func(t *testing.T) {
		set := cells.ValuesSet(0b010)

		set.Remove(cells.Value(2))

		assert.Equal(uint16(0b0), uint16(set))
	})

	t.Run("Remove value should lower bit and keep others", func(t *testing.T) {
		set := cells.ValuesSet(0b0101)

		set.Remove(cells.Value(3))

		assert.Equal(uint16(0b1), uint16(set))
	})

	t.Run("Remove value not in set should not change", func(t *testing.T) {
		set := cells.ValuesSet(0b010)

		set.Remove(cells.Value(1))

		assert.Equal(uint16(0b10), uint16(set))
	})
}

func TestHas(t *testing.T) {
	assert := assert.New(t)

	t.Run("Has value should return true only if bit is raised", func(t *testing.T) {
		set := cells.ValuesSet(0b011010)

		assert.False(set.Has(cells.Value(1)))
		assert.True(set.Has(cells.Value(2)))
		assert.False(set.Has(cells.Value(3)))
		assert.True(set.Has(cells.Value(4)))
		assert.True(set.Has(cells.Value(5)))
		assert.False(set.Has(cells.Value(6)))
	})
}
