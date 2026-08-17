package services

import (
	"iter"

	"github.com/Lechenco/sudoku-solver/models/regions"
)



func UnclearRegionsIterator(board BoardService) iter.Seq[regions.Region] {
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
