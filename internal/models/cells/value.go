package cells

type Value uint8

func ToValue(c rune) Value {
	if c >= '1' && c <= '9' {
		return Value(c - '0')
	}
	return 0
}

func (v Value) String() string {
	if v == 0 {
		return "_"
	}
	return string(v + '0')
}
