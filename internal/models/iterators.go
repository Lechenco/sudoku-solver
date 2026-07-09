package models

import (
	"Lechenco/sudoku-solver/internal/models/cells"
	"Lechenco/sudoku-solver/internal/models/regions"
	"iter"
)

func CellsIterator(grid CellGrid) iter.Seq[cells.Cell] {
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

func EmptyCellsIterator(grid CellGrid) iter.Seq[cells.Cell] {
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

func UnclearRegionsIterator(board Board) iter.Seq[regions.Region] {
	return func(yield func(regions.Region) bool) {
		for i, v := range board.CleanedRegions {
			if v {
				continue
			}

			if !yield(board.AllRegions[i]) {
				return
			}
		}
	}
}
