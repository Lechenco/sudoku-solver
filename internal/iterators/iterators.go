package iterators

import (
	"Lechenco/sudoku-solver/internal/models"
	"Lechenco/sudoku-solver/internal/models/cells"
	"iter"
)


func CellsIterator(grid models.CellGrid) iter.Seq[cells.Cell] {
	return func(yield func(cells.Cell) bool) {
		for _, row := range grid {
			for _, c := range row {
				if !yield(*c) {
					return 
				}
			}
		}
	}
}

func EmptyCellsIterator(grid models.CellGrid) iter.Seq[cells.Cell] {
	return func(yield func(cells.Cell) bool) {
		for _, row := range grid {
			for _, c := range row {
				if !c.IsEmpty() {
 					continue
				}

				if !yield(*c) {
					return 
				}
			}
		}
	}
}
