package cells

import (
	"fmt"
)

type Cell struct {
	Value      Value
	Candidates ValuesSet
	Position   Position
}

func (c Cell) IsEmpty() bool {
	return c.Value == 0
}

func (c Cell) String() string {
	return fmt.Sprintf("{Cell[%v]: %v}", c.Position, c.Value)
}

func ToCell(value Value) *Cell {
	var candidates ValuesSet
	if value == 0 {
		candidates = ValuesSet(0x1FF) // All candidates (1-9) are possible
	}

	return &Cell{
		Value:      value,
		Candidates: candidates,
	}
}

func (c *Cell) SetValue(value Value) error {
	if c.Value != Value(0) {
		return fmt.Errorf("%v already filled, overrides are not allowed.", c)
	}

	if value == Value(0) {
		return fmt.Errorf("%v cannot set a cell value to zero", c)
	}

	c.Value = value
	c.Candidates.Clear()

	return nil
}
