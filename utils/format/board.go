package format

import (
	"Lechenco/sudoku-solver/internal/models"
)

func BoardFromString(s string) models.Board {
	var board models.Board
	for i, c := range s {
		if c >= '1' && c <= '9' {
			board.Cells[i/9][i%9] = models.ToCell(models.ToValue(c))
		}
	}
	return board
}

func BoardToString(board models.Board) string {
	var s string
	topBorder := "┌───────┬───────┬───────┐\n"
	midBorder := "├───────┼───────┼───────┤\n"
	bottomBorder := "└───────┴───────┴───────┘\n"

	s += topBorder
	for i, row := range board.Cells {
		s += "│ "

		for j, cell := range row {
			if cell.IsEmpty() {
				s += "_ "
			} else {
				s += cell.String() + " "
			}
			if (j+1)%3 == 0 {
				s += "│ "
			}
		}
		s += "\n"
		if (i+1)%3 == 0 && i != 8 {
			s += midBorder
		}
	}
	s += bottomBorder
	return s
}
