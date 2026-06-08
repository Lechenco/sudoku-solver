package cells

type ValuesSet uint16

func (c ValuesSet) Has(value Value) bool {
	return c&(1<<(value-1)) != 0
}

func (c *ValuesSet) Add(value Value) {
	*c |= 1 << (value - 1)
}

func (c *ValuesSet) Remove(value Value) {
	*c &^= 1 << (value - 1)
}

func (c ValuesSet) IsEmpty() bool {
	return c == ValuesSet(0x0)
}

func (c *ValuesSet) Clear() {
	*c = 0x0
}
