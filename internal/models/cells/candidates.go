package cells

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
