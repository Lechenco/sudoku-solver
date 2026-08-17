package format

import (
	"strings"

	"github.com/Lechenco/sudoku-solver/models/board"
	"github.com/Lechenco/sudoku-solver/models/cells"
)

func BoardFromString(s string) board.Board {
	var board board.Board
	var count int

	if len(s) > 81 {
		s = s[:81]
	}

	for i, c := range s {
		if c >= '1' && c <= '9' {
			board.Cells[i/9][i%9] = cells.ToCell(cells.ToValue(c))
		} else {
			board.Cells[i/9][i%9] = cells.ToCell(0)
		}
		count++
	}

	for ; count < 81; count++ {
		board.Cells[count/9][count%9] = cells.ToCell(0)
	}

	return board
}

func BoardToString(board board.Board) string {
	var s strings.Builder
	topBorder := "┌───────┬───────┬───────┐\n"
	midBorder := "├───────┼───────┼───────┤\n"
	bottomBorder := "└───────┴───────┴───────┘\n"

	s.WriteString(topBorder)
	for i, row := range board.Cells {
		s.WriteString("│ ")

		for j, cell := range row {
			if cell.IsEmpty() {
				s.WriteString("_ ")
			} else {
				s.WriteString(cell.Value.String())
				s.WriteString(" ")
			}
			if (j+1)%3 == 0 {
				s.WriteString("│ ")
			}
		}
		s.WriteString("\n")
		if (i+1)%3 == 0 && i != 8 {
			s.WriteString(midBorder)
		}
	}
	s.WriteString(bottomBorder)
	return s.String()
}
