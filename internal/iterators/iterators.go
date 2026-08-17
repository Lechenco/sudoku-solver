package iterators

import (
	"iter"

	"github.com/Lechenco/sudoku-solver/models/board"
	"github.com/Lechenco/sudoku-solver/models/cells"
)

func CellsIterator(grid board.CellGrid) iter.Seq[cells.Cell] {
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

func EmptyCellsIterator(grid board.CellGrid) iter.Seq[cells.Cell] {
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

