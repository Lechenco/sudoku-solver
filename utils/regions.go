package utils

import "Lechenco/sudoku-solver/internal/models/cells"

func IndexRegion(i, j int) (int, int) {
	row, col := (i/3)*3+(j/3), (i%3)*3+(j%3)

	return row, col
}

func IndexRegionPos(position cells.Position) (int, int) {
	return IndexRegion(int(position.RowNumber), int(position.ColumnNumber))
}
