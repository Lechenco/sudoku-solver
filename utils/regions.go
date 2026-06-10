package utils

import "Lechenco/sudoku-solver/internal/models/cells"

func IndexSquareRegion(i, j int) (int, int) {
	row, col := (i/3)%3, (j/3)%3

	return row, col
}

func IndexSquareRegionPos(position cells.Position) (int, int) {
	return IndexSquareRegion(int(position.RowNumber), int(position.ColumnNumber))
}
