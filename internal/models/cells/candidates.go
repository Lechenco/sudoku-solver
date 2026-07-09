package cells

import (
	"math/bits"
)

// ValuesSet is a bitset of Value uint8 variables, with methods to help
// verify the puzzle operations
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

// GetNumberOfValues return the number of bits raised in the ValuesSet
func (c ValuesSet) GetNumberOfValues() int  {
	return bits.OnesCount16(uint16(c))
}

// Has verify if value is present in the ValueSet
func (c ValuesSet) Has(value Value) bool {
	return c&(1<<(value-1)) != 0
}

// Add raises the corresponding bit of value
func (c *ValuesSet) Add(value Value) {
	*c |= 1 << (value - 1)
}

// Remove drop the corresponding bit of value
func (c *ValuesSet) Remove(value Value) {
	*c &^= 1 << (value - 1)
}

// IsEmpty returns if ValuesSet has none bit raised
func (c ValuesSet) IsEmpty() bool {
	return c == ValuesSet(0x0)
}

// Clear erase all bits in ValuesSet
func (c *ValuesSet) Clear() {
	*c = 0x0
}

// Uniques iterate over a list of ValuesSet, identifying which Value are set 
// in only one of the ValuesSet, returning a new set with these values.
// Follows the algorithm:
//   - mask: remaining unique bits
//   - seen: bits already seen
//   - for each value v:
//      - erase from mask bits seen twice (in v and in seen)
//		- add to seen the bits in values
//	 - return a AND from mask and seen
func Uniques(values ...ValuesSet) ValuesSet  {
	mask := ValuesSet(0x1FF) 
	seen := ValuesSet(0x0)

	for _, v := range values {
		mask &^= v & seen
		seen |= v 
	}
	return mask & seen
}
