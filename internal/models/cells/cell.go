package cells

type Cell struct {
	Value Value
	Candidates
}

func (c Cell) IsEmpty() bool {
	return c.Value == 0
}

func (c Cell) String() string {
	return c.Value.String()
}

func ToCell(value Value) *Cell {
	var candidates Candidates
	if value == 0 {
		candidates = Candidates(0x1FF) // All candidates (1-9) are possible
	}
	return &Cell{
		Value:      value,
		Candidates: candidates,
	}
}
