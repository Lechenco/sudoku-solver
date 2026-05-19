package models

type Candidates uint16

func (c Candidates) Has(value Value) bool {
	return c&(1<<(value-1)) != 0
}

func (c *Candidates) Add(value Value) {
	*c |= 1 << (value - 1)
}

func (c *Candidates) Remove(value Value) {
	*c &^= 1 << (value - 1)
}

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

func ToCell(value Value) Cell {
	var candidates Candidates
	if value == 0 {
		candidates = Candidates(0x1FF) // All candidates (1-9) are possible
	}
	return Cell{
		Value:      value,
		Candidates: candidates,
	}
}
