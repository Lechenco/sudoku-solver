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

func TestUniques(t *testing.T) {
	t.Run("Unique bits between two set", func(t *testing.T) {
		assert := assert.New(t)
		res := cells.Uniques(cells.ValuesSet(0b011), cells.ValuesSet(0b110))

		assert.True(res.Has(1))
		assert.False(res.Has(2))
		assert.True(res.Has(3))
	})
	
	t.Run("Unique bits between three set", func(t *testing.T) {
		assert := assert.New(t)
		res := cells.Uniques(
			cells.ValuesSet(0b011), 
			cells.ValuesSet(0b110),
			cells.ValuesSet(0b001),
		)

		assert.False(res.Has(1))
		assert.False(res.Has(2))
		assert.True(res.Has(3))
	})

	t.Run("Unique bits between severous sets", func(t *testing.T) {
		assert := assert.New(t)
		res := cells.Uniques(
			cells.ValuesSet(0b00011), 
			cells.ValuesSet(0b11101),
			cells.ValuesSet(0b11011),
			cells.ValuesSet(0b01001),
		)

		assert.False(res.Has(1))
		assert.False(res.Has(2))
		assert.True(res.Has(3))
		assert.False(res.Has(4))
		assert.False(res.Has(5))
	})

	t.Run("Unique bits between sets exclude zeros", func(t *testing.T) {
		assert := assert.New(t)
		res := cells.Uniques(cells.ValuesSet(0b011), cells.ValuesSet(0b010))

		assert.True(res.Has(1))
		assert.False(res.Has(2))
		assert.False(res.Has(3))
	})
	
}
