package format

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
)

func BoardFromString(s string) models.Board {
	var board models.Board
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

func BoardToString(board models.Board) string {
	var s string
	topBorder := "┌───────┬───────┬───────┐\n"
	midBorder := "├───────┼───────┼───────┤\n"
	bottomBorder := "└───────┴───────┴───────┘\n"

	s += topBorder
	for i, row := range board.GetCells() {
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
