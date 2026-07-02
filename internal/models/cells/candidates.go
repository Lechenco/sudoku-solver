package cells

import (
	"math/bits"
)

type ValuesSet uint16

func (c ValuesSet) GetValues() []Value {
	values := []Value{}

	for i := range bits.Len(uint(c)) {
		v := Value(i +1)
		if c.Has(v) {
			values = append(values, v)
		}
	}

	return values
}

func (c ValuesSet) GetNumberOfValues() int  {
	return bits.OnesCount16(uint16(c))
}

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

func Uniques(values ...ValuesSet) ValuesSet  {
	mask := ValuesSet(0x1FF) // mask of remaining unique bits
	seen := ValuesSet(0x0) // bits already seen

	for _, v := range values {
		mask &^= v & seen // erase from mask bits been seen twice
		seen |= v // Update seen bits 
	}
	return mask & seen
}
