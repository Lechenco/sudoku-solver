package models

type Board struct {
	Cells [9][9]Value
}

func NewBoard(cells string) Board {
	var board Board
	for i, c := range cells {
		if c >= '1' && c <= '9' {
			board.Cells[i/9][i%9] = ToValue(c)
		}
	}
	return board
}

func (b Board) String() string {
	var result string
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if b.Cells[i][j] == 0 {
				result += "_"
			} else {
				result += b.Cells[i][j].String()
			}
		}
		result += "\n"
	}
	return result
}
