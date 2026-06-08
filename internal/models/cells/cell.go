package cells

type Cell struct {
	Value      Value
	Candidates ValuesSet
	Position   Position
}

func (c Cell) GetValue() Value {
	return c.Value
}

func (c Cell) IsEmpty() bool {
	return c.Value == 0
}

func (c Cell) String() string {
	return c.Value.String()
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
