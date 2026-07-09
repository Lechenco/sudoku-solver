package utils

import "Lechenco/sudoku-solver/internal/models/cells"

// IndexSquareRegion take a grid position and returns the relative square region
// position.
func IndexSquareRegion(i, j int) (row, column int) {
	row, column = (i/3)%3, (j/3)%3

	return 
}

func IndexSquareRegionPos(position cells.Position) (row, column int) {
	return IndexSquareRegion(int(position.RowNumber), int(position.ColumnNumber))
}
